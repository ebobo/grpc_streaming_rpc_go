package main

import (
	"log"
	"net"
	"os"

	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ebobo/grpc_streaming_rpc_go/pkg/go/pb/v1"
)

var opt struct {
	GRPCAddr string `long:"grpc-addr" default:":9092" description:"gRPC listen address"`
}

func main() {
	_, err := flags.ParseArgs(&opt, os.Args)
	if err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}

	var streamServer = &StreamServer{}
	lis, err := net.Listen("tcp", opt.GRPCAddr)
	if err != nil {
		log.Fatalf("unable to listen %v", err)
	}
	gs := grpc.NewServer()
	reflection.Register(gs)

	pb.RegisterStreamServiceServer(gs, streamServer)

	log.Printf("server listening at %v", lis.Addr())

	gs.Serve(lis)
}
