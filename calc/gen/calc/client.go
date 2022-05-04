// Code generated by goa v3.7.3, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen calc/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "calc" service client.
type Client struct {
	AddEndpoint    goa.Endpoint
	SubEndpoint    goa.Endpoint
	MulEndpoint    goa.Endpoint
	DivideEndpoint goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(add, sub, mul, divide goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:    add,
		SubEndpoint:    sub,
		MulEndpoint:    mul,
		DivideEndpoint: divide,
	}
}

// Add calls the "add" endpoint of the "calc" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Sub calls the "sub" endpoint of the "calc" service.
func (c *Client) Sub(ctx context.Context, p *SubPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.SubEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Mul calls the "mul" endpoint of the "calc" service.
func (c *Client) Mul(ctx context.Context, p *MulPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.MulEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Divide calls the "divide" endpoint of the "calc" service.
// Divide may return the following errors:
//	- "DivByZero" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Divide(ctx context.Context, p *DividePayload) (res int, err error) {
	var ires interface{}
	ires, err = c.DivideEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}
