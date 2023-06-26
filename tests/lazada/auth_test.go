package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/jubelio/go-sdk-marketplace/lazada"
	"github.com/stretchr/testify/assert"
)

func Test_GetAuthURL(t *testing.T) {
	setup()
	defer teardown()

	redirectUrl := fmt.Sprintf(os.Getenv("API_STAGING")+"/marketplace/lazada/connect?tenantId=%s&companyId=%d", "sdljkaslj234j2k4j2", 129929)
	authURL := client.Auth.GetAuthURL("", redirectUrl)
	t.Logf("auth url: %s", authURL)
}

func Test_GetAccessToken(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		lazada.ApiNames["AccessToken"],
		httpmock.NewBytesResponder(200, loadFixture("access_token.json")),
	)

	// Call GetToken
	token, err := client.Auth.GetAccessToken(context.Background(), "0_117532_PiJJlotSBDPKwsoi4s2Jc13E38014")

	// Check that there was no error
	assert.Nil(t, err)

	if token != nil {
		assert.Equal(t, "50000600000218479995AvcwUvCeyufRljGpxnPZZ3gxClvEWudZEBthstUIT4si", token.AccessToken)
	}
}
