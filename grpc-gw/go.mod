module github.com/gucchisk/go-samples/grpc-gw

go 1.14

replace github.com/gucchisk/go-samples/grpc-gw => ./

require (
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	golang.org/x/net v0.5.0
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)
