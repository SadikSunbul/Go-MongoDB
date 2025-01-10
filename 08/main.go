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
	coll := client.Database("deneme1").Collection("courses2")

	// Query

	Query3(coll)

}

type Course struct {
	Title      string `bson:"title,omitempty"`
	CourseId   string `bson:"course_id,omitempty"`
	Enrollment int32  `bson:"enrollment,omitempty"`
}

func Query1(coll *mongo.Collection) {
	docs := []interface{}{
		Course{Title: "Primate Behavior", CourseId: "PSY2030", Enrollment: 40},
		Course{Title: "Revolution and Reform", CourseId: "HIST3080", Enrollment: 12},
	}

	result, err := coll.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}
	fmt.Println("ıds:", result.InsertedIDs)
}

func Query2(coll *mongo.Collection) {
	filter := bson.D{}
	// Hangi alanların dönüceğini beliritir
	opts := options.Find().SetProjection(bson.D{{"course_id", 0}, {"enrollment", 0}})

	cursor, err := coll.Find(context.TODO(), filter, opts)
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

func Query3(coll *mongo.Collection) {

	projectStage := bson.D{{"$project", bson.D{{"title", 1}}}}

	filter := bson.D{
		{"$match", bson.D{
			{"enrollment", bson.D{{"$gt", 20}}}, // Örneğin, sadece "status: active" olan belgeler
		}},
	}
	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{filter, projectStage})
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
