package rpc

import (
	"context"

	"github.com/mixarchitecture/arch/shared/genproto/example"
)

func (s Server) GetExample(ctx context.Context, req *example.GetExampleRequest) (*example.GetExampleResponse, error) {
	return nil, nil
}

func (s Server) ListExample(ctx context.Context, req *example.ListExampleRequest) (*example.ListExampleResponse, error) {
	return nil, nil
}

func (s Server) CreateExample(ctx context.Context, req *example.CreateExampleRequest) (*example.CreateExampleResponse, error) {
	return nil, nil
}

func (s Server) UpdateExample(ctx context.Context, req *example.UpdateExampleRequest) (*example.UpdateExampleResponse, error) {
	return nil, nil
}
