package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodbConnectionString = "mongodb://admin:password@localhost:27017"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(mongodbConnectionString))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("deneme1").Collection("Plant")

	// Query
	Query2(coll)
}

type Plant struct {
	Species string
	PlantID int32 `bson:"plant_id"`
	Height  float64
}

func Query1(coll *mongo.Collection) {
	docs := []interface{}{
		Plant{Species: "Polyscias fruticosa", PlantID: 1, Height: 27.6},
		Plant{Species: "Polyscias fruticosa", PlantID: 2, Height: 34.9},
		Plant{Species: "Ledebouria socialis", PlantID: 1, Height: 11.4},
	}
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}

	fmt.Println("ıds:", result.InsertedIDs)
}

/*
Upsert İşlemi
MongoDB'deki upsert (update or insert) işlemi, bir dokümanı şu şekilde ele alır:

Eğer verilen filtreye uygun bir doküman varsa, bu doküman güncellenir.
Eğer filtreye uygun bir doküman yoksa, yeni bir doküman eklenir.
*/

func Query2(coll *mongo.Collection) {
	filter := bson.D{{"species", "Ledebouria socialis"}, {"plant_id", 10}}
	/*
			Türü "Ledebouria socialis" olan,
			plant_id değeri 3 olan bir dokümanı arar.
		Eğer böyle bir doküman varsa, güncellenecek. Yoksa yeni bir doküman oluşturulacak.
	*/
	update := bson.D{{"$set", bson.D{{"species", "Ledebouria socialis"}, {"plant_id", 5}, {"height", 8.3}}}}
	opts := options.Update().SetUpsert(true)
	//	Eğer filtreye uygun bir doküman bulunamazsa, güncelleme işlemi yerine yeni bir doküman ekler.

	result, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of documents updated: %v\n", result.ModifiedCount)
	fmt.Printf("Number of documents upserted: %v\n", result.UpsertedCount)
}
