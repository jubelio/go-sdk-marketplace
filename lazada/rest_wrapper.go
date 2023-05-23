package lazada

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

func GetClient(URL string, lc LazadaClient) *resty.Client {
	client := resty.New().
		SetBaseURL(lc.BaseURL)

	client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		// Before request middleware
		c.SetQueryParam("sign_method", "sha256")
		c.SetQueryParam("timestamp", fmt.Sprintf("%d", time.Now().Unix()*1000))
		c.SetQueryParam("app_key", lc.AppKey)

		if lc.AccessToken != "" {
			c.SetQueryParam("access_token", lc.AccessToken)
		}

		sign := lc.Signature(strings.TrimPrefix(r.URL, "/rest"), c.QueryParam)
		c.SetQueryParam("sign", sign)
		return nil
	})

	return client
}
