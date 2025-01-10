package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
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
	Query3(coll)
}

///////// Aggregation ////////////

type Tea struct {
	Type     string
	Category string
	Toppings []string
	Price    float32
}

func Query1(coll *mongo.Collection) {
	docs := []interface{}{
		Tea{Type: "Masala", Category: "black", Toppings: []string{"ginger", "pumpkin spice", "cinnamon"}, Price: 6.75},
		Tea{Type: "Gyokuro", Category: "green", Toppings: []string{"berries", "milk foam"}, Price: 5.65},
		Tea{Type: "English Breakfast", Category: "black", Toppings: []string{"whipped cream", "honey"}, Price: 5.75},
		Tea{Type: "Sencha", Category: "green", Toppings: []string{"lemon", "whipped cream"}, Price: 5.15},
		Tea{Type: "Assam", Category: "black", Toppings: []string{"milk foam", "honey", "berries"}, Price: 5.65},
		Tea{Type: "Matcha", Category: "green", Toppings: []string{"whipped cream", "honey"}, Price: 6.45},
		Tea{Type: "Earl Grey", Category: "black", Toppings: []string{"milk foam", "pumpkin spice"}, Price: 6.15},
		Tea{Type: "Hojicha", Category: "green", Toppings: []string{"lemon", "ginger", "milk foam"}, Price: 5.55},
	}

	result, err := coll.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	fmt.Println("ıds:", result.InsertedIDs)
}

// Sınırlamalar

func Query2(coll *mongo.Collection) {
	// Aşağıdaki örnekte her çay kategorisi için ortalama puan ve puan sayısı hesaplanmakta ve gösterilmektedir.

	// create group stage : Toplama hattı, belgeleri kategori alanına göre gruplandırmak için $group aşamasını kullanır, $avg ifade operatörünü kullanarak ortalamayı hesaplar ve $sum ifade operatörünü kullanarak belge sayısını sayar.
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$category"},
			{"average_price", bson.D{{"$avg", "$price"}}}, // price: Sadece bir kelime gibi görünür ve MongoDB bunun bir alan olduğunu anlayamaz.  $price: Bu, MongoDB'ye "Bu, bir dokümandaki price alanıdır" demektir.
			{"type_total", bson.D{{"$sum", 1}}},
		}}}

	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	if err != nil {
		panic(err)
	}

	// display the results
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("Average price of %v tea options: $%v \n", result["_id"], result["average_price"])
		fmt.Printf("Number of %v tea options: %v \n\n", result["_id"], result["type_total"])
	}
}

// Sonuçlardaki Alanları Atla
func Query3(coll *mongo.Collection) {
	/*
			Toplama işlem hattı aşağıdaki aşamaları içerir:
		$match aşaması toppings alanının "süt köpüğü" içerdiği belgeleri eşleştirmek için
		$unset aşaması _id ve kategori alanlarını atlamak için
		$sort aşaması fiyat ve toppings'i artan sırada sıralamak için
		$limit aşaması ilk iki belgeyi göstermek için
	*/
	// create the stages
	matchStage := bson.D{{"$match", bson.D{{"toppings", "milk foam"}}}}   // filtreleme yaptık
	unsetStage := bson.D{{"$unset", bson.A{"_id", "category"}}}           // belşirli alanları çıakrtık
	sortStage := bson.D{{"$sort", bson.D{{"price", 1}, {"toppings", 1}}}} // sıralama yaptık
	limitStage := bson.D{{"$limit", 2}}                                   // ilk iki belgeyi göstermek icin

	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{matchStage, unsetStage, sortStage, limitStage})

	if err != nil {
		panic(err)
	}

	// display the results
	var results []Tea
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("Tea: %v \nToppings: %v \nPrice: $%v \n\n", result.Type, strings.Join(result.Toppings, ", "), result.Price)
	}

}
