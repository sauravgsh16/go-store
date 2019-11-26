package service

import (
	"context"

	"github.com/segmentio/ksuid"

	"github.com/sauravgsh16/go-store/account/domain/account"
)

// Service interface
type Service interface {
	GetAccount(ctx context.Context, id string) (*account.Account, error)
	GetAccounts(ctx context.Context, skip, take uint64) ([]account.Account, error)
	PostAccount(ctx context.Context, name string) (*account.Account, error)
}

// AccService service struct
type AccService struct {
	repo account.Repository
}

// NewAccService returns a new account service struct
func NewAccService(r account.Repository) Service {
	return &AccService{
		repo: r,
	}
}

// GetAccount returns an account, given an id
func (as *AccService) GetAccount(ctx context.Context, id string) (*account.Account, error) {
	acc, err := as.repo.Select(ctx, id)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

// GetAccounts returns all accounts
func (as *AccService) GetAccounts(ctx context.Context, skip, take uint64) ([]account.Account, error) {
	if take > 50 || (skip == 0 && take == 0) {
		take = 50
	}
	return as.repo.SelectAll(ctx, skip, take)
}

// PostAccount creates an account
func (as *AccService) PostAccount(ctx context.Context, name string) (*account.Account, error) {
	a := &account.Account{
		Name: name,
		ID:   ksuid.New().String(),
	}
	if err := as.repo.Insert(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}
