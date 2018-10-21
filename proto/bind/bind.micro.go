// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/bind/bind.proto

/*
Package go_micro_srv_tenno_ucenter is a generated protocol buffer package.

It is generated from these files:
	proto/bind/bind.proto

It has these top-level messages:
	BindRequest
	UnbindRequest
	GetBindListRequest
	Response
*/
package go_micro_srv_tenno_ucenter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Bind service

type BindService interface {
	Bind(ctx context.Context, in *BindRequest, opts ...client.CallOption) (*Response, error)
	UnBind(ctx context.Context, in *UnbindRequest, opts ...client.CallOption) (*Response, error)
	GetBindList(ctx context.Context, in *GetBindListRequest, opts ...client.CallOption) (*Response, error)
}

type bindService struct {
	c    client.Client
	name string
}

func NewBindService(name string, c client.Client) BindService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.tenno.ucenter"
	}
	return &bindService{
		c:    c,
		name: name,
	}
}

func (c *bindService) Bind(ctx context.Context, in *BindRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Bind.Bind", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindService) UnBind(ctx context.Context, in *UnbindRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Bind.UnBind", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindService) GetBindList(ctx context.Context, in *GetBindListRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Bind.GetBindList", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bind service

type BindHandler interface {
	Bind(context.Context, *BindRequest, *Response) error
	UnBind(context.Context, *UnbindRequest, *Response) error
	GetBindList(context.Context, *GetBindListRequest, *Response) error
}

func RegisterBindHandler(s server.Server, hdlr BindHandler, opts ...server.HandlerOption) error {
	type bind interface {
		Bind(ctx context.Context, in *BindRequest, out *Response) error
		UnBind(ctx context.Context, in *UnbindRequest, out *Response) error
		GetBindList(ctx context.Context, in *GetBindListRequest, out *Response) error
	}
	type Bind struct {
		bind
	}
	h := &bindHandler{hdlr}
	return s.Handle(s.NewHandler(&Bind{h}, opts...))
}

type bindHandler struct {
	BindHandler
}

func (h *bindHandler) Bind(ctx context.Context, in *BindRequest, out *Response) error {
	return h.BindHandler.Bind(ctx, in, out)
}

func (h *bindHandler) UnBind(ctx context.Context, in *UnbindRequest, out *Response) error {
	return h.BindHandler.UnBind(ctx, in, out)
}

func (h *bindHandler) GetBindList(ctx context.Context, in *GetBindListRequest, out *Response) error {
	return h.BindHandler.GetBindList(ctx, in, out)
}