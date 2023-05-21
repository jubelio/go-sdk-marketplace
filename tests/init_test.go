package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/apsyadira-jubelio/go-marketplace-sdk/shopee"
	"github.com/caarlos0/env"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

const (
	maxRetries  = 3
	shopID      = 1234567
	merchantID  = 0
	accessToken = "accesstoken"
)

var (
	client *shopee.ShopeeClient
	app    shopee.AppConfig
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file")
		app = shopee.AppConfig{
			PartnerID:   12345678,
			PartnerKey:  "hush",
			RedirectURL: "https://example.com/callback",
			APIURL:      "https://partner.test-stable.shopeemobile.com",
		}
	} else {
		env.Parse(&app)
	}
	client = shopee.NewClient(app,
		shopee.WithRetry(maxRetries))
	httpmock.ActivateNonDefault(client.Client)
}

func teardown() {
	httpmock.DeactivateAndReset()
}

func loadFixture(filename string) []byte {
	f, err := ioutil.ReadFile("../mockdata/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}
	return f
}

func loadMockData(filename string, out interface{}) {
	f, err := ioutil.ReadFile("../mockdata/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}
	if err := json.Unmarshal(f, &out); err != nil {
		panic(fmt.Sprintf("decode mock data error: %s", err))
	}
}
