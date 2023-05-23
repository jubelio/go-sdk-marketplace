package lazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// The Auth Service deals with doing the OAuth flow and exchanging codes for tokens.
// It also lets you refresh tokens in order to get new credentials.
type AuthService interface {
	GetAuthURL(state string, redirectURL string) string
	GetAccessToken(ctx context.Context, code string) (*Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (*Token, error)
}

type AuthServiceOp struct {
	sdk *LazadaClient
}

// Token is the data returned when doing an Oauth Flow through the open platform
type Token struct {
	AccessToken      string `json:"access_token"`
	Country          string `json:"country"`
	RefreshToken     string `json:"refresh_token"`
	AccountID        string `json:"account_id"`
	Code             string `json:"code"`
	AccountPlatform  string `json:"account_platform"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	CountryUserInfo  []struct {
		Country   string `json:"country"`
		UserID    string `json:"user_id"`
		SellerID  string `json:"seller_id"`
		ShortCode string `json:"short_code"`
	} `json:"country_user_info"`
	ExpiresIn   int    `json:"expires_in"`
	RequestID   string `json:"request_id"`
	Account     string `json:"account"`
	retrievedAt time.Time
}

// ExpiresAt tells you the point in time when this token will expire
func (t *Token) ExpiresAt() time.Time {
	return t.calculateExpires(t.ExpiresIn)
}

// RefreshExpiresAt tells you the point in time when the refresh token will expire
func (t *Token) RefreshExpiresAt() time.Time {
	return t.calculateExpires(t.RefreshExpiresIn)
}

// Valid tells you if the token is valid right now
func (t *Token) Valid() bool {
	if t.AccessToken == "" || time.Now().After(t.ExpiresAt()) {
		return false
	}
	return true
}

func (t *Token) calculateExpires(exp int) time.Time {
	return t.retrievedAt.Add(time.Second * time.Duration(exp))
}

// AuthURL returns the URL you should use to start the OAuth flow
// It takes in the URL that the user should be returned to as redirect
// and a state variable which should be a random string
func (c *AuthServiceOp) GetAuthURL(state string, redirectURL string) string {
	baseURL, _ := url.Parse("https://auth.lazada.com/oauth/authorize")

	q := baseURL.Query()
	q.Set("client_id", c.sdk.APIKey)
	q.Set("redirect_uri", redirectURL)
	q.Set("response_type", "code")
	q.Set("state", state)

	baseURL.RawQuery = q.Encode()
	fmt.Println(baseURL.String())
	return baseURL.String()
}

// Exchange sends the received oauth code to the open platform and returns a token
func (c *AuthServiceOp) GetAccessToken(ctx context.Context, code string) (*Token, error) {
	resp, err := c.sdk.Client.R().
		SetResult(Token{}).
		SetQueryParam("code", code).
		Get(ApiNames["AccessToken"])

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, err
	}

	tokenResponse := &Token{}
	err = json.Unmarshal(resp.Body(), tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return tokenResponse, nil
}

func (c *AuthServiceOp) RefreshToken(ctx context.Context, refreshToken string) (*Token, error) {
	resp, err := c.sdk.Client.R().
		SetResult(Token{}).
		SetQueryParam("refresh_token", refreshToken).
		Get(ApiNames["RefreshToken"])

	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("API request error: %v", resp.Status())
	}

	tokenResponse := &Token{}
	err = json.Unmarshal(resp.Body(), tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return tokenResponse, nil
}
