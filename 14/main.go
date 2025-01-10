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

	coll := client.Database("deneme1").Collection("Book")

	// Query
	Query1(coll)
}

type Book struct {
	Title  string
	Author string
	Length int32
}

/*
Toplu İşlemler (Bulk Operations)
Toplu işlemler, çok sayıda yazma işlemini tek bir veritabanı çağrısıyla gerçekleştirmenizi sağlar.
Her işlem için ayrı ayrı çağrı yapmak yerine, toplu işlemler birden fazla işlemi tek seferde
veritabanına gönderir. Bu, performansı artırır ve işlem maliyetlerini azaltır.
*/

func Query1(coll *mongo.Collection) {
	docs := []interface{}{
		Book{Title: "My Brilliant Friend", Author: "Elena Ferrante", Length: 331},
		Book{Title: "Lucy", Author: "Jamaica Kincaid", Length: 103},
	}
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}

	fmt.Println("ıds:", result.InsertedIDs)
}

func Query2(coll *mongo.Collection) {
	/*
		Örnek: Şu işlemleri sırasız olarak gerçekleştiren bir toplu işlem:

		İki belge ekler.
		"My Brilliant Friend" kitabını yenisiyle değiştirir.
		Uzunluğu 200’den az olan kitapların uzunluğunu 10 artırır.
		Yazarında "Jam" geçen tüm kitapları siler.
	*/
	models := []mongo.WriteModel{
		mongo.NewInsertOneModel().SetDocument(Book{Title: "Middlemarch", Author: "George Eliot", Length: 904}),
		mongo.NewInsertOneModel().SetDocument(Book{Title: "Pale Fire", Author: "Vladimir Nabokov", Length: 246}),
		mongo.NewReplaceOneModel().SetFilter(bson.D{{"title", "My Brilliant Friend"}}).
			SetReplacement(Book{Title: "Atonement", Author: "Ian McEwan", Length: 351}),
		mongo.NewUpdateManyModel().SetFilter(bson.D{{"length", bson.D{{"$lt", 200}}}}).
			SetUpdate(bson.D{{"$inc", bson.D{{"length", 10}}}}),
		mongo.NewDeleteManyModel().SetFilter(bson.D{{"author", bson.D{{"$regex", "Jam"}}}}),
	}
	opts := options.BulkWrite().SetOrdered(false) // İşlemler sırasız olarak gerçekleştirilir.,Hangi işlem önce, hangi işlem sonra yapılır, bu garanti edilmez.,Bir işlem hata verirse, hata o işlemle sınırlı kalır ve diğer işlemler devam eder.
	results, err := coll.BulkWrite(context.TODO(), models, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Eklenen belge sayısı: %d\n", results.InsertedCount)
	fmt.Printf("Güncellenen veya değiştirilen belge sayısı: %d\n", results.ModifiedCount)
	fmt.Printf("Silinen belge sayısı: %d\n", results.DeletedCount)

}
