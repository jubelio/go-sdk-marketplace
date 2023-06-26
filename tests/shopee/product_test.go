package tests

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/jubelio/go-sdk-marketplace/shopee"
)

func Test_GetProduct(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_base_info", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_product_resp.json")))

	res, err := client.Product.GetProductById(123456, accessToken, shopee.GetProductParamRequest{
		ItemIDList:          []int{3400133011},
		NeedTaxInfo:         true,
		NeedComplaintPolicy: true,
	})

	if err != nil {
		t.Errorf("Product.GetProductResponse error: %s", err)
	}

	t.Logf("return tok: %#v", res)

	var expectedMsgID int64 = 3400133011
	if res.Response.ItemList[0].ItemID != expectedMsgID {
		t.Errorf("MessageList.MessageID returned %+v, expected %+v", res.Response.ItemList[0].ItemID, expectedMsgID)
	}
}
