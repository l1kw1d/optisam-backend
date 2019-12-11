protoc --proto_path=auth-service/api/proto/v1 --proto_path=common/third_party --go_out=plugins=grpc:auth-service/pkg/api/v1 auth.proto
protoc --proto_path=auth-service/api/proto/v1 --proto_path=common/third_party --grpc-gateway_out=logtostderr=true:auth-service/pkg/api/v1 auth.proto
protoc --proto_path=auth-service/api/proto/v1 --proto_path=common/third_party --swagger_out=logtostderr=true:auth-service/api/swagger/v1 auth.proto
# protoc --proto_path=auth-service/api/proto/v1 --proto_path=common/third_party --go_out=auth-service/pkg/api/v1 --validate_out=lang=go:auth-service/pkg/api/v1 auth.proto