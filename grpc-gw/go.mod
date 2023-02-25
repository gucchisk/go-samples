module github.com/gucchisk/go-samples/grpc-gw

go 1.14

replace github.com/gucchisk/go-samples/grpc-gw => ./

require (
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	golang.org/x/net v0.7.0
	google.golang.org/genproto v0.0.0-20200605102947-12044bf5ea91
	google.golang.org/grpc v1.30.0-dev.1
	google.golang.org/protobuf v1.24.0
)
