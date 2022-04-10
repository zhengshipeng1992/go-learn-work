package service

import (
	"context"
	work_v1 "go-learn-work/work03/api/work/v1"
	"go-learn-work/work03/internel/biz"
)

type Hello struct {
	biz.HelloRequest
}

func NewHello(request biz.HelloRequest) *Hello {
	return &Hello{request}
}

func (h *Hello) SayHello(ctx context.Context, request *work_v1.HelloRequest) (*work_v1.HelloResponse, error) {
	bizRes := h.BizHello(ctx, biz.NewHelloReq(request.Name))
	return &work_v1.HelloResponse{Message: bizRes.GetMessage()}, nil
}
