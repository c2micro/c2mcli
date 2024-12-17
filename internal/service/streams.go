package service

import (
	operatorv1 "github.com/c2micro/c2mshr/proto/gen/operator/v1"
	"github.com/lrita/cmap"
	"google.golang.org/grpc"
)

type streams struct {
	controlStream grpc.ServerStreamingClient[operatorv1.HelloResponse]
	tasksStream   grpc.BidiStreamingClient[operatorv1.SubscribeTasksRequest, operatorv1.SubscribeTasksResponse]
	groupStreams  cmap.Map[uint32, grpc.ClientStreamingClient[operatorv1.NewGroupRequest, operatorv1.NewGroupResponse]]
}
