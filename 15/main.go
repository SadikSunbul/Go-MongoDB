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

	coll := client.Database("deneme1").Collection("Course")

	// Query
	Query3(coll)
}

// Kurs bilgilerini tutmak için Course yapısı
type Course struct {
	Title      string // Kurs başlığı
	Enrollment int32  // Kaydolma sayısı
}

func Query1(coll *mongo.Collection) {

	// Örnek veri eklemek için koleksiyon ve belgeler oluşturuluyor
	docs := []interface{}{
		Course{Title: "Representation Theory", Enrollment: 40},
		Course{Title: "Early Modern Philosophy", Enrollment: 25},
		Course{Title: "Animal Communication", Enrollment: 18},
	}
	result, err := coll.InsertMany(context.TODO(), docs) // Belgeler ekleniyor
	if err != nil {
		panic(err) // Eğer bir hata oluşursa program durur
	}

	fmt.Println("ıds:", result.InsertedIDs)
}

func Query2(coll *mongo.Collection) {
	// "Enrollment" değeri 20'den küçük olan ilk belgeyi bul ve sil
	filter := bson.D{{"enrollment", bson.D{{"$lt", 20}}}} // Filtre koşulu
	var deletedDoc Course                                 // Silinen belgeyi tutmak için değişken

	err := coll.FindOneAndDelete(context.TODO(), filter).Decode(&deletedDoc)
	if err != nil {
		panic(err) // Hata durumunda program durur
	}

	// Silinen belgeyi JSON formatında yazdır
	res, _ := bson.MarshalExtJSON(deletedDoc, false, false)
	fmt.Println("Silinen Belge:", string(res))

}

func Query3(coll *mongo.Collection) {
	// "Title" değeri "Modern" kelimesini içeren ilk belgeyi bul ve güncelle
	filter := bson.D{{"title", bson.D{{"$regex", "Modern"}}}} // Filtre koşulu
	update := bson.D{{"$set", bson.D{{"enrollment", 32}}}}    // Güncelleme işlemi

	// Güncellenmiş belgeyi döndürmek için seçenekler
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedDoc Course // Güncellenen belgeyi tutmak için değişken
	err := coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDoc)
	if err != nil {
		panic(err) // Hata durumunda program durur
	}

	// Güncellenmiş belgeyi JSON formatında yazdır
	res, _ := bson.MarshalExtJSON(updatedDoc, false, false)
	fmt.Println("Güncellenen Belge:", string(res))
}

func Query4(coll *mongo.Collection) {
	// "Title" değeri "Representation Theory" olan belgeyi bul ve değiştir
	filter := bson.D{{"title", "Representation Theory"}} // Filtre koşulu

	// Yeni belge (eski belge bununla değiştirilecek)
	replacement := Course{Title: "Combinatorial Theory", Enrollment: 35}

	var outdatedDoc Course // Değiştirilen eski belgeyi tutmak için değişken
	err := coll.FindOneAndReplace(context.TODO(), filter, replacement).Decode(&outdatedDoc)
	if err != nil {
		panic(err) // Hata durumunda program durur
	}

	// Eski belgeyi JSON formatında yazdır
	res, _ := bson.MarshalExtJSON(outdatedDoc, false, false)
	fmt.Println("Değiştirilen Belge:", string(res))
}
