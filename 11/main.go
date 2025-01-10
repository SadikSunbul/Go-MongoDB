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
	Query7(client)
	Query8(client)
	//	coll := client.Database("deneme1").Collection("insertion")

	// Query

}

func Query3(client *mongo.Client) {
	db := client.Database("testdb")

	// Koleksiyon oluşturma (şema doğrulama ile)
	validationOptions := options.CreateCollection().SetValidator(bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"name", "age"},
			"properties": bson.M{
				"name": bson.M{"bsonType": "string"},
				"age":  bson.M{"bsonType": "int"},
			},
		},
	})

	err := db.CreateCollection(context.TODO(), "users", validationOptions)
	if err != nil {
		fmt.Println("Koleksiyon zaten mevcut olabilir:", err)
		return
	}

	fmt.Println("Koleksiyon başarıyla oluşturuldu.")
}

func Query4(client *mongo.Client) {
	db := client.Database("testdb")
	coll := db.Collection("users")

	// Geçerli bir belge ekleme
	_, err := coll.InsertOne(context.TODO(), bson.M{"name": "John", "age": 30})
	if err != nil {
		fmt.Println("Belge eklenemedi:", err)
	} else {
		fmt.Println("Geçerli belge başarıyla eklendi.")
	}

	// Geçersiz bir belge ekleme
	_, err = coll.InsertOne(context.TODO(), bson.M{"name": "Jane", "age": 30, "test": 30})
	if err != nil {
		fmt.Println("Geçersiz belge eklenemedi:", err)
	} else {
		fmt.Println("Geçersiz belge başarıyla eklendi.")
	}
}

func Query5(client *mongo.Client) {
	db := client.Database("testdb")
	coll := db.Collection("users")

	// Geçerli bir belge ekleme
	_, err := coll.InsertOne(context.TODO(), bson.M{"name": "John", "age": 30})
	if err != nil {
		fmt.Println("Belge eklenemedi:", err)
	} else {
		fmt.Println("Geçerli belge başarıyla eklendi.")
	}

	// Geçersiz bir belge ekleme
	_, err = coll.InsertOne(context.TODO(), bson.M{"name": "Jane", "age": "otuz"})
	if err != nil {
		fmt.Println("Geçersiz belge eklenemedi:", err)
	} else {
		fmt.Println("Geçersiz belge başarıyla eklendi.")
	}
}

func Query7(client *mongo.Client) {
	db := client.Database("testdb1")

	// Koleksiyon oluşturma (şema doğrulama ile)
	validationOptions := options.CreateCollection().SetValidator(bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"name", "age"},
			"properties": bson.M{
				"_id":  bson.M{"bsonType": "objectId"}, // "_id" alanı için izin (bunu eklemeyi unutmayın degılse calısmıyor)
				"name": bson.M{"bsonType": "string"},
				"age":  bson.M{"bsonType": "int"},
			},
			"additionalProperties": false, // Ekstra alanlara izin verilmez
		},
	})

	err := db.CreateCollection(context.TODO(), "users", validationOptions)
	if err != nil {
		fmt.Println("Koleksiyon zaten mevcut olabilir:", err)
		return
	}

	fmt.Println("Koleksiyon başarıyla oluşturuldu.")
}

func Query8(client *mongo.Client) {
	db := client.Database("testdb1")
	coll := db.Collection("users")

	// Geçerli bir belge ekleme
	_, err := coll.InsertOne(context.TODO(), bson.M{"name": "John", "age": 30})
	if err != nil {
		fmt.Println("Belge eklenemedi:", err)
	} else {
		fmt.Println("Geçerli belge başarıyla eklendi.")
	}

	// Geçersiz bir belge ekleme (ekstra alan nedeniyle başarısız olmalı)
	_, err = coll.InsertOne(context.TODO(), bson.M{"name": "Jane", "age": 30, "test": 30})
	if err != nil {
		fmt.Println("Geçersiz belge eklenemedi:", err)
	} else {
		fmt.Println("Geçersiz belge başarıyla eklendi.")
	}
}

func Query9(coll *mongo.Collection) {
	/*
		1. SetHint()
		Bu seçenek, hangi dizinin (index) silme işleminde kullanılacağını belirtir.
		MongoDB, belgeleri daha hızlı bulmak için dizinleri kullanır. Eğer filtreye uygun bir
		dizin varsa bunu belirterek işlemin performansını artırabilirsiniz.

		Neden kullanılır? Özellikle büyük koleksiyonlarda silme işlemini hızlandırmak için kullanılır.
		Eğer MongoDB’nin varsayılan bir dizin seçmesini istemiyorsanız veya özel bir dizin kullanmak
		istiyorsanız bu seçenek yararlıdır.
	*/

	filter := bson.D{{"length", bson.D{{"$gt", 300}}}}
	opts := options.Delete().SetHint(bson.D{{"_id", 1}}) // "_id" dizinini kullanmasını belirtiyoruz.
	result, err := coll.DeleteMany(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)

}

func Query10(coll *mongo.Collection) {
	/*
		 SetCollation()
		Bu seçenek, sorguların dilsel ve bölgesel sıralama kurallarına (collation) göre nasıl e
		şleştirileceğini belirler. Örneğin, büyük/küçük harfe duyarlı mı olacak veya Unicode
		sıralaması mı kullanılacak gibi detayları kontrol eder.

		Neden kullanılır? Dilsel sıralama kurallarıyla uyumlu belgeleri silmek istediğinizde bu
		seçenek önemlidir. Örneğin, İngilizce büyük/küçük harf duyarlılığı olmayan bir sıralama
		kullanabilirsiniz.
	*/
	collation := options.Collation{Locale: "en", Strength: 2} // Büyük/küçük harfe duyarlılığı kapalı
	opts := options.Delete().SetCollation(&collation)
	filter := bson.D{{"title", "lucy"}} // "Lucy" veya "lucy" eşleşebilir.
	result, err := coll.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)

}
