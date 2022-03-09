package main

import (
	"flag"
	"fmt"
	"os"
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	
	"github.com/slabs-forge/grpctst/pkg/common"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "20060102 15:04:05.000"})
	
	port := flag.Int("port",9090,"Listen port of the server")
	
	flag.Parse()

	var opts []grpc.DialOption
	
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	conn , err := grpc.Dial(fmt.Sprintf(":%d",*port), opts...)
	if err != nil  {
		log.Fatal().Err(err).Int("port",*port).Msg("Can't connect")
	}
	defer conn.Close()
	
	log.Info().Int("port",*port).Msg("Connected")

	client := common.NewGrpcTestClient(conn)

	for i := 1 ; i < 10 ; i++ {
		res , err := client.Test(context.TODO(),&common.TestRequest{Message: fmt.Sprintf("Hello %d",i)})
		if err != nil  {
			log.Error().Err(err).Int("i",i).Msg("Can't call Test")
		}
			
		log.Info().Str("result", res.Message).Msg("Test called")
	}	

}
