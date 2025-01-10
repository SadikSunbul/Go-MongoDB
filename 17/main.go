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

	coll := client.Database("deneme1").Collection("movies")

	// Query
	Query8(coll)
}

// ///// Indexes ////////
// Dizin _id_, tek alan dizininin bir örneğidir. Bu dizin, _idyeni bir koleksiyon oluşturduğunuzda alanda otomatik olarak oluşturulur.

func Query1(coll *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"title", 1}},
	}
	/*
		Aşağıdaki örnek, titlekoleksiyondaki alan üzerinde artan sırada bir dizin oluşturur sample_mflix.movies:
	*/
	name, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created: " + name)
}

// Bileşik index
func Query2(coll *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{"fullplot", -1},
			{"title", 1},
		},
	}
	name, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created: " + name)
}

// Çok Anahtarlı İndeksler (Dizi Alanlarındaki İndeksler)
/*
Multikey Index Nedir?
MongoDB'de Multikey Index, bir dokümanda array (dizi) tipi bir alanda oluşturulan bir tür indekstir.
Bu tür indeksler, array'in her elemanına ayrı ayrı indeks oluşturur. Bu sayede, sorgularınızda array alanlarını hedeflediğinizde sorgu performansı artar.
*/
func Query3(coll *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"cast", -1}},
	}
	name, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name of Index Created: " + name)
}

// Kümelenmiş Endeksler
// Kümelenmiş dizinler, kümelenmiş koleksiyonlardaki ekleme, güncelleme ve silme işlemlerinin performansını artırır . Kümelenmiş koleksiyonlar, belgeleri kümelenmiş dizin anahtar değerine göre sıralanmış şekilde depolar.
func Query4(client *mongo.Client) {
	db := client.Database("deneme1")
	cio := bson.D{{"key", bson.D{{"_id", 1}}}, {"unique", true}}
	opts := options.CreateCollection().SetClusteredIndex(cio)
	db.CreateCollection(context.TODO(), "tea4", opts)
	/*
		MongoDB, _id alanı için varsayılan bir endeks sağlar, ancak bu kümelenmiş bir endeks değildir.
		Eğer verilerinizin fiziksel olarak belirli bir düzene göre sıralanmasını istiyorsanız
		(örneğin, aralıklı sorgular için), kümelenmiş endeks oluşturmayı düşünmelisiniz.
		Ancak varsayılan _id endeksi çoğu senaryoda yeterlidir ve kümelenmiş endeks genellikle
		büyük veri kümeleri veya performansın çok kritik olduğu durumlarda gereklidir.
	*/
}

// Metin İndeksleri
// Metin indeksleri, MongoDB'de metin tabanlı veriler üzerinde hızlı ve güçlü bir arama yapmanıza olanak tanır. Dize (string) türündeki alanlar üzerinde çalışır ve belirli kelimeleri, ifadeleri ya da anahtar kelimeleri aramak için kullanılır.
func Query5(client *mongo.Client) {
	coll := client.Database("sample_mflix").Collection("movies6")
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{"plot", "text"},
		},
	}

	name, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name of Index Created: " + name)
	/*
		{ "title": "Movie A", "plot": "A thrilling adventure in Italy." }
		{ "title": "Movie B", "plot": "An exciting journey through the mountains." }
		{ "title": "Movie C", "plot": "A romantic story in Rome." }

		filter := bson.M{"$text": bson.M{"$search": "adventure"}}
		cursor, err := coll.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}

		$text:
		Metin indekslerinde kullanılan özel bir operatördür.
		"$search": "adventure": adventure kelimesini plot alanında arar.
	*/
}

// Coğrafi Uzamsal Endeksler
// Coğrafi uzamsal indeksler, coğrafi veriler üzerinde sorgular yapmanıza ve işlemler gerçekleştirmenize olanak tanır. MongoDB, coğrafi koordinatlarla çalışmak için 2dsphere ve 2d olmak üzere iki ana indeks türü sunar.
func Query6(coll *mongo.Collection) {

	indexModel := mongo.IndexModel{
		Keys: bson.D{{"location.geo", "2dsphere"}},
	}
	name, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name of Index Created: " + name)
	/*
		filter := bson.M{
		    "location.geo": bson.M{
		        "$near": bson.M{
		            "$geometry": bson.M{
		                "type":        "Point",
		                "coordinates": []float64{-118.3, 33.9},
		            },
		            "$maxDistance": 5000, // 5 kilometre
		        },
		    },
		}
		cursor, err := coll.Find(context.TODO(), filter)
		if err != nil {
		    panic(err)
		}

	*/

}

/// Benzersiz Endeksler

func Query7(coll *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"theaterId", -1}},
		Options: options.Index().SetUnique(true), // Benzersiz bir dizin oluşturmak için, çoğaltılmasını engellemek istediğiniz alanı veya alan birleşimini belirtin ve uniqueseçeneği olarak ayarlayın true.
	}
	name, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created: " + name)
}

// Bir Dizin'i Kaldır

func Query8(coll *mongo.Collection) {
	res, err := coll.Indexes().DropOne(context.TODO(), "title_1")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
