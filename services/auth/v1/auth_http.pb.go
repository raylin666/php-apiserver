// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.4

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

type AuthHTTPServer interface {
	GetRolesForUser(context.Context, *GetRolesForUserRequest) (*GetRolesForUserReply, error)
}

func RegisterAuthHTTPServer(s *http.Server, srv AuthHTTPServer) {
	r := s.Route("/")
	r.GET("/roles/forUser/{user}", _Auth_GetRolesForUser0_HTTP_Handler(srv))
}

func _Auth_GetRolesForUser0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRolesForUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/services.auth.v1.Auth/GetRolesForUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRolesForUser(ctx, req.(*GetRolesForUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRolesForUserReply)
		return ctx.Result(200, reply)
	}
}

type AuthHTTPClient interface {
	GetRolesForUser(ctx context.Context, req *GetRolesForUserRequest, opts ...http.CallOption) (rsp *GetRolesForUserReply, err error)
}

type AuthHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthHTTPClient(client *http.Client) AuthHTTPClient {
	return &AuthHTTPClientImpl{client}
}

func (c *AuthHTTPClientImpl) GetRolesForUser(ctx context.Context, in *GetRolesForUserRequest, opts ...http.CallOption) (*GetRolesForUserReply, error) {
	var out GetRolesForUserReply
	pattern := "/roles/forUser/{user}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/services.auth.v1.Auth/GetRolesForUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
