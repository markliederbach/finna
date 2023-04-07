package client

import (
	"context"

	"github.com/plaid/plaid-go/v3/plaid"
)

type DefaultPlaidClient struct {
	Client *plaid.PlaidApiService
	Config *plaid.Configuration
	// accessToken is stored in memory for this example,
	// but should be stored in a secure persistent data store for production
	accessToken string
	itemId      string
}

type NewPlaidClientInput struct {
	ClientID    string
	Secret      string
	Environment string
}

// ToPlaidEnvironment converts a string to a plaid.Environment URL
func ToPlaidEnvironment(environment string) plaid.Environment {
	pladEnvMap := map[string]plaid.Environment{
		"sandbox":     plaid.Sandbox,
		"development": plaid.Development,
		"production":  plaid.Production,
	}
	return pladEnvMap[environment]
}

// NewPlaidClient creates a new Plaid client with the given credentials
func NewPlaidClient(input NewPlaidClientInput) PlaidClient {
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", input.ClientID)
	configuration.AddDefaultHeader("PLAID-SECRET", input.Secret)
	configuration.UseEnvironment(ToPlaidEnvironment(input.Environment))
	client := plaid.NewAPIClient(configuration)
	return &DefaultPlaidClient{
		Client: client.PlaidApi,
		Config: client.GetConfig(),
	}
}

// GetAccessToken returns the access token for the client from cache, or
// fetches it from the Plaid API if it is not cached
func (c *DefaultPlaidClient) GetAccessToken(ctx context.Context, publicToken string) (string, string, error) {
	if c.accessToken == "" || c.itemId == "" {
		accessToken, itemId, err := c.exchangePublicToken(ctx, publicToken)
		if err != nil {
			return "", "", err
		}
		c.accessToken = accessToken
		c.itemId = itemId
	}
	return c.accessToken, c.itemId, nil
}

func (c *DefaultPlaidClient) exchangePublicToken(ctx context.Context, publicToken string) (string, string, error) {
	exchangePublicTokenResp, _, err := c.Client.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken),
	).Execute()
	if err != nil {
		return "", "", err
	}
	return exchangePublicTokenResp.GetAccessToken(), exchangePublicTokenResp.GetItemId(), nil
}
