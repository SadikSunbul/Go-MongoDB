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

	Query5(coll)

}

type Tea struct {
	Type   string
	Rating int32
}

func Query1(coll *mongo.Collection) {
	docs := []interface{}{
		Tea{Type: "Masala", Rating: 10},
		Tea{Type: "Matcha", Rating: 7},
		Tea{Type: "Assam", Rating: 4},
		Tea{Type: "Oolong", Rating: 9},
		Tea{Type: "Chrysanthemum", Rating: 5},
		Tea{Type: "Earl Grey", Rating: 8},
		Tea{Type: "Jasmine", Rating: 3},
		Tea{Type: "English Breakfast", Rating: 6},
		Tea{Type: "White Peony", Rating: 4},
	}
	result, err := coll.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	fmt.Println("ıds:", result.InsertedIDs)
}

func Query2(coll *mongo.Collection) {
	/*
		Sorgu filtrenizle eşleşen belge sayısını saymak için CountDocuments()yöntemini kullanın. Boş bir sorgu filtresi geçirirseniz, bu yöntem koleksiyondaki toplam belge sayısını döndürür.
	*/
	opts := options.Count().SetHint("_id_")
	//options.Count(): MongoDB'nin CountDocuments fonksiyonunda kullanılabilecek bir opsiyonlar (seçenekler) nesnesi oluşturur.
	// SetHint("_id_"): Veritabanına, belge sayımı işlemini hızlandırmak için _id indeksini kullanmasını belirtir.

	count, err := coll.CountDocuments(context.TODO(), bson.D{}, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Count: %d\n", count)

	/*
			Bu CountOptionstür, seçenekleri aşağıdaki yöntemlerle yapılandırmanıza olanak tanır:

			SetCollation() -> Sonuçları sıralarken kullanılacak dil sıralama türü. Varsayılan:nil
			SetHint()      ->Sayım yapılacak belgeleri taramak için kullanılacak dizin.Varsayılan:nil
			SetLimit()     ->Sayılacak maksimum belge sayısı. Varsayılan:0
			SetMaxTime()   ->Sorgunun sunucuda çalışabileceği maksimum süre. Varsayılan:nil
		    SetSkip()      ->Sayılmadan önce atlanacak belge sayısı. Varsayılan:0
	*/
}

func Query3(coll *mongo.Collection) {
	filter := bson.D{{"rating", bson.D{{"$lt", 6}}}}

	count, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of documents with a rating less than six: %d\n", count)
}

func Query4(coll *mongo.Collection) {
	matchStage := bson.D{{"$match", bson.D{{"rating", bson.D{{"$gt", 5}}}}}}
	countStage := bson.D{{"$count", "counted_documents"}}

	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{matchStage, countStage})
	if err != nil {
		panic(err)
	}

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

func Query5(coll *mongo.Collection) {
	// Koleksiyonunuzdaki belge sayısını tahmin etmek için EstimatedDocumentCount()yöntemini kullanın.
	// EstimatedDocumentCount() yöntemi, tüm koleksiyonu taramak yerine koleksiyonun meta verilerini kullandığı için CountDocuments() yönteminden daha hızlıdır.
	count, err := coll.EstimatedDocumentCount(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Estimated number of documents in the tea collection: %d\n", count)
	/*
			Davranışı Değiştir
			EstimatedDocumentCount()Bir tür geçirerek davranışını değiştirebilirsiniz EstimatedDocumentCountOptions. Herhangi bir seçenek belirtmezseniz, sürücü varsayılan değerlerini kullanır.

			Bu EstimatedDocumentCountOptionstür, seçenekleri aşağıdaki yöntemlerle yapılandırmanıza olanak tanır:

		    SetMaxTime()  -->Sorgunun sunucuda çalışabileceği maksimum süre. Varsayılan:nil
	*/
}
