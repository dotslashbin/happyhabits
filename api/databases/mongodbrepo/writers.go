package mongodbrepo

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// Create executes a database call to create a record
func (db *MongoDBRepo) Create(inputs interface{}, collectionName string) interface{} {

	godotenv.Load()
	configs, _ := godotenv.Read()
	collection := db.Client.Database(configs["MONGO_DATABASE"]).Collection(collectionName)
	result, err := collection.InsertOne(context.TODO(), inputs)
	if err != nil {
		fmt.Println("Problem in inserting data to mongodb")
		log.Fatal(err)
	}

	return result
}
