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

	Query4(coll)

}

type Course struct {
	Title      string
	Enrollment int32
}

func Query1(coll *mongo.Collection) {
	filter := bson.D{{"enrollment", bson.D{{"$gt", 20}}}}
	opts := options.Find().SetLimit(2)

	cursor, err := coll.Find(context.TODO(), filter, opts)

	var results []Course
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}

func Query2(coll *mongo.Collection) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"enrollment", 1}}).SetSkip(1).SetLimit(2)

	cursor, err := coll.Find(context.TODO(), filter, opts)

	var results []Course
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}

func Query4(coll *mongo.Collection) {
	limitStage := bson.D{{"$limit", 3}}

	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{limitStage})
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
