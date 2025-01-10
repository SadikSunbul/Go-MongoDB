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

	coll := client.Database("deneme1").Collection("Drink")

	// Query
	Query4(coll)
}

type Drink struct {
	Description string
	Sizes       []float64 `bson:"sizes,truncate"`
	Styles      []string
}

func Query1(coll *mongo.Collection) {
	docs := []interface{}{
		Drink{Description: "Matcha Latte", Sizes: []float64{12, 16, 20}, Styles: []string{"iced", "hot", "extra hot"}},
	}
	result, err := coll.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	fmt.Println("ıds:", result.InsertedIDs)
}

func Query2(coll *mongo.Collection) {
	filter := bson.D{{"sizes", bson.D{{"$lte", 16}}}}
	// Burada "sizes" adlı bir alanın değerlerinin 16'dan küçük veya eşit ($lte) olduğu belgeler seçiliyor.
	// Yani, sadece `sizes` alanında 16 veya daha küçük bir değer içeren belgeleri hedefliyoruz.
	update := bson.D{{"$inc", bson.D{{"sizes.$", -2}}}}
	// "sizes.$" kullanılarak, filtreye uyan listedeki **ilk eşleşen değerden** 2 çıkarılır.
	// $inc operatörü, sayısal bir değeri artırmak veya azaltmak için kullanılır.
	// Burada -2 kullanıldığı için değer 2 azaltılacaktır.
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After)
	// SetReturnDocument(options.After), güncellemeden **sonra** belgenin döndürülmesini sağlar.
	// Yani, güncellenmiş halini alırız.

	var updatedDoc Drink
	err := coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDoc)
	if err != nil {
		panic(err)
	}

	res, _ := bson.MarshalExtJSON(updatedDoc, false, false)
	fmt.Println(string(res))
}

func Query3(coll *mongo.Collection) {
	// Filtre tanımı: "hotOptions" alanında "hot" kelimesini içeren dizi elemanlarını hedefliyoruz.
	identifier := []interface{}{bson.D{{"hotOptions", bson.D{{"$regex", "hot"}}}}}

	// Güncelleme işlemi: "styles" dizisinde, `hotOptions` kriterine uyan elemanları kaldırıyoruz.
	update := bson.D{{"$unset", bson.D{{"styles.$[hotOptions]", ""}}}}

	// 1. ArrayFilters ile hedef dizi elemanlarını belirtiyoruz.
	// 2. ReturnDocument ile güncelleme sonrası belgeyi döndürmesini istiyoruz.
	opts := options.FindOneAndUpdate().
		SetArrayFilters(options.ArrayFilters{Filters: identifier}).
		SetReturnDocument(options.After)

	var updatedDoc Drink
	err := coll.FindOneAndUpdate(context.TODO(), bson.D{}, update, opts).Decode(&updatedDoc)
	if err != nil {
		panic(err)
	}

	res, _ := bson.MarshalExtJSON(updatedDoc, false, false)
	fmt.Println(string(res))
}

func Query4(coll *mongo.Collection) {
	// $mul operatörü, belirtilen alanın değerlerini verilen katsayıyla çarpmak için kullanılır.
	// "sizes.$[]" ifadesi, `sizes` dizisinin **tüm elemanlarını** hedef alır.
	// Bu örnekte, `sizes` dizisinin tüm elemanları 29.57 ile çarpılacak.
	update := bson.D{{"$mul", bson.D{{"sizes.$[]", 29.57}}}}

	// SetReturnDocument(options.After), güncellemeden **sonra** oluşan belgeyi döndürmesini sağlar.
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After)

	var updatedDoc Drink
	err := coll.FindOneAndUpdate(context.TODO(), bson.D{}, update, opts).Decode(&updatedDoc)
	if err != nil {
		panic(err)
	}

	res, _ := bson.MarshalExtJSON(updatedDoc, false, false)
	fmt.Println(string(res))
}
