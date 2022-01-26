package main

import (
	"context"
	"io"
	"log"

	"github.com/ebobo/grpc_streaming_rpc_go/pkg/go/pb/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	streamServerAddress = "localhost:9092"
)

func main() {
	conn, err := grpc.Dial(streamServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	c := pb.NewStreamServiceClient(conn)

	stream, err := c.TimeInfo(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.GetTimeData())
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")

}

// use go mod tidy to download all the pakages we imported
