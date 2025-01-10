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
	coll := client.Database("deneme1").Collection("dish")

	// Query

	Query8(coll)

}

type Dish struct {
	Name        string
	Description string
}

func Query1(coll *mongo.Collection) {
	docs := []interface{}{
		Dish{Name: "Shepherd’s Pie", Description: "A vegetarian take on the classic dish that uses lentils as a base. Serves 2."},
		Dish{Name: "Green Curry", Description: "A flavorful Thai curry, made vegetarian with fried tofu. Vegetarian and vegan friendly."},
		Dish{Name: "Herbed Whole Branzino", Description: "Grilled whole fish stuffed with herbs and pomegranate seeds. Serves 3-4."},
		Dish{Name: "Kale Tabbouleh", Description: "A bright, herb-based salad. A perfect starter for vegetarians and vegans."},
		Dish{Name: "Garlic Butter Trout", Description: "Baked trout seasoned with garlic, lemon, dill, and, of course, butter. Serves 2."},
	}

	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}

	fmt.Println("ıds:", result.InsertedIDs)
}

func Query2(coll *mongo.Collection) {
	model := mongo.IndexModel{Keys: bson.D{{"description", "text"}}}
	/*
		mongo.IndexModel: MongoDB'de bir index tanımlamak için kullanılan yapı.
		Keys: Hangi alanlarda ve nasıl bir index oluşturulacağını belirler.
		bson.D{{"description", "text"}}:
		"description" alanında bir text index oluşturulmasını sağlar.
		"text": Bu, MongoDB'ye bu alanın metin aramaları için optimize edilmiş bir text index olduğunu belirtir.
	*/
	name, err := coll.Indexes().CreateOne(context.TODO(), model)
	/*
		coll.Indexes(): Koleksiyon üzerindeki index'lerle ilgili işlemler yapmamızı sağlayan bir API.
		CreateOne:
		Bu, belirtilen model (burada model) ile bir index oluşturur.
		context.TODO(): İşlemin hangi bağlamda yapılacağını belirtir. Basit işlemler için TODO kullanılır.
		Geri dönen değerler:
		name: Oluşturulan index'in adı. MongoDB, otomatik olarak bir ad oluşturur, ancak özel bir ad belirlemek isterseniz bunu Options ile yapabilirsiniz.
		err: İşlem sırasında bir hata meydana gelirse bu değişkende tutulur.
	*/
	if err != nil {
		panic(err)
	}
	/*
		Bu Kod Ne İşe Yarar?
		Text Index Oluşturma: Kod, description alanı için bir text index oluşturur. Bu index sayesinde:
		$text filtresi kullanılarak, description alanında kelime bazlı arama yapılabilir.
		Arama işlemleri, index sayesinde çok daha hızlı çalışır.
	*/

	fmt.Println("Name of index created: " + name)
}

func Query3(coll *mongo.Collection) { // Query 2 yi çalıştırmadan burası calısmıyacaktır
	filter := bson.D{{"$text", bson.D{{"$search", "SERVES fish"}}}} // bosluk veya gibi calısır buyuk kucuk harf duyarlıgı yoktur
	/*
		$text, MongoDB'nin metin arama sorgularında kullanılan bir operatördür.
		Bu operatör yalnızca bir veya daha fazla alan için text index tanımlandığında çalışır.
		Text index, yalnızca belirttiğiniz alanlarda kelime bazlı arama yapar. Bu nedenle, description gibi belirli bir alanı text index olarak belirlemeden, $text ile arama yapamazsınız.
	*/

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []Dish
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}

/*
MongoDB'nin text index ve $text operatörünü kullanırken varsayılan davranışıdır. $text sorgularında belirli kurallar vardır ve bunları doğru şekilde anlamak önemlidir:

1. "and" gibi kelimeler neden sonuç dönmüyor?
MongoDB'nin text search özelliği, durdurma kelimeleri (stop words) olarak bilinen bazı kelimeleri otomatik olarak filtreler. Bu kelimeler, dil bağımsız olarak genellikle sıkça kullanılan, ancak arama için pek anlam ifade etmeyen kelimelerdir. Örneğin:

İngilizce için: and, the, is, a, of, in, vb.
"and" kelimesi bu durdurma kelimelerinden biridir, bu yüzden text search sorgusunda dikkate alınmaz.

2. Bu davranışı nasıl değiştirebilirim?
Eğer "and" gibi durdurma kelimelerini de arama sonuçlarında kullanmak istiyorsanız, şunları yapabilirsiniz:

a. Text Index için Özel Ayarlar Kullanmak
default_language özelliğini ayarlayarak durdurma kelimeleri devre dışı bırakabilirsiniz. Örneğin, index oluştururken şu şekilde bir ayar yapabilirsiniz:

model := mongo.IndexModel{
    Keys:    bson.D{{"description", "text"}},
    Options: options.Index().SetDefaultLanguage("none"),
}
name, err := coll.Indexes().CreateOne(context.TODO(), model)
if err != nil {
    panic(err)
}
Bu ayar, text index'in dil ayarını kaldırır ve tüm kelimeleri olduğu gibi aramanıza izin verir.


*/

func Query4(coll *mongo.Collection) {
	filter := bson.D{{"$text", bson.D{{"$search", "\"serves 2\""}}}} // bu sadece serves 2 yı arar normalde bosluk or yerıne gecıyordu oradan kacmıs olduk

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []Dish
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}

func Query5(coll *mongo.Collection) {
	//	Aşağıdaki örnek, "vegan" terimini içeren, ancak "tofu" terimini içermeyen açıklamalar için bir metin araması çalıştırır:
	filter := bson.D{{"$text", bson.D{{"$search", "vegan -tofu"}}}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []Dish
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}

func Query6(coll *mongo.Collection) {
	filter := bson.D{{"$text", bson.D{{"$search", "vegetarian"}}}}
	sort := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
	projection := bson.D{{"name", 1}, {"description", 1}, {"score", bson.D{{"$meta", "textScore"}}}, {"_id", 0}}
	opts := options.Find().SetSort(sort).SetProjection(projection)

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}

func Query7(coll *mongo.Collection) {
	matchStage := bson.D{{"$match", bson.D{{"$text", bson.D{{"$search", "herb"}}}}}}
	// Ayrıca , bir toplama boru hattında metin araması yapmak için $match$text aşamasına değerlendirme sorgu operatörünü de ekleyebilirsiniz .
	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{matchStage})
	if err != nil {
		panic(err)
	}

	var results []Dish
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}

func Query8(coll *mongo.Collection) {
	matchStage := bson.D{{"$match", bson.D{{"$text", bson.D{{"$search", "vegetarian"}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"score", bson.D{{"$meta", "textScore"}}}}}}
	projectStage := bson.D{{"$project", bson.D{{"name", 1}, {"score", bson.D{{"$meta", "textScore"}}}, {"_id", 0}}}}

	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{matchStage, sortStage, projectStage})
	if err != nil {
		panic(err)
	}

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}
