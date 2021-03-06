// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: platform/proto/signup/signup.proto

package signup

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Signup service

func NewSignupEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Signup service

type SignupService interface {
	// Send a verification email
	SendVerificationEmail(ctx context.Context, in *SendVerificationEmailRequest, opts ...client.CallOption) (*SendVerificationEmailResponse, error)
	// Verify a user who's been sent an email
	Verify(ctx context.Context, in *VerifyRequest, opts ...client.CallOption) (*VerifyResponse, error)
	// Complete the signup process
	CompleteSignup(ctx context.Context, in *CompleteSignupRequest, opts ...client.CallOption) (*CompleteSignupResponse, error)
}

type signupService struct {
	c    client.Client
	name string
}

func NewSignupService(name string, c client.Client) SignupService {
	return &signupService{
		c:    c,
		name: name,
	}
}

func (c *signupService) SendVerificationEmail(ctx context.Context, in *SendVerificationEmailRequest, opts ...client.CallOption) (*SendVerificationEmailResponse, error) {
	req := c.c.NewRequest(c.name, "Signup.SendVerificationEmail", in)
	out := new(SendVerificationEmailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupService) Verify(ctx context.Context, in *VerifyRequest, opts ...client.CallOption) (*VerifyResponse, error) {
	req := c.c.NewRequest(c.name, "Signup.Verify", in)
	out := new(VerifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupService) CompleteSignup(ctx context.Context, in *CompleteSignupRequest, opts ...client.CallOption) (*CompleteSignupResponse, error) {
	req := c.c.NewRequest(c.name, "Signup.CompleteSignup", in)
	out := new(CompleteSignupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Signup service

type SignupHandler interface {
	// Send a verification email
	SendVerificationEmail(context.Context, *SendVerificationEmailRequest, *SendVerificationEmailResponse) error
	// Verify a user who's been sent an email
	Verify(context.Context, *VerifyRequest, *VerifyResponse) error
	// Complete the signup process
	CompleteSignup(context.Context, *CompleteSignupRequest, *CompleteSignupResponse) error
}

func RegisterSignupHandler(s server.Server, hdlr SignupHandler, opts ...server.HandlerOption) error {
	type signup interface {
		SendVerificationEmail(ctx context.Context, in *SendVerificationEmailRequest, out *SendVerificationEmailResponse) error
		Verify(ctx context.Context, in *VerifyRequest, out *VerifyResponse) error
		CompleteSignup(ctx context.Context, in *CompleteSignupRequest, out *CompleteSignupResponse) error
	}
	type Signup struct {
		signup
	}
	h := &signupHandler{hdlr}
	return s.Handle(s.NewHandler(&Signup{h}, opts...))
}

type signupHandler struct {
	SignupHandler
}

func (h *signupHandler) SendVerificationEmail(ctx context.Context, in *SendVerificationEmailRequest, out *SendVerificationEmailResponse) error {
	return h.SignupHandler.SendVerificationEmail(ctx, in, out)
}

func (h *signupHandler) Verify(ctx context.Context, in *VerifyRequest, out *VerifyResponse) error {
	return h.SignupHandler.Verify(ctx, in, out)
}

func (h *signupHandler) CompleteSignup(ctx context.Context, in *CompleteSignupRequest, out *CompleteSignupResponse) error {
	return h.SignupHandler.CompleteSignup(ctx, in, out)
}
