package server

import (
	"context"
	"log"
	"os/exec"
	"time"

	"github.com/elalmirante/elalmirante-agent/conf"
	"github.com/elalmirante/elalmirante-agent/rpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeployServiceServer struct {
	Conf *conf.Configuration
}

func (s *DeployServiceServer) Deploy(ctx context.Context, req *rpc.DeployRequest) (*rpc.DeployResponse, error) {
	if req.Key == "" || req.Key != s.Conf.Key {
		log.Println("DeployService: Unauthenticated")
		return nil, status.Errorf(codes.Unauthenticated, "Invalid Key")
	}

	t := time.Now()
	log.Println("DeployService: Started")

	cmd := exec.Command("/bin/sh", "-c", s.Conf.ScriptLine(req.Ref))

	out, err := cmd.CombinedOutput()
	output := string(out)

	log.Printf("DeployService: Output:\n%s\n", output)

	errStr := ""
	if err != nil {
		errStr = err.Error()
		log.Printf("DeployService: Error:\n%s\n", errStr)
	}

	log.Println("DeployService: Finished, took", time.Since(t))
	return &rpc.DeployResponse{Output: output, Error: errStr}, nil
}
