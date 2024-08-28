mkdir -p ./output
protoc --go_out=./output --go_opt=paths=source_relative --go-grpc_out=./output --go-grpc_opt=paths=source_relative ssl_*