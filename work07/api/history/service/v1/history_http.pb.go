// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type HistoryHTTPServer interface {
	SaveHistory(context.Context, *SaveHistoryRequest) (*SaveHistoryReply, error)
}

func RegisterHistoryHTTPServer(s *http.Server, srv HistoryHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/history", _History_SaveHistory0_HTTP_Handler(srv))
}

func _History_SaveHistory0_HTTP_Handler(srv HistoryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveHistoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.history.service.v1.History/SaveHistory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveHistory(ctx, req.(*SaveHistoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveHistoryReply)
		return ctx.Result(200, reply)
	}
}

type HistoryHTTPClient interface {
	SaveHistory(ctx context.Context, req *SaveHistoryRequest, opts ...http.CallOption) (rsp *SaveHistoryReply, err error)
}

type HistoryHTTPClientImpl struct {
	cc *http.Client
}

func NewHistoryHTTPClient(client *http.Client) HistoryHTTPClient {
	return &HistoryHTTPClientImpl{client}
}

func (c *HistoryHTTPClientImpl) SaveHistory(ctx context.Context, in *SaveHistoryRequest, opts ...http.CallOption) (*SaveHistoryReply, error) {
	var out SaveHistoryReply
	pattern := "/v1/history"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.history.service.v1.History/SaveHistory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}