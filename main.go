package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	router "github.com/RJD02/whatsapp-elections-go/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r := router.Router()

	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
