dev-recreate:
	@docker-compose --project-name=go-go --env-file deploy/dev/.env -f deploy/dev/docker-compose.yaml up -d --build --force-recreate

build-and-push-image: build-image push-image

build-image:
	@docker build -f release.Dockerfile . -t aqaurius6666/mainserver:pre-release

push-image:
	@docker tag supermedicalchain/main-service:pre-release supermedicalchain/main-service${TAG}
	@docker push supermedicalchain/main-service${TAG}

build:
	@go build -o ./dist/server ./src/internal

debug:
	@/bin/sh -c "dlv --listen=0.0.0.0:2345 --headless=true --api-version=2 exec ./dist/server -- --disable-profiler --allow-kill serve"

serve:
	@./dist/server serve

dev:
	@./dist/server --log-format plain --log-level debug --disable-profiler --allow-kill serve

cleanDB:
	@./dist/server clean

seed:
	@./dist/server seed-data --clean
	@echo Hello

test:
	#go test ./src/cockroach/... -v -check.f "CockroachDbGraphTestSuite.*"
	@go test ./... -v

test-prepare-up:
	@docker exec  up -f deploy/dev/docker-compose.yaml main-cdb -d

test-prepare-down:
	@docker-compose down -f deploy/dev/docker-compose.yaml main-cdb

grpc-client:
	@grpc-client-cli localhost:${GRPC_PORT}

kill:
	@(echo '{}' | grpc-client-cli -service Common -method Kill localhost:${GRPC_PORT}) > /nil 2> /nil || return 0

logs:
	@docker-compose --project-name=go-go -f deploy/dev/docker-compose.yaml logs -f mainservice

proto:
	@./genproto.sh

swagger: proto
	@go generate ./src/swagger github.com/sotanext-team/medical-chain/src/mainservice/src/swagger 

rebase: 
	@git pull --rebase origin dev

push:
	@git push origin HEAD:automatic-branch -f

prom:
	@docker-compose --project-name=go-go -f deploy/dev/docker-prometheus.yaml up -d --build --force-recreate

lint:
	@golangci-lint run