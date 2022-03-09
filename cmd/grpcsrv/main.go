package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"

	"github.com/slabs-forge/grpctst/pkg/common"
)

type grpcTestServer struct {
	common.UnimplementedGrpcTestServer
}

func (s *grpcTestServer) Test(ctx context.Context, request *common.TestRequest) (*common.TestResponse, error) {
	response := &common.TestResponse{}

	log.Info().Msg("Test call")

	response.Message = fmt.Sprintf("Result is %s",request.Message)
	return response,nil
} 

func newServer() *grpcTestServer {
	return &grpcTestServer{}
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "20060102 15:04:05.000"})
	
	port := flag.Int("port",9090,"Listen port of the server")
	
	flag.Parse()

	ls , err := net.Listen("tcp",fmt.Sprintf(":%d",*port))
	if err != nil {
		log.Fatal().Err(err).Msg("Can't connect")
	}

	var opts []grpc.ServerOption
	
	gs := grpc.NewServer(opts...)
	common.RegisterGrpcTestServer(gs,newServer())
	
	log.Info().Msg("Listening...")
	gs.Serve(ls)
}
