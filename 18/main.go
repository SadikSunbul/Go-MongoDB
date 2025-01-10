package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
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

	coll := client.Database("deneme1").Collection("movies")

	// Query
	Query1(coll, client)
}

////// Transactions  //////

// Bu kılavuzda, MongoDB Go Sürücüsünü kullanarak işlemleri nasıl gerçekleştireceğinizi öğrenebilirsiniz . İşlemler, işlem tamamlanana kadar hiçbir veriyi değiştirmeyen bir dizi işlem çalıştırmanıza olanak tanır. İşlemdeki herhangi bir işlem bir hata döndürürse, sürücü işlemi iptal eder ve görünür hale gelmeden önce tüm veri değişikliklerini atar.

func Query1(coll *mongo.Collection, client *mongo.Client) {
	session, err := client.StartSession()
	if err != nil {
		log.Fatalf("Oturum başlatılamadı: %v", err)
	}
	defer session.EndSession(context.TODO())
	/*
	   . Oturum Başlatma
	   MongoDB'de işlemler oturumlar içinde çalışır. Bir oturum başlatmak için istemciden StartSession() yöntemi çağrılır:
	*/
	txnOptions := options.Transaction().SetWriteConcern(writeconcern.Majority())
	result, err := session.WithTransaction(context.TODO(), func(ctx mongo.SessionContext) (interface{}, error) {
		// İşlem kapsamında birden fazla işlem gerçekleştirin
		result, err := coll.InsertMany(ctx, []interface{}{
			bson.D{{"title", "The Bluest Eye"}, {"author", "Toni Morrison"}},
			bson.D{{"title", "Sula"}, {"author", "Toni Morrison"}},
		})
		if err != nil {
			return nil, err // Hata oluşursa işlem iptal edilir
		}
		return result, nil
	}, txnOptions)

	if err != nil {
		log.Fatalf("İşlem başarısız oldu: %v", err)
	}

	fmt.Printf("Sonuc: %v\n", result)
}
