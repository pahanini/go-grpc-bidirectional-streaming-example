package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"os"
	pb "proto"
	"time"

	"google.golang.org/grpc"
)

func main() {

	// dail server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewMathClient(conn)
	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	var max int32
	const step = 5
	ctx := stream.Context()

	// first goroutine send random increasing numbers to stream
	go func() {
		for i := 1; i <= 10; i++ {
			// generate random nummber and send it to stream
			rnd := int32(rand.Intn(i))
			req := pb.Request{Num: rnd}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf("%d sent", req.Num)
			time.Sleep(time.Second)
		}

		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
		os.Exit(0)
	}()

	// second goroutine receives data from stream
	go func() {
		for {
			// receive number from stream and update max
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			max = resp.Result
			log.Printf("new max=%d received", max)
		}
	}()

	<-ctx.Done()
	if err := ctx.Err(); err != nil {
		log.Println(err)
	} else {
		log.Printf("finished with max=%d", max)
	}
}
