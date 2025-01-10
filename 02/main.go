package main

import (
	"context"
	"encoding/json"
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
	coll := client.Database("deneme1").Collection("tea")

	// Query

	Query8(coll)

}

type Tea struct {
	Type   string
	Rating int32
	Vendor []string `bson:"vendor,omitempty" json:"vendor,omitempty"`
}

func Query1(coll *mongo.Collection) {
	/*
		Kriterleri gerçek değerlerle eşleştirmek için aşağıdaki formatı kullanın:
		filter := bson.D{{"<field>", "<value>"}}

		Kriterleri bir sorgu operatörüyle eşleştirmek için aşağıdaki biçimi kullanın:
		filter := bson.D{{"<field>", bson.D{{"<operator>", "<value>"}}}}

		Varolmayan Veritabanları ve Koleksiyonlar
		Bir yazma işlemi gerçekleştirdiğinizde gerekli veritabanı ve koleksiyon mevcut değilse, sunucu bunları dolaylı olarak oluşturur.
	*/

	docs := []interface{}{
		Tea{Type: "Masala", Rating: 10, Vendor: []string{"A", "C"}},
		Tea{Type: "English Breakfast", Rating: 6},
		Tea{Type: "Oolong", Rating: 7, Vendor: []string{"C"}},
		Tea{Type: "Assam", Rating: 5},
		Tea{Type: "Earl Grey", Rating: 8, Vendor: []string{"A", "B"}},
	}
	result, err := coll.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	fmt.Println("ıds:", result.InsertedIDs)
}

func Query2(coll *mongo.Collection) {

	filter := bson.D{{"type", "Oolong"}}
	// filter := bson.D{{"type", bson.D{{"$eq", "Oolong"}}}} // buda aynı sonucu verır

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var teas []Tea
	if err = cursor.All(context.TODO(), &teas); err != nil {
		panic(err)
	}
	for _, tea := range teas {
		fmt.Println(tea)
	}
}

func Query3(coll *mongo.Collection) {

	/*
		$gt -> Daha büyük
		$gte -> Daha büyük veya eşit
		$lt -> Daha kucuk
		$lte -> Daha kucuk veya eşit
		$eq -> Esit
		$ne -> Esit degil
		$in -> Belirtilen degerlerden biri
		$nin -> Belirtilen degerlerden hicbiri
		$all -> Belirtilen degerlerden herhangi biri
	*/

	filter := bson.D{{"rating", bson.D{{"$lt", 7}}}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []Tea
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		//res, _ := bson.MarshalExtJSONIndent(result, false, false, "", "   ")
		res, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(string(res))
	}
}

func Query4(coll *mongo.Collection) {

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"rating", bson.D{{"$gt", 7}}}},
				bson.D{{"rating", bson.D{{"$lte", 10}}}},
			},
		},
	}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []Tea
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(string(res))
	}

	/*
				filter := bson.D{{"type", "Oolong"}, {"rating", 7}}

			    filter := bson.D{
			  {"$and",
			    bson.A{
			      bson.D{{"type", "Oolong"}},
			      bson.D{{"rating", 7}},
			    }},
			}


		bson.A Nedir ve Ne İşe Yarar?
		bson.A, Go'nun MongoDB Driver kütüphanesinde tanımlı bir veri tipidir. Bu, BSON formatında array (dizi) verilerini temsil etmek için kullanılan bir veri yapısıdır. Aslında, []interface{} türüne denk gelir ve BSON belgelerinde dizi olarak kullanılan veri tiplerini Go'da ifade eder.
			ikiside aynı sonucu donduru

	*/
}

func Query5(coll *mongo.Collection) {

	// Aşağıdaki örnek, vendoralanın mevcut olmadığı belgelerle eşleşir:

	filter := bson.D{{"vendor", bson.D{{"$exists", false}}}}

	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		panic(err)
	}
	var results []Tea

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(string(res))
	}
}

func Query6(coll *mongo.Collection) {
	//Aşağıdaki örnek, type"E" harfiyle başlayan belgelerle eşleşir:

	filter := bson.D{{"type", bson.D{{"$regex", "^E"}}}}

	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		panic(err)
	}
	var results []Tea

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(string(res))
	}
}

func Query7(coll *mongo.Collection) {
	filter := bson.D{{"vendor", bson.D{{"$all", bson.A{"A"}}}}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []Tea
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(string(res))
	}
}

func Query8(coll *mongo.Collection) {
	filter := bson.D{{"rating", bson.D{{"$bitsAllSet", 6}}}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []Tea
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(string(res))
	}
}
