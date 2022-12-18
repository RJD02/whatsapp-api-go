package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ImageModel struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
}

type PostData struct {
	Messaging_product string       `json:"messaging_product"`
	To                string       `json:"to"`
	Type              string       `json:"type"`
	Image             ImageModel   `json:"image"`
	Headers           HeadersModel `json:"headers"`
}

type HeadersModel struct {
	Content_Type  string `json:"Content-Type"`
	Authorization string `json:"Authorization"`
}

type PostRequest struct {
	Data    PostData     `json:"data"`
	Headers HeadersModel `json:"headers"`
}

type WhatsappObject struct {
	Object string `json:"object"`
	Entry  []struct {
		ID      string `json:"id"`
		Changes []struct {
			Value struct {
				MessagingProduct string `json:"messaging_product"`
				Metadata         struct {
					DisplayPhoneNumber string `json:"display_phone_number"`
					PhoneNumberID      string `json:"phone_number_id"`
				} `json:"metadata"`
				Contacts []struct {
					Profile struct {
						Name string `json:"name"`
					} `json:"profile"`
					WaID string `json:"wa_id"`
				} `json:"contacts"`
				Messages []struct {
					From      string `json:"from"`
					ID        string `json:"id"`
					Timestamp string `json:"timestamp"`
					Text      struct {
						Body string `json:"body"`
					} `json:"text"`
					Type string `json:"type"`
				} `json:"messages"`
			} `json:"value"`
			Field string `json:"field"`
		} `json:"changes"`
	} `json:"entry"`
}

type Voter struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	WardNo       string             `json:"ward_no" bson:"Ward_no"`
	SLNO         string             `json:"slno" bson:"SLNO"`
	HouseNo      string             `json:"houseno" bson:"houseno"`
	VnameEnglish string             `json:"VNAME_ENGLISH" bson:"VNAME_ENGLISH"`
	Cardno       string             `json:"cardno" bson:"cardno"`
	Age          int                `json:"Age" bson:"Age"`
}
