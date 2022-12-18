package router

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/RJD02/whatsapp-elections-go/db"
	"github.com/RJD02/whatsapp-elections-go/model"
	"github.com/RJD02/whatsapp-elections-go/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = db.GetCollection()

const VERIFY_TOKEN = "helloworldthisiswhatsappelectionswebhookintesting"

func WebhookGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	fmt.Println("Get request for /webhook")
	result, _ := url.Parse(r.RequestURI)
	qparams := result.Query()
	mode := qparams["hub.mode"][0]
	token := qparams["hub.verify_token"][0]
	challenge := qparams["hub.challenge"][0]

	if mode != "" && token != "" {
		if mode == "subscribe" && token == VERIFY_TOKEN {
			fmt.Println("WEBHOOK VERIFIED")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(challenge)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
	defer r.Body.Close()
}

func WebHookPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	var wg = &sync.WaitGroup{}
	fmt.Println("POST request for /webhook")
	result, _ := url.Parse(r.RequestURI)
	qparams := result.Query()
	challenge := qparams["hub.challenge"]
	fmt.Println(r.Body, challenge)
	var jsonObj model.WhatsappObject
	err := json.NewDecoder(r.Body).Decode(&jsonObj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Request Object", jsonObj)

	if jsonObj.Entry[0].Changes[0].Value.Messages != nil {
		phoneNumberID := jsonObj.Entry[0].Changes[0].Value.Metadata.PhoneNumberID
		from := jsonObj.Entry[0].Changes[0].Value.Messages[0].From
		msgBody := ""
		msgBody = jsonObj.Entry[0].Changes[0].Value.Messages[0].Text.Body

		filter := bson.M{"cardno": msgBody}
		res := collection.FindOne(context.Background(), filter)

		var voter model.Voter
		if err := res.Decode(&voter); err != nil {
			fmt.Println(err)
			fmt.Println("Voter not found")
			msgBody = "Voter not found"
		} else {
			fmt.Println("Voter found", voter)
			msgBody = "Here are your details:\nWard no: " + voter.WardNo + "\nCard no: " + voter.Cardno + "\nHouse no: " + voter.HouseNo + "\nSLNO: " + voter.SLNO + "\nVname English: " + voter.VnameEnglish + "\nAge: " + strconv.FormatInt(int64(voter.Age), 10)
		}

		wg.Add(1)
		go utils.SendMessage(wg, phoneNumberID, from, msgBody)

	}
	wg.Wait()
	r.Body.Close()
}

func WebHookRoutes(s *mux.Router) {
	fmt.Println("Subrouting enabled")
	s.HandleFunc("/", WebhookGet).Methods("GET")
	s.HandleFunc("/", WebHookPost).Methods("POST")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requesting home route")
}

func Router() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/webhook").Subrouter()
	WebHookRoutes(s)
	// s.HandleFunc("/", WebhookGet).Methods("GET")

	r.HandleFunc("/", serveHome)

	return r
}
