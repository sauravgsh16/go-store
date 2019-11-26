package controller

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sauravgsh16/go-store/account/pb"
	"github.com/sauravgsh16/go-store/account/service"
)

// AccGrpc struct
type AccGrpc struct {
	Serv service.Service
}

// GetAccount to get account
func (g *AccGrpc) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	resp, err := g.Serv.GetAccount(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error getting account: %s", err.Error()))
	}
	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:   resp.ID,
			Name: resp.Name,
		},
	}, nil
}

// GetAccounts to get accounts
func (g *AccGrpc) GetAccounts(ctx context.Context, req *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	resp, err := g.Serv.GetAccounts(ctx, req.Skip, req.Take)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error getting all accounts: %s", err.Error()))
	}
	accs := []*pb.Account{}
	for _, acc := range resp {
		accs = append(
			accs,
			&pb.Account{
				Id:   acc.ID,
				Name: acc.Name,
			},
		)
	}
	return &pb.GetAccountsResponse{
		Accounts: accs,
	}, nil
}

// PostAccount to create an account
func (g *AccGrpc) PostAccount(ctx context.Context, req *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	resp, err := g.Serv.PostAccount(ctx, req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Errors created account: %s", err.Error()))
	}
	return &pb.PostAccountResponse{
		Account: &pb.Account{
			Id:   resp.ID,
			Name: resp.Name,
		},
	}, nil
}
