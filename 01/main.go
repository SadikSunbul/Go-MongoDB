package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	coll := client.Database("deneme1").Collection("movies")

	//moviesGet1(coll)
	//moviesInsert(coll)
	//moviesUpdate(coll)
	//moviesUpdateMany(coll)
	//moviesReplace(coll)
	//moviesDelete(coll)
	//moviesDeleteMany(coll)
	//moviesBulkWrite(coll)
	//moviesMonitorDataChages(coll)
	//moviesEstimatedDocumentCount(coll)
	//moviesdistinctTitles(coll)
	moviesRunACommand(coll)

}

type Movie struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"` // omitempty yazarsak ıstege baglı olur boş ise yokmuş gibi gönderılır mesele create için boş gitmeli ki mongodb kendısı ıd atasın , ama get yaparken ıd nın gelmesını ıstıyoruz.
	Title string             `bson:"title"  json:"title"`
	Year  int                `bson:"year"  json:"year"`
	Genre string             `bson:"genre"  json:"genre"`
}

func moviesGet1(coll *mongo.Collection) {

	/*
		TODO : $lte sorguları

		$eq  :Belirtilen değere eşit olan değerleri eşleştirir.
		$gt  :Belirtilen değerden büyük olan değerlerle eşleşir.
		$gte :Belirtilen değerden büyük veya ona eşit olan değerlerle eşleşir.
		$in  :Bir dizide belirtilen değerlerden herhangi biriyle eşleşir.
		$lt  :Matches values that are less than a specified value.
		$lte :Belirtilen değerden küçük veya ona eşit olan değerlerle eşleşir.
		$ne  :Belirtilen değere eşit olmayan tüm değerlerle eşleşir.
		$nin :Bir dizide belirtilen değerlerden hiçbiriyle eşleşmez.

		$$$ Example Usage

		bson.D{{"year",bson.D{"$gte",2000}}} // burada year ı 200 den buyuk veya eşit olanları filtreler

	*/

	filter := bson.D{{"year", bson.D{{"$lte", 1999}}}, {"title", "Back to the Future"}}
	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}

	var movies []Movie
	if err = cursor.All(context.TODO(), &movies); err != nil {
		panic(err)
	}

	for _, movie := range movies {
		jsondata, err := json.MarshalIndent(movie, "", "  ")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("\n%s\n", jsondata)
	}
}

func moviesInsert(coll *mongo.Collection) {
	movies := Movie{Title: "Test", Year: 2222, Genre: "Test"}

	result, err := coll.InsertOne(context.TODO(), movies)
	if err != nil {
		fmt.Errorf("Erorrs:", err.Error())
		return
	}

	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}

func moviesUpdate(coll *mongo.Collection) {

	id, _ := primitive.ObjectIDFromHex("675d3a322979f406206c8341")

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"title", "DENEME"}, {"year", 9999}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Errorf("Error:", err.Error())
		return
	}

	fmt.Printf("Matched %d documents and updated %d documents.\n", result.MatchedCount, result.ModifiedCount)

}

func moviesUpdateMany(coll *mongo.Collection) {
	filter := bson.D{{"year", bson.D{{"$gte", 2000}}}} // TODO : $gte
	update := bson.D{{"$set", bson.D{{"year", 2024}}}} //year alanını güncelle

	result, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		fmt.Errorf("Error:", err.Error())
		return
	}
	fmt.Printf("Matched %d documents and updated %d documents.\n", result.MatchedCount, result.ModifiedCount)
}

func moviesReplace(coll *mongo.Collection) {

	id, _ := primitive.ObjectIDFromHex("675d3a322979f406206c8341")
	filter := bson.D{{"_id", id}}

	movies := Movie{Title: "Alo"}

	result, err := coll.ReplaceOne(context.TODO(), filter, movies)

	if err != nil {
		fmt.Errorf("Erorr:", err.Error())
		return
	}

	fmt.Printf("Matched %d documents and replaced %d documents.\n", result.MatchedCount, result.ModifiedCount)
}

func moviesDelete(coll *mongo.Collection) {
	id, _ := primitive.ObjectIDFromHex("675d3a322979f406206c8341")
	filter := bson.D{{"_id", id}}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Errorf("Erorr:", err.Error())
		return
	}
	fmt.Printf("Deleted %d documents in the trainers collection\n", result.DeletedCount)
}

func moviesDeleteMany(coll *mongo.Collection) {

	filter := bson.D{{"year", 2024}}

	result, err := coll.DeleteMany(context.TODO(), filter)

	if err != nil {
		fmt.Errorf("Erorr:", err.Error())
		return
	}

	fmt.Printf("Deleted %d documents in the trainers collection\n", result.DeletedCount)
}

func moviesBulkWrite(coll *mongo.Collection) { // Toplu işlem yapmayı ve performansı artırır

	models := []mongo.WriteModel{
		mongo.NewInsertOneModel().SetDocument(Movie{Title: "Sadık", Year: 2024, Genre: "Dram"}),
		mongo.NewUpdateManyModel().SetFilter(bson.D{{"year", 1999}}).SetUpdate(bson.D{{"$set", bson.D{{"year", 2025}}}}),
	}

	opts := options.BulkWrite().SetOrdered(true) // sıralı yap işlemleri

	results, err := coll.BulkWrite(context.TODO(), models, opts)

	if err != nil {
		fmt.Errorf("Erorr:", err.Error())
		return
	}

	fmt.Printf("Inserted %d documents\n", results.InsertedCount)
	fmt.Printf("Updated %d documents\n", results.ModifiedCount)
	fmt.Printf("Deleted %d documents\n", results.DeletedCount)
}

func moviesMonitorDataChages(coll *mongo.Collection) { // burası ıcın databasenın en bastan yapılandırılmaası lazımdır

	// insert işlemi izlenicek dedik
	pipeline := mongo.Pipeline{bson.D{{"$match", bson.D{{"operationType", "insert"}}}}}

	cs, err := coll.Watch(context.TODO(), pipeline)
	if err != nil {
		panic(err)
	}
	defer cs.Close(context.TODO())

	fmt.Println("Waiting For Change Events. Insert something in MongoDB!")

	// Prints a message each time the change stream receives an event
	for cs.Next(context.TODO()) {
		var event bson.M
		if err := cs.Decode(&event); err != nil {
			panic(err)
		}
		output, err := json.MarshalIndent(event["fullDocument"], "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
	if err := cs.Err(); err != nil {
		panic(err)
	}
}

func moviesEstimatedDocumentCount(coll *mongo.Collection) {
	filter := bson.D{{"title", "Test"}}

	/*
		Bir koleksiyondaki belge sayısı hakkında yaklaşık bir bilgi edinmek için
		EstimatedDocumentCount() metodunu kullanabilirsiniz.

		Tam belge sayısı için ise CountDocuments() metodunu kullanabilirsiniz.
	*/

	// Koleksiyondaki tahmini belge sayısını alır ve yazdırır
	estCount, estCountErr := coll.EstimatedDocumentCount(context.TODO())
	if estCountErr != nil {
		panic(estCountErr)
	}

	// Koleksiyondaki belgelerin sayısını alır ve yazdırır
	// filtreye uyan
	count, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Estimated document count: %d\n", estCount)
	fmt.Printf("Actual document count: %d\n", count)
}

func moviesdistinctTitles(coll *mongo.Collection) {

	// Filtre: year alanında 2025 geçen belgeleri eşleştir
	filter := bson.D{{"title", "Back to the Future"}}
	// Distinct() ile "title" alanındaki benzersiz değerleri al
	results, err := coll.Distinct(context.TODO(), "year", filter)
	// Bu kod, MongoDB koleksiyonunda "title" değeri "Back to the Future" olan tüm belgeleri bulur
	//Sonra bu belgelerin "year" alanlarındaki tekrar etmeyen farklı değerlerin listesini alır.
	if err != nil {
		fmt.Printf("Error while performing distinct operation: %v\n", err)
		return
	}

	// Sonuçları yazdır
	fmt.Println("Distinct year:")
	for _, result := range results {
		fmt.Println(result)
	}
}

func moviesRunACommand(coll *mongo.Collection) {

	command := bson.D{{"dbStats", 1}}

	// RunCommand() metodu, belirtilen komutu veritabanına göndererek sonucu alır
	var result bson.M // bson.M: MongoDB'den dönen verileri anahtar-değer çiftleri (map[string]interface{}) olarak saklar.

	err := coll.Database().RunCommand(context.TODO(), command).Decode(&result)

	// Prints a message if any errors occur during the command execution
	if err != nil {
		panic(err)
	}
	// 2. parametre başata ne olucak 3. parametre ise elemanların basında kac boluk olucak
	output, err := json.MarshalIndent(result, "", "    ") // result değişkenindeki veriyi JSON formatına dönüştürür.
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
}
