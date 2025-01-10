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
	coll := client.Database("deneme1").Collection("tea2")

	// Query

	Query2(coll)

}

type Course struct {
	Title      string
	Department string
	Enrollment int32
}

func Query1(coll *mongo.Collection) {
	docs := []interface{}{
		Course{Title: "World Fiction", Department: "English", Enrollment: 35},
		Course{Title: "Abstract Algebra", Department: "Mathematics", Enrollment: 60},
		Course{Title: "Modern Poetry", Department: "English", Enrollment: 12},
		Course{Title: "Plate Tectonics", Department: "Earth Science", Enrollment: 30},
	}

	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}

	fmt.Println("ıds:", result.InsertedIDs)
}

func Query2(coll *mongo.Collection) {
	results, err := coll.Distinct(context.TODO(), "department", bson.D{{"enrollment", bson.D{{"$lt", 50}}}})
	// enrollment 50 den kucuk ve department degerı farklı olanları alır
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
