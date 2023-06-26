package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jubelio/go-sdk-marketplace/lazada"
)

func main() {
	lazadaClient := lazada.NewClient("1111222334455", "JJJJJJJJJJJJJJJJ@@@@@JJSJJS", lazada.Indonesia)
	lazadaClient.NewTokenClient("1111111128378273817238723817DJHSKJHD899787389728")

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
