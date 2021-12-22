GOPATH=$HOME/go
PATH=$PATH:$GOPATH/bin
protodir=./proto

#protoc --go-grpc_out=plugins=grpc:./src -I $protodir $protodir/*.proto
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out ./src/pb --go_opt paths=source_relative \
   -I $protodir \
   --go-grpc_out ./src/pb --go-grpc_opt paths=source_relative \
   --grpc-gateway_out ./src/pb --grpc-gateway_opt paths=source_relative --grpc-gateway_opt allow_delete_body=true \
   --openapiv2_out ./src/internal/services/swagger --openapiv2_opt logtostderr=true --openapiv2_opt allow_delete_body=true \
   $protodir/api.proto 