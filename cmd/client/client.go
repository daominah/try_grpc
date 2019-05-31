package main

import (
	"context"
	"fmt"
	"log"
	"sync"
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

	var locker sync.Mutex
	var waiter sync.WaitGroup
	nSuccess, nAll := 0, 0
	for i := 0; i < 30; i++ {
		waiter.Add(1)
		go func(i int) {
			r2, err := Func2(c, 2, 3)
			_, _ = r2, err
			//fmt.Println("pussy", i, r2, err)
			locker.Lock()
			nAll +=1
			if err == nil {
				nSuccess += 1
			}
			locker.Unlock()
			waiter.Add(-1)
		}(i)
	}
	waiter.Wait()
	fmt.Printf("result: nAll %v, nSuc %v.\n", nAll, nSuccess)
}
