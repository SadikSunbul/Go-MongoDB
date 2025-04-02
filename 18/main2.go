package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Age     int                `bson:"age"`
	Hobbies []string           `bson:"hobbies"`
	Scores  []int              `bson:"scores"`
}

func insertSampleData(coll *mongo.Collection) error {
	people := []interface{}{
		Person{Name: "Sadık", Age: 21, Hobbies: []string{"reading", "gaming", "swimming"}, Scores: []int{85, 90, 75}},
		Person{Name: "Ahmet", Age: 25, Hobbies: []string{"painting", "cooking", "traveling"}, Scores: []int{95, 80, 88}},
	}
	_, err := coll.InsertMany(context.TODO(), people)
	return err
}

func findFirstHobby(coll *mongo.Collection, name string) (string, error) {
	filter := bson.D{{"name", name}}
	var result Person
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return "", err
	}
	if len(result.Hobbies) > 0 {
		return result.Hobbies[0], nil
	}
	return "", fmt.Errorf("Hobi bulunamadı")
}

func hasHobby(coll *mongo.Collection, hobby string) ([]Person, error) {
	filter := bson.D{{"hobbies", hobby}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var results []Person
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func removeHobby(coll *mongo.Collection, hobby string) error {
	filter := bson.D{{"hobbies", hobby}}
	update := bson.D{{"$pull", bson.D{{"hobbies", hobby}}}}
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	return err
}

func updateScore(coll *mongo.Collection, oldScore, newScore int) error {
	filter := bson.D{{"scores", oldScore}}
	update := bson.D{{"$set", bson.D{{"scores.$", newScore}}}}
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	return err
}

func updateHighScores(coll *mongo.Collection) error {
	filter := bson.D{{"scores", bson.D{{"$gt", 80}}}}
	update := bson.D{{"$set", bson.D{{"scores.$[elem]", 100}}}}
	opts := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.D{{"elem", bson.D{{"$gt", 80}}}}},
	})
	_, err := coll.UpdateMany(context.TODO(), filter, update, opts)
	return err
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	coll := client.Database("testdb").Collection("people")

	// Verileri ekle
	if err := insertSampleData(coll); err != nil {
		log.Fatal(err)
	}

	// İlk hobi sorgulama
	hobby, err := findFirstHobby(coll, "Sadık")
	if err != nil {
		log.Printf("Hata: %v", err)
	} else {
		fmt.Printf("Sadık’ın ilk hobisi: %s\n", hobby)
	}

	// Hobi kontrolü
	people, err := hasHobby(coll, "gaming")
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range people {
		fmt.Printf("%s gaming hobisine sahip\n", p.Name)
	}

	// Hobi silme
	if err := removeHobby(coll, "gaming"); err != nil {
		log.Fatal(err)
	}

	// Skor güncelleme
	if err := updateScore(coll, 90, 95); err != nil {
		log.Fatal(err)
	}

	// Yüksek skorları güncelleme
	if err := updateHighScores(coll); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tüm işlemler tamamlandı!")
}
