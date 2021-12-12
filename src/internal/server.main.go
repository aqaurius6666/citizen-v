package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"cloud.google.com/go/profiler"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/services/jwt"
	commongrpc "github.com/aqaurius6666/go-utils/common_grpc"
	commonpb "github.com/aqaurius6666/go-utils/common_grpc/pb"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func runMain(appCtx *cli.Context) (err error) {

	var wg sync.WaitGroup
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	err = validate.RegisterValidator()
	if err != nil {
		return err
	}
	errChan := make(chan error, 1)
	if appCtx.Bool("disable-tracing") {
		logger.Info("Tracing disabled.")
	} else {
		logger.Info("Tracing enabled.")
		go initTracing()
	}
	if appCtx.Bool("disable-profiler") {
		logger.Info("Profiling disabled.")
	} else {
		logger.Info("Profiling enabled.")
		go initProfiling(serviceName, appCtx.String("runtime-version"))
	}

	// Start HTTP Server
	httpListner, err := net.Listen("tcp", fmt.Sprintf(":%d", appCtx.Int("http-port")))
	if err != nil {
		logger.Fatal(err)
		return err
	}
	mainServer, err := InitMainServer(ctx, logger, ServerOptions{
		DBDsn: db.DBDsn(appCtx.String("db-uri")),
		Sec:   jwt.SecretKey(appCtx.String("jwt-secret-key")),
	})
	if err != nil {
		logger.Fatal(err)
		return err
	}
	defer func() {
		_ = httpListner.Close()
		_ = mainServer.MainRepo.Close()
	}()
	err = mainServer.MainRepo.Migrate()
	if err != nil {
		logger.Fatal(err)
		return err
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		mainServer.ApiServer.RegisterEndpoint()
		logger.WithField("port", appCtx.Int("http-port")).Info("listening for HTTP connections")
		if err := http.Serve(httpListner, mainServer.ApiServer.G); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start GRPC Server
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", appCtx.Int("grpc-port")))
	if err != nil {
		logger.Fatal(err)
		return err
	}
	defer func() { _ = grpcListener.Close() }()
	var srv *grpc.Server
	commonServer := commongrpc.NewCommonServer(logger, appCtx.Bool("allow-kill"))
	wg.Add(1)
	go func() {
		defer wg.Done()
		if appCtx.Bool("disable-stats") {
			logger.Info("Stats disabled.")
			srv = grpc.NewServer()
		} else {
			logger.Info("Stats enabled.")
			srv = grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
		}
		healthpb.RegisterHealthServer(srv, commonServer)
		commonpb.RegisterCommonServer(srv, commonServer)
		reflection.Register(srv)
		logger.WithField("port", appCtx.Int("grpc-port")).Info("listening for gRPC connections")
		if err := srv.Serve(grpcListener); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start pprof server
	pprofListener, err := net.Listen("tcp", fmt.Sprintf(":%d", appCtx.Int("pprof-port")))
	if err != nil {
		logger.Fatal(err)
		return err
	}
	defer func() {
		_ = pprofListener.Close()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.WithField("port", appCtx.Int("pprof-port")).Info("listening for pprof requests")
		sSrv := new(http.Server)
		_ = sSrv.Serve(pprofListener)
	}()

	if !appCtx.Bool("disable-prometheus") {
		// Start Prometheus Server
		promListener, err := net.Listen("tcp", fmt.Sprintf(":%d", appCtx.Int("prometheus-port")))
		if err != nil {
			return err
		}
		defer func() {
			_ = promListener.Close()
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			promServer := http.NewServeMux()
			promServer.Handle("/metrics", promhttp.Handler())
			logger.WithField("port", appCtx.Int("prometheus-port")).Info("listening for metrics requests")
			if err := http.Serve(promListener, promServer); err != nil {
				errChan <- err
			}
		}()
	} else {
		logger.Info("Prometheus disabled.")
	}

	// Watch kill signal
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
		select {
		case s := <-sigCh:
			cancelFn()
			// Handle graceful shutdown here

			logger.WithField("signal", s.String()).Infof("shutting down due to signal")
		case <-ctx.Done():

		case err := <-errChan:
			cancelFn()
			logger.WithField("error", err.Error()).Errorf("shutting down due to error")
		}
	}()
	wg.Wait()
	return nil
}

func initTracing() {
	// initJaegerTracing()
	initStackdriverTracing()
}

func initStats(exporter *stackdriver.Exporter) {
	view.SetReportingPeriod(60 * time.Second)
	view.RegisterExporter(exporter)
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		logger.Warn("Error registering default server views")
	} else {
		logger.Info("Registered default server views")
	}
}

func initStackdriverTracing() {
	// TODO(ahmetb) this method is duplicated in other microservices using Go
	// since they are not sharing packages.
	for i := 1; i <= 3; i++ {
		exporter, err := stackdriver.NewExporter(stackdriver.Options{})
		if err != nil {
			logger.Infof("failed to initialize stackdriver exporter: %+v", err)
		} else {
			trace.RegisterExporter(exporter)
			logger.Info("registered Stackdriver tracing")

			// Register the views to collect server stats.
			initStats(exporter)
			return
		}
		d := time.Second * 10 * time.Duration(i)
		logger.Infof("sleeping %v to retry initializing Stackdriver exporter", d)
		time.Sleep(d)
	}
	logger.Warn("could not initialize Stackdriver exporter after retrying, giving up")
}

func initProfiling(service, version string) {
	// TODO(ahmetb) this method is duplicated in other microservices using Go
	// since they are not sharing packages.
	for i := 1; i <= 3; i++ {
		if err := profiler.Start(profiler.Config{
			Service:        service,
			ServiceVersion: version,
			// ProjectID must be set if not running on GCP.
			// ProjectID: "my-project",
		}); err != nil {
			logger.Warnf("failed to start profiler: %+v", err)
		} else {
			logger.Info("started Stackdriver profiler")
			return
		}
		d := time.Second * 10 * time.Duration(i)
		logger.Infof("sleeping %v to retry initializing Stackdriver profiler", d)
		time.Sleep(d)
	}
	logger.Warn("could not initialize Stackdriver profiler after retrying, giving up")
}
