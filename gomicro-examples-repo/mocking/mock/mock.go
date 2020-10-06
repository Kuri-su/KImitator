package mock

import (
	"context"

	proto "github.com/micro/examples/helloworld/proto"
	"github.com/micro/go-micro/v2/client"
)

type mockGreeterService struct {
}

func (m *mockGreeterService) Call(ctx context.Context, req *proto.Request, opts ...client.CallOption) (*proto.Response, error) {
	return &proto.Response{
		Msg: "Hello " + req.Name,
	}, nil
}

func NewGreeterService() proto.HelloworldService {
	return new(mockGreeterService)
}
