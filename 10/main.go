package main

import (
	"context"
	"fmt"
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
	coll := client.Database("deneme1").Collection("insertion")

	// Query

	Query2(coll)

}

type Dish struct {
	Name        string
	Description string
}
type Dish2 struct {
	Name        string
	Description string
	Test        string
}

func Query1(coll *mongo.Collection) {

	opt := options.InsertOne().SetBypassDocumentValidation(false)
	ds := Dish{
		Name:        "Sadaık",
		Description: "saflhjkgahs",
	}

	result, err := coll.InsertOne(context.TODO(), ds, opt)

	if err != nil {
		panic(err)
	}
	fmt.Println("ıds:", result.InsertedID)
}

func Query2(coll *mongo.Collection) {
	opt := options.InsertOne().SetBypassDocumentValidation(false)
	ds := Dish2{
		Name:        "Sadaık",
		Description: "saflhjkgahs",
		Test:        "aslfjkhka",
	}

	result, err := coll.InsertOne(context.TODO(), ds, opt)

	if err != nil {
		panic(err)
	}
	fmt.Println("ıds:", result.InsertedID)
}

/*
SetBypassDocumentValidation(true), InsertOneOptions nesnesinin bir parçasıdır.
Yukarıdaki örnekte, yaş alanı bir integer olmalıydı, ancak bir string olarak girildi. Bu normalde doğrulama hatasına yol açardı. Ancak BypassDocumentValidation: true ile eklenmesine izin verildi.
*/
