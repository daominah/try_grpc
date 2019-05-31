package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/daominah/try_grpc/minahproto"
	"google.golang.org/grpc"
)

func Func1(c minahproto.HelloClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &minahproto.HelloRequest{Name: "hihi"})
	if err != nil {
		return "", err
	}
	return r.Message, nil
}

func Func2(c minahproto.HelloClient, arg1 int64, arg2 int64) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	r, err := c.Add(ctx, &minahproto.AddRequest{Arg1: arg1, Arg2: arg2})
	if err != nil {
		return 0, err
	}
	return r.Sum, nil
}

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := minahproto.NewHelloClient(conn)

	r1, err := Func1(c)
	fmt.Println("pussy", r1, err)

	r2, err := Func2(c, 2, 3)
	fmt.Println("pussy2", r2, err)

	r3, err := Func2(c, 4, 7)
	fmt.Println("pussy2", r3, err)
}
