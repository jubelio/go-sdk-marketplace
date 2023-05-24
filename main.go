package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/apsyadira-jubelio/go-marketplace-sdk/lazada"
)

func main() {
	lazadaClient := lazada.NewClient("117532", "G5VwB0wyhk3XQEsklCfmHSF2kP2luEqS", lazada.Indonesia)
	lazadaClient.NewTokenClient("50000600116cWYzTphDtTDshMBux1993574eoq9YzkugHtfWTiXeDQ7OzvDLRkFx")

	res, err := lazadaClient.Chat.SendMessage(context.Background(), &lazada.SendMessageParams{
		SessionID:  "400096074640_1_400601424036_2_103",
		TemplateID: lazada.NormalTextMessage,
		Txt:        "Testing Lazada SDK jubelio chat",
	})

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	jsonData, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
