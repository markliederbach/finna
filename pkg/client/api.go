package client

import (
	"context"

	"github.com/plaid/plaid-go/v3/plaid"
)

type PlaidSdkClient interface {
	ItemPublicTokenExchange(ctx context.Context) plaid.ApiItemPublicTokenExchangeRequest
}

type PlaidClient interface {
	GetAccessToken(ctx context.Context, publicToken string) (string, string, error)
}
