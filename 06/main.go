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
	coll := client.Database("deneme1").Collection("courses")

	// Query

	Query3(coll)

}

type Course struct {
	Title      string
	Enrollment int32
}

func Query1(coll *mongo.Collection) {
	docs := []interface{}{
		Course{Title: "World Fiction", Enrollment: 35},
		Course{Title: "Abstract Algebra", Enrollment: 60},
		Course{Title: "Modern Poetry", Enrollment: 12},
		Course{Title: "Plate Tectonics", Enrollment: 35},
	}

	result, err := coll.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	fmt.Println("ıds:", result.InsertedIDs)
}

func Query2(coll *mongo.Collection) {

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"enrollment", 1}}) // -1 de azalan yapar

	cursor, err := coll.Find(context.TODO(), filter, opts)

	if err != nil {
		panic(err)
	}

	var results []Course
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

}

func Query3(coll *mongo.Collection) {

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"enrollment", 1}, {"title", -1}}) // -1 de azalan yapar

	cursor, err := coll.Find(context.TODO(), filter, opts)

	if err != nil {
		panic(err)
	}

	var results []Course
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

}

func Query4(coll *mongo.Collection) {
	sortStage := bson.D{{"$sort", bson.D{{"enrollment", -1}, {"title", 1}}}}

	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{sortStage})
	if err != nil {
		panic(err)
	}

	var results []Course
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}
