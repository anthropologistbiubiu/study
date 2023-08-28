package service

import (
	"consul/pb"
	"context"
	"fmt"
)

type JobServiceServer struct {
}

func (job *JobServiceServer) GetJobService(ctx context.Context, request *pb.Request) (*pb.Response, error) {

	reposne := &pb.Response{
		Reply: request.Name + ":" + request.Job,
	}
	fmt.Println("run......8081")
	return reposne, nil
}
