package gateway

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/gucchisk/go-samples/grpc-gw/proto/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func run(serverPort string, gwPort string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf(":" + serverPort)
	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	log.Printf("gateway port:" + gwPort)
	log.Printf("server listen: " + serverPort)
	return http.ListenAndServe(":"+gwPort, mux)
}

func Start(serverPort string, gwPort string) {
	flag.Parse()
	if err := run(serverPort, gwPort); err != nil {
		log.Fatalln(err)
	}
}
