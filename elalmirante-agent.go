package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/elalmirante/elalmirante-agent/conf"
	"github.com/elalmirante/elalmirante-agent/rpc"
	"github.com/elalmirante/elalmirante-agent/rpc/server"

	"google.golang.org/grpc"
)

func main() {
	// read config file
	var confPath string
	flag.StringVar(&confPath, "c", "/etc/elalmirante-agent.conf", "The path to the configuration file.")
	flag.Parse()

	conf, err := conf.Parse(confPath)
	if err != nil {
		log.Fatal(err)
	}

	// rpc server listen:
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	rpc.RegisterDeployServiceServer(grpcServer, &server.DeployServiceServer{Conf: conf})
	log.Println("Starting gRPC server...")
	grpcServer.Serve(lis)
}
