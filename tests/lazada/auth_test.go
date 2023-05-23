package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/apsyadira-jubelio/go-marketplace-sdk/lazada"
	"github.com/jarcoal/httpmock"
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
	token, err := client.Auth.GetAccessToken(context.Background(), "0_117532_85Q0IK2Q4nw4mcXAk4k29RCe36621")

	// Check that there was no error
	assert.Nil(t, err)
	// Check that tokenResponse is not nil

	if token != nil {
		assert.Equal(t, "access_token", token.AccessToken)
	}
}

func Test_RefreshAccessToken(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		lazada.ApiNames["RefreshToken"],
		httpmock.NewBytesResponder(200, loadFixture("access_token.json")),
	)

	// Call GetToken
	token, err := client.Auth.RefreshToken(context.Background(), "500016000300bwa2WteaQyfwBMnPxurcA0mXGhQdTt18356663CfcDTYpWoi")

	// Check that there was no error
	assert.Nil(t, err)
	// Check that tokenResponse is not nil

	if token != nil {
		assert.Equal(t, "access_token", token.AccessToken)
	}
}
