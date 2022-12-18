package db

import (
	"context"
	"fmt"
	"log"

	"github.com/RJD02/whatsapp-elections-go/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "whatsappIntegration"
const collName = "voters"

var configurations = config.GetConfig()
var MONGODB_USERNAME = configurations["MONGODB_USERNAME"]
var MONGODB_PASSWORD = configurations["MONGODB_PASSWORD"]
var connectionString = "mongodb+srv://admin-" + MONGODB_USERNAME + ":" + MONGODB_PASSWORD + "@cluster0.lkxsz.mongodb.net/?retryWrites=true&w=majority"

func GetCollection() *mongo.Collection {
	fmt.Println(connectionString)
	var collection *mongo.Collection
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	dbClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	collection = dbClient.Database(dbName).Collection(collName)
	fmt.Println("Collection instance is ready")

	return collection
}
