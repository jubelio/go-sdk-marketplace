package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/apsyadira-jubelio/go-marketplace-sdk/lazada"
	"github.com/jarcoal/httpmock"
)

var (
	client *lazada.Client
)

func setup() {
	client = lazada.NewClient("1999192", "JSHDjhdjhhjhdjshdjshdjhsjdhsjhd1982781", lazada.Indonesia)

	httpmock.ActivateNonDefault(client.Client)
}

func teardown() {
	httpmock.DeactivateAndReset()
}

func loadFixture(filename string) []byte {
	f, err := ioutil.ReadFile("../../mockdata/lazada/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}

	return f
}

func loadMockData(filename string, out interface{}) {
	f, err := ioutil.ReadFile("../../mockdata/lazada/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}
	if err := json.Unmarshal(f, &out); err != nil {
		panic(fmt.Sprintf("decode mock data error: %s", err))
	}
}
