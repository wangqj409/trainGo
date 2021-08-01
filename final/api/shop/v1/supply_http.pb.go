// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.3

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

type SupplyHTTPServer interface {
	CreateSupply(context.Context, *CreateSupplyRequest) (*CreateSupplyReply, error)
	DeleteSupply(context.Context, *DeleteSupplyRequest) (*DeleteSupplyReply, error)
	GetSupply(context.Context, *GetSupplyRequest) (*GetSupplyReply, error)
	ListSupply(context.Context, *ListSupplyRequest) (*ListSupplyReply, error)
	UpdateSupply(context.Context, *UpdateSupplyRequest) (*UpdateSupplyReply, error)
}

func RegisterSupplyHTTPServer(s *http.Server, srv SupplyHTTPServer) {
	r := s.Route("/")
	r.POST("/supply", _Supply_CreateSupply0_HTTP_Handler(srv))
	r.POST("/supply", _Supply_UpdateSupply0_HTTP_Handler(srv))
	r.POST("/supply", _Supply_DeleteSupply0_HTTP_Handler(srv))
	r.GET("/supply/{supply_id}", _Supply_GetSupply0_HTTP_Handler(srv))
	r.GET("/supply", _Supply_ListSupply0_HTTP_Handler(srv))
}

func _Supply_CreateSupply0_HTTP_Handler(srv SupplyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSupplyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.v1.Supply/CreateSupply")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSupply(ctx, req.(*CreateSupplyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateSupplyReply)
		return ctx.Result(200, reply)
	}
}

func _Supply_UpdateSupply0_HTTP_Handler(srv SupplyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateSupplyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.v1.Supply/UpdateSupply")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSupply(ctx, req.(*UpdateSupplyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateSupplyReply)
		return ctx.Result(200, reply)
	}
}

func _Supply_DeleteSupply0_HTTP_Handler(srv SupplyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSupplyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.v1.Supply/DeleteSupply")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSupply(ctx, req.(*DeleteSupplyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteSupplyReply)
		return ctx.Result(200, reply)
	}
}

func _Supply_GetSupply0_HTTP_Handler(srv SupplyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSupplyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.v1.Supply/GetSupply")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSupply(ctx, req.(*GetSupplyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSupplyReply)
		return ctx.Result(200, reply)
	}
}

func _Supply_ListSupply0_HTTP_Handler(srv SupplyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSupplyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.v1.Supply/ListSupply")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSupply(ctx, req.(*ListSupplyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSupplyReply)
		return ctx.Result(200, reply)
	}
}

type SupplyHTTPClient interface {
	CreateSupply(ctx context.Context, req *CreateSupplyRequest, opts ...http.CallOption) (rsp *CreateSupplyReply, err error)
	DeleteSupply(ctx context.Context, req *DeleteSupplyRequest, opts ...http.CallOption) (rsp *DeleteSupplyReply, err error)
	GetSupply(ctx context.Context, req *GetSupplyRequest, opts ...http.CallOption) (rsp *GetSupplyReply, err error)
	ListSupply(ctx context.Context, req *ListSupplyRequest, opts ...http.CallOption) (rsp *ListSupplyReply, err error)
	UpdateSupply(ctx context.Context, req *UpdateSupplyRequest, opts ...http.CallOption) (rsp *UpdateSupplyReply, err error)
}

type SupplyHTTPClientImpl struct {
	cc *http.Client
}

func NewSupplyHTTPClient(client *http.Client) SupplyHTTPClient {
	return &SupplyHTTPClientImpl{client}
}

func (c *SupplyHTTPClientImpl) CreateSupply(ctx context.Context, in *CreateSupplyRequest, opts ...http.CallOption) (*CreateSupplyReply, error) {
	var out CreateSupplyReply
	pattern := "/supply"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.shop.v1.Supply/CreateSupply"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SupplyHTTPClientImpl) DeleteSupply(ctx context.Context, in *DeleteSupplyRequest, opts ...http.CallOption) (*DeleteSupplyReply, error) {
	var out DeleteSupplyReply
	pattern := "/supply"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.shop.v1.Supply/DeleteSupply"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SupplyHTTPClientImpl) GetSupply(ctx context.Context, in *GetSupplyRequest, opts ...http.CallOption) (*GetSupplyReply, error) {
	var out GetSupplyReply
	pattern := "/supply/{supply_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.shop.v1.Supply/GetSupply"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SupplyHTTPClientImpl) ListSupply(ctx context.Context, in *ListSupplyRequest, opts ...http.CallOption) (*ListSupplyReply, error) {
	var out ListSupplyReply
	pattern := "/supply"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.shop.v1.Supply/ListSupply"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SupplyHTTPClientImpl) UpdateSupply(ctx context.Context, in *UpdateSupplyRequest, opts ...http.CallOption) (*UpdateSupplyReply, error) {
	var out UpdateSupplyReply
	pattern := "/supply"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.shop.v1.Supply/UpdateSupply"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}