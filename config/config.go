package config

import (
	"os"
)

func GetConfig() map[string]string {

	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	config := make(map[string]string)

	config["MONGODB_USERNAME"] = os.Getenv("MONGODB_USERNAME")
	config["MONGODB_PASSWORD"] = os.Getenv("MONGODB_PASSWORD")
	config["WHATSAPP_TOKEN"] = os.Getenv("WHATSAPP_TOKEN")
	config["WHATSAPP_PHONE_NUMBER_ID"] = os.Getenv("WHATSAPP_PHONE_NUMBER_ID")

	return config
}
