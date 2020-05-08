// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: prodService.proto

package Service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ProdService service

type ProdService interface {
	GetProdsList(ctx context.Context, in *ProdsReq, opts ...client.CallOption) (*ProdListResp, error)
	GetProdsDetail(ctx context.Context, in *ProdsReq, opts ...client.CallOption) (*ProdDetailtResp, error)
}

type prodService struct {
	c    client.Client
	name string
}

func NewProdService(name string, c client.Client) ProdService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "Service"
	}
	return &prodService{
		c:    c,
		name: name,
	}
}

func (c *prodService) GetProdsList(ctx context.Context, in *ProdsReq, opts ...client.CallOption) (*ProdListResp, error) {
	req := c.c.NewRequest(c.name, "ProdService.GetProdsList", in)
	out := new(ProdListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodService) GetProdsDetail(ctx context.Context, in *ProdsReq, opts ...client.CallOption) (*ProdDetailtResp, error) {
	req := c.c.NewRequest(c.name, "ProdService.GetProdsDetail", in)
	out := new(ProdDetailtResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProdService service

type ProdServiceHandler interface {
	GetProdsList(context.Context, *ProdsReq, *ProdListResp) error
	GetProdsDetail(context.Context, *ProdsReq, *ProdDetailtResp) error
}

func RegisterProdServiceHandler(s server.Server, hdlr ProdServiceHandler, opts ...server.HandlerOption) error {
	type prodService interface {
		GetProdsList(ctx context.Context, in *ProdsReq, out *ProdListResp) error
		GetProdsDetail(ctx context.Context, in *ProdsReq, out *ProdDetailtResp) error
	}
	type ProdService struct {
		prodService
	}
	h := &prodServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProdService{h}, opts...))
}

type prodServiceHandler struct {
	ProdServiceHandler
}

func (h *prodServiceHandler) GetProdsList(ctx context.Context, in *ProdsReq, out *ProdListResp) error {
	return h.ProdServiceHandler.GetProdsList(ctx, in, out)
}

func (h *prodServiceHandler) GetProdsDetail(ctx context.Context, in *ProdsReq, out *ProdDetailtResp) error {
	return h.ProdServiceHandler.GetProdsDetail(ctx, in, out)
}