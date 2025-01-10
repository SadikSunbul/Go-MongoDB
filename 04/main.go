package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	coll := client.Database("deneme1").Collection("tea2")

	// Query

	Query2(coll)

}

type Tea struct {
	Type   string
	Rating int32
}

/*
Bir mahalledeki tüm mektupları dağıtan bir postacıyı düşünün:

Postacı mektupları taşırken elindeki çantada tüm mektupları değil, yalnızca bir kısmını (bir toplu) taşır.
Her eve geldiğinde çantasından sırayla bir mektup çıkarır ve teslim eder.
Teslimat tamamlanınca postacı çantasını yeniden doldurur ve işlemi tekrarlar.
Bu örnekte:

İmleç (Cursor): Postacının çantasıdır. Belgelerin bir alt kümesini içerir.
Belgeler: Mektuplardır. İşlenecek veriyi temsil eder.
Next() veya TryNext(): Postacının çantasından sırayla mektup alıp teslim etmesi.
All(): Tüm mektupları tek seferde almak ve işlem yapmak.
Close(): Çantayı kapatmak (kaynakları serbest bırakmak).
*/

func Query1(coll *mongo.Collection) {
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) { // Postacı, çantasından mektupları sırayla çıkarır ve her mektubu bir eve bırakır.
		// Amaç: Belgeleri sırayla işlemek. Büyük veri kümeleriyle çalışırken belleği etkin kullanır.
		var result Tea
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Document: %+v\n", result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
}

func Query2(coll *mongo.Collection) {
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	for {
		if cursor.TryNext(context.TODO()) { // Postacı çantasındaki bir sonraki mektubu teslim etmeye çalışır, ancak çantası boşsa yeni bir mektup almak için bekler.
			/*
						Metafor: Postacı mektup bırakmaya çalışır, ama çantası boşsa yeni bir mektup gelene kadar bekler.
				        Amaç: Belgelerin hemen mevcut olmadığı durumlarda verimli bir şekilde çalışmak.
			*/
			var result Tea
			if err := cursor.Decode(&result); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Document: %+v\n", result)
			continue
		}
		if err := cursor.Err(); err != nil {
			log.Fatal(err)
		}
		if cursor.ID() == 0 {
			break
		}
	}

}

func Query3(coll *mongo.Collection) {
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO()) // Postacı çantasını kullanmayı bitirdiğinde, çantayı kapatır ve dağıtım işlemini sonlandırır.

	var results []Tea
	/*
		Metafor: Postacı, bir çanta yerine büyük bir koli alır ve içindeki tüm mektupları bir seferde mahalleye bırakır.
		Amaç: Küçük bir veri kümesini hızla işlemek. Ancak, büyük veri kümelerinde bellek sorunlarına yol açabilir.
	*/
	if err = cursor.All(context.TODO(), &results); err != nil { // Postacı, tüm mektupları bir arada alır ve hepsini aynı anda mahalleye dağıtır.
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("Document: %+v\n", result)
	}

}

/*
Next() ve TryNext() büyük veri kümeleriyle çalışırken verimli bir şekilde işlenir.
All(), küçük veri kümelerinde hızlı ve basit bir yöntemdir.
Close(), kaynak yönetimi için kritik bir adımdır.
*/

/* ___OZET___
Next()  	-> Teker teker belge işlemek gerektiğinde.
TryNext()   -> Kuyruklanabilir sonuçları işlemek gerektiğinde.
All()       ->Tüm sonuçları bir diziye topluca almak gerektiğinde.
Close()     -> İmleç kullanımını bitirip kaynakları serbest bırakmak gerektiğinde.
*/
