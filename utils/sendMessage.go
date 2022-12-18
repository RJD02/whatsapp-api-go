package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/RJD02/whatsapp-elections-go/config"
	"github.com/RJD02/whatsapp-elections-go/model"
)

var configurations map[string]string = config.GetConfig()

func SendMessage(wg *sync.WaitGroup, phoneNumberId, from, msgBody string) {

	fmt.Println("Sending message")
	defer wg.Done()
	headers := model.HeadersModel{
		Content_Type:  "application/json",
		Authorization: "Bearer " + configurations["WHATSAPP_TOKEN"],
	}
	image := model.ImageModel{
		Link:    "https://images.unsplash.com/photo-1670031652376-e2b853e67390?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyfHx8ZW58MHx8fHw%3D&auto=format&fit=crop&w=500&q=60",
		Caption: "Ack: " + msgBody,
	}
	data := model.PostData{
		Messaging_product: "whatsapp",
		To:                from,
		Type:              "image",
		Image:             image,
		Headers:           headers,
	}

	Url := "https://graph.facebook.com/v15.0/" + configurations["WHATSAPP_PHONE_NUMBER_ID"] + "/messages?access_token=EAAJkuqEv6zABAAr7L0MSepSNJdcpkKUTsFn9PZCKprB4DWwv67vf9TZCmVVmRD5d8t7DIApcZA9jzZC0NyjkL2GtbfjroRhcG6DMrmRfQzivmoZCgHVnehCPnzP6u006jCRWXSqBcLuCqOGiLr7pfIuMyGRsrLTg13Q0AYIdVZCfKTX3vcihY3DyISqtUohwkZD"

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(data)

	resp, err := http.Post(Url, headers.Content_Type, payload)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		jsonStr := string(body)
		fmt.Println("Response:", jsonStr)
	} else {
		fmt.Println("Get failed with error", resp.StatusCode)
	}
}
