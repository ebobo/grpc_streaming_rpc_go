package main

import (
	"time"

	"github.com/ebobo/grpc_streaming_rpc_go/pkg/go/pb/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StreamServer struct {
	pb.UnimplementedStreamServiceServer
}

func (server *StreamServer) TimeInfo(_ *emptypb.Empty, stream pb.StreamService_TimeInfoServer) error {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		if err := stream.Send(&pb.CurrentTime{MessageId: int32(i), TimeData: time.Now().String()}); err != nil {
			return err
		}
	}
	return nil
}
