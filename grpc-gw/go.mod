module github.com/gucchisk/go-samples/grpc-gw

go 1.14

replace github.com/gucchisk/go-samples/grpc-gw => ./

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9
	google.golang.org/genproto v0.0.0-20200605102947-12044bf5ea91
	google.golang.org/grpc v1.30.0-dev.1
	google.golang.org/grpc/examples v0.0.0-20200604175613-ad51f572fd27 // indirect
	google.golang.org/protobuf v1.24.0
)
