package lazada

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"sort"
	"strings"

	"github.com/go-resty/resty/v2"
)

type LazadaClient struct {
	AppKey      string
	AppSecret   string
	AccessToken string
	Region      string
	Client      *resty.Client
	APIKey      string
	APISecret   string
	RedirectURL string `env:"LAZADA_REDIRECT_URL"`
	BaseURL     string

	Auth *AuthServiceOp
	Chat *ChatServiceOp
}

type Option func(c *LazadaClient)

func NewClient(appKey, secret string, region Region) *LazadaClient {
	baseURL := endpoints[Region(region)]
	lazadaClient := &LazadaClient{
		BaseURL:   baseURL,
		APIKey:    appKey,
		APISecret: secret,
	}

	lazadaClient.Auth = &AuthServiceOp{sdk: lazadaClient}
	lazadaClient.Chat = &ChatServiceOp{sdk: lazadaClient}

	lazadaClient.WrapperApi()

	return lazadaClient
}

// NewTokenClient takes a client access token and returns a copy of the client with the token set.
func (c *LazadaClient) SetTokenClient(token string) *LazadaClient {
	newC := *c
	newC.AccessToken = token
	return &newC
}

func (lc *LazadaClient) WrapperApi() *LazadaClient {
	rc := GetClient(lc.BaseURL, *lc)
	lc.Client = rc
	return lc
}

// Signature calculates the signature for the query parameter needed with every request
func (c *LazadaClient) Signature(api string, val url.Values) string {
	var buf bytes.Buffer

	buf.WriteString(api)
	keys := make([]string, 0, len(val))
	for k := range val {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		vs := val[k]
		keyEscaped := url.QueryEscape(k)

		for _, v := range vs {
			buf.WriteString(keyEscaped)
			buf.WriteString(v)
		}
	}

	signer := hmac.New(sha256.New, []byte(c.APISecret))
	signer.Write(buf.Bytes())
	sig := signer.Sum(nil)

	return strings.ToUpper(hex.EncodeToString(sig))
}

func (c *LazadaClient) Do(ctx context.Context, method string, url string, body interface{}, res interface{}) (*resty.Response, error) {
	request := c.Client.R().
		SetContext(ctx).
		SetResult(res)

	if body != nil {
		request = request.SetBody(body)
	}

	resp, err := request.Execute(method, url)

	if err != nil {
		log.Fatalf("ERROR %v", err)
		return nil, err
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}

	return resp, nil
}
