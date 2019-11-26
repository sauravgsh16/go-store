package client

import (
	"context"
	"google.golang.org/grpc"

	"github.com/sauravgsh16/go-store/account/pb"
)

// Client struct
type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

// NewClient return new client
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewAccountServiceClient(conn)
	return &Client{
		conn:    conn,
		service: c,
	}, nil
}

// Close closes client connection
func (c *Client) Close() {
	c.conn.Close()
}

// PostAccount call a post account on the client
func (c *Client) PostAccount(ctx context.Context, name string) (*pb.Account, error) {
	req := &pb.PostAccountRequest{
		Name: name,
	}

	resp, err := c.service.PostAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Account, nil
}

// GetAccount calls a get account on the client
func (c *Client) GetAccount(ctx context.Context, id string) (*pb.Account, error) {
	req := &pb.GetAccountRequest{
		Id: id,
	}

	resp, err := c.service.GetAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Account, nil
}

// GetAccounts call a get accounts on the client
func (c *Client) GetAccounts(ctx context.Context, skip, take uint64) ([]*pb.Account, error) {
	req := &pb.GetAccountsRequest{
		Skip: skip,
		Take: take,
	}
	resp, err := c.service.GetAccounts(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Accounts, nil
}
