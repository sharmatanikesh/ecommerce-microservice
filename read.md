protoc --go_out=. --go_opt=module=github.com/sharamatanikesh/ecom-microservice \
       --go-grpc_out=. --go-grpc_opt=module=github.com/sharamatanikesh/ecom-microservice,require_unimplemented_servers=false \
       account/account.proto
