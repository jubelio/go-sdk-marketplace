package shopee

import (
	"net/http"
	"net/url"
)

// Option is used to configure client with options
type Option func(c *ShopeeClient)

func WithRetry(retries int) Option {
	return func(c *ShopeeClient) {
		c.retries = retries
	}
}

func WithLogger(logger LeveledLoggerInterface) Option {
	return func(c *ShopeeClient) {
		c.log = logger
	}
}

func WithProxy(proxyHost string) Option {
	return func(c *ShopeeClient) {
		proxyURL, err := url.Parse(proxyHost)
		if err != nil {
			return
		}
		c.Client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
}
