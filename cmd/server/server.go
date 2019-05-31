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

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/daominah/try_grpc/minahproto"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type HelloServerImpl struct{}

// SayHello implements helloworld.GreeterServer
func (s *HelloServerImpl) SayHello(ctx context.Context, in *minahproto.HelloRequest) (
	*minahproto.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &minahproto.HelloReply{Message: "pussy " + in.Name}, nil
}

func (s *HelloServerImpl) Add(ctx context.Context, r *minahproto.AddRequest) (
	*minahproto.AddResponse, error) {
	_ = time.Second
	time.Sleep(2 * time.Second)
	return &minahproto.AddResponse{Sum: r.Arg1 + r.Arg2}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	var hs minahproto.HelloServer
	hs = &HelloServerImpl{}
	minahproto.RegisterHelloServer(s, hs)

	log.Println("Listening on ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
