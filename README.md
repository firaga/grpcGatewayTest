protoc -I . \
--go_out ./gen/go/ --go_opt paths=source_relative \
--go-grpc_out ./gen/go/ --go-grpc_opt paths=source_relative \
your/service/v1/your_service.proto

protoc -I . --grpc-gateway_out ./gen/go \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt paths=source_relative \
your/service/v1/your_service.proto

https://github.com/grpc-ecosystem/grpc-gateway#usage

curl -X POST http://localhost:8081/v1/example/echo -H "Content-Type: application/json" -d'{"value": "foo"}'