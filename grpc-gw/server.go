/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net"
	"net/http"
	"strings"

	pb "github.com/gucchisk/go-samples/grpc-gw/proto/helloworld"
	"google.golang.org/grpc"
)

const (
	port = "9999"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func handlerFunc(ctx context.Context, grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Content-Type")
		fmt.Printf("%s", t)
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r.WithContext(ctx))
		} else {
			httpHandler.ServeHTTP(w, r.WithContext(ctx))
		}
	}), &http2.Server{})
}

func initHTTPServer(ctx context.Context, grpcServer *grpc.Server, gwmux *runtime.ServeMux, addr string) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	server := &http.Server{
		Addr:    addr,
		Handler: handlerFunc(ctx, grpcServer, mux),
	}
	return server
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	ss := &server{}
	pb.RegisterGreeterServer(grpcServer, ss)
	gwmux := runtime.NewServeMux()

	pb.RegisterGreeterHandlerServer(ctx, gwmux, ss)
	server := initHTTPServer(ctx, grpcServer, gwmux, ":"+port)

	log.Println("start server...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
