package main

func main() {

}

/*
	TODO : Comparison Query Operators (Karşılaştırma Sorgu Operatörleri)✅
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

/*
	TODO : Convert ObjectId✅
		objectId, _ := primitiv	e.ObjectIDFromHex("675d3a322979f406206c8341") // strıngi objectId ye ceviri
		stringID:=id.Hex()  // objectId yı stringe ceviri
*/

/*
	TODO : Logical Query Operators (Mantıksal Sorgu Operatörleri)
		$and  :Sorgu ifadelerini mantıksal VE ile birleştirerek her iki ifadenin koşullarına uyan tüm belgeleri döndürür. [bson.D{{"$and", bson.A{bson.D{{"roles", bson.D{{"$eq", "user"}}}},bson.D{{"isActive", bson.D{{"$eq", true}}}},}}}]
		$not  :Sorgu yordamının etkisini tersine çevirir ve sorgu yordamıyla eşleşmeyen belgeleri döndürür.
		$nor  :Sorgu ifadelerini mantıksal bir NOR ile birleştirmek, her iki ifadeyle de eşleşmeyen tüm belgeleri döndürür.
		$or   :Sorgu ifadelerini mantıksal VEYA ile birleştirmek, her iki ifadenin koşullarına uyan tüm belgeleri döndürür.
*/

/*
	TODO : Element Query Operators (Eleman Sorgu Operatörleri)
		$exists  :Belirtilen alana sahip belgelerle eşleşir.  [bson.D{{"vendor", bson.D{{"$exists", false}}}}]
		$type    :Bir alan belirtilen türdeyse belgeleri seçer.

*/

/*
	TODO : Evaluation Query Operators (Değerlendirme Sorgu Operatörleri)
		$expr 		:Sorgu dili içerisinde toplama ifadelerinin kullanılmasına izin verir.
		$jsonSchema :Belgeleri verilen JSON Şemasına göre doğrulayın.
		$mod 		:Bir alanın değeri üzerinde bir modül işlemi gerçekleştirir ve belirtilen sonuca sahip belgeleri seçer.
		$regex 		:Değerleri belirtilen düzenli ifadeyle eşleşen belgeleri seçer. [bson.D{{"type", bson.D{{"$regex", "^E"}}}}]
		$where 		:JavaScript ifadesini karşılayan belgelerle eşleşir.
*/

/*
	TODO : Array Query Operators (Dizi Sorgu Operatörleri) ✅
		✅$all 		 :Sorguda belirtilen tüm öğeleri içeren dizilerle eşleşir. [bson.D{{"roles", bson.M{"$all": bson.A{"user"}}}}] ornke [a,b,c,d] ise all a ise bunu getırı veya all a,c ise de getiri bunu
		✅$elemMatch   :Dizi alanındaki öğe belirtilen tüm $elemMatch koşullarıyla eşleşiyorsa belgeleri seçer. [bson.D{{"scores", bson.M{"$elemMatch": bson.M{"subject": "math","score": bson.M{"$gt": 90}}}}] bırden fazla sorguyu bırlestırı
		✅$size		 :Dizi alanı belirtilen bir boyutta ise belgeleri seçer. [bson.D{{"roles", bson.M{"$size": 1}}}]
*/

/*
	TODO : Projection Operators
		$   		:Sorgu koşuluyla eşleşen dizideki ilk öğeyi yansıtır.
		$elemMatch  :Belirtilen $elemMatch koşuluyla eşleşen bir dizideki ilk öğeyi yansıtır.
		$meta		:Belge başına kullanılabilir meta verileri yansıtır.
		$slice		:Bir diziden yansıtılan eleman sayısını sınırlar. Atlama ve sınırlama dilimlerini destekler.
*/

/*
	TODO : Miscellaneous Query Operators (Çeşitli Sorgu Operatörleri)
		$rand	 :0 ile 1 arasında rastgele bir float değeri üretir.
		$natural :sort() veya hint() yöntemleri aracılığıyla sağlanabilen ve ileri veya geri koleksiyon taramasını zorlamak için kullanılabilen özel bir ipucu.
*/

/*
	TODO : Field Update Operators (Alan Güncelleme Operatörleri)
		$currentDate    :Bir alanın değerini geçerli tarihe, Tarih veya Zaman Damgası olarak ayarlar.
		$inc		    :Alanın değerini belirtilen miktarda artırır.
		$min		    :Yalnızca belirtilen değer mevcut alan değerinden küçükse alanı günceller.
		$max 		    :Yalnızca belirtilen değer mevcut alan değerinden büyükse alanı günceller.
		$mul		    :Alanın değerini belirtilen miktarla çarpar.
		$rename			:Bir alanı yeniden adlandırır.
		✅$set			:Bir belgedeki bir alanın değerini ayarlar.  [bson.D{{"$set", bson.D{{"year", 2024}}}} year alanını 2024 ile güncelle]
		$setOnInsert	:Bir güncellemenin bir belgenin eklenmesiyle sonuçlanması durumunda bir alanın değerini ayarlar. Mevcut belgeleri değiştiren güncelleme işlemleri üzerinde hiçbir etkisi yoktur.
		$unset			:Belirtilen alanı belgeden kaldırır.
*/

/*
	TODO : Array Update Operators (Dizi Güncelleme Operatörleri)
		$					:Sorgu koşuluyla eşleşen ilk öğeyi güncellemek için yer tutucu görevi görür.
		$[]					:Sorgu koşuluyla eşleşen belgeler için dizideki tüm öğeleri güncellemek üzere bir yer tutucu görevi görür.
		$[<identifier>]		:Sorgu koşuluyla eşleşen belgeler için arrayFilters koşuluyla eşleşen tüm öğeleri güncellemek için bir yer tutucu görevi görür.
		$addToSet			:Diziye yalnızca kümede halihazırda mevcut olmayan öğeleri ekler.
		$pop				:Bir dizinin ilk veya son öğesini kaldırır.
		$pull				:Belirtilen sorguyla eşleşen tüm dizi öğelerini kaldırır.
		$push				:Bir diziye bir öğe ekler.
		$pullAll			:Bir diziden eşleşen tüm değerleri kaldırır.
		$each				:Dizi güncellemeleri için birden fazla öğeyi eklemek üzere $push ve $addToSet operatörlerini değiştirir.
		$position			:$push operatörünü, dizideki öğelerin ekleneceği konumu belirtecek şekilde değiştirir.
		$slice				:Güncellenen dizilerin boyutunu sınırlamak için $push operatörünü değiştirir.
		$sort				:Bir dizide saklanan belgeleri yeniden sıralayacak şekilde $push operatörünü değiştirir.
*/

/*
	TODO : Methotlar
		ReplaceOne metodu, filter ile eşleşen bir belgeyi bulur ve onu movies adlı yeni belge ile tamamen değiştirir (yerine koyar).

*/

/*
	TODO : moviesBulkWrite	  (Toplu işlem yapmayı ve performansı artırır)
		models := []mongo.WriteModel{
		mongo.NewInsertOneModel().SetDocument(Movie{Title: "Sadık", Year: 2024, Genre: "Dram"}),
		mongo.NewUpdateManyModel().SetFilter(bson.D{{"year", 1999}}).SetUpdate(bson.D{{"$set", bson.D{{"year", 2025}}}}),
		}
		opts := options.BulkWrite().SetOrdered(true) // sıralı yap işlemleri
*/

/*
	TODO : Belge sayıları:
		Bir koleksiyondaki belge sayısı hakkında yaklaşık bir bilgi edinmek için
		EstimatedDocumentCount() metodunu kullanabilirsiniz.
		Tam belge sayısı için ise CountDocuments() metodunu kullanabilirsiniz.
*/

/*
	TODO : Distinc
		filter := bson.D{{"title", "Back to the Future"}}
		//Distinct() ile "title" alanındaki benzersiz değerleri al
		results, err := coll.Distinct(context(), "year", filter)
		// Bu kod, MongoDB koleksiyonunda "title" değeri "Back to the Future" olan tüm belgeleri bulur
		//Sonra bu belgelerin "year" alanlarındaki tekrar etmeyen farklı değerlerin listesini alır.
*/

/*
	TODO :
		Kriterleri gerçek değerlerle eşleştirmek için aşağıdaki formatı kullanın:
		filter := bson.D{{"<field>", "<value>"}}
		|
		Kriterleri bir sorgu operatörüyle eşleştirmek için aşağıdaki biçimi kullanın:
		filter := bson.D{{"<field>", bson.D{{"<operator>", "<value>"}}}}
*/

/*
	TODO :
		Bu CountOptionstür, seçenekleri aşağıdaki yöntemlerle yapılandırmanıza olanak tanır:
			SetCollation() -> Sonuçları sıralarken kullanılacak dil sıralama türü. Varsayılan:nil
			SetHint()      ->Sayım yapılacak belgeleri taramak için kullanılacak dizin.Varsayılan:nil
			SetLimit()     ->Sayılacak maksimum belge sayısı. Varsayılan:0
			SetMaxTime()   ->Sorgunun sunucuda çalışabileceği maksimum süre. Varsayılan:nil
		    SetSkip()      ->Sayılmadan önce atlanacak belge sayısı. Varsayılan:0

*/

/*
	TODO :
		$match     :
		$count	   :
*/

/*
	TODO:
			results, err := coll.Distinct(context, "department", bson.D{{"enrollment", bson.D{{"$lt", 50}}}})
				opts := options.Find().SetSort(bson.D{{"enrollment", 1}}) // -1 de azalan yapar
				|
				filter := bson.D{{"enrollment", bson.D{{"$gt", 20}}}}
				opts := options.Find().SetLimit(2)
				|
				opts := options.Find().SetSort(bson.D{{"enrollment", 1}}).SetSkip(1).SetLimit(2)
				|
				bson.D{{"$limit", 3}}
				|
				opts := options.Find().SetProjection(bson.D{{"course_id", 0}, {"enrollment", 0}})
				|
				bson.D{{"$project", bson.D{{"title", 1}}}}
				|
				model := mongo.IndexModel{Keys: bson.D{{"description", "text"}}}
				|
				 bson.D{{"$text", bson.D{{"$search", "SERVES fish"}}}} // bosluk veya gibi calısır buyuk kucuk harf duyarlıgı yoktur
				|
				bson.D{{"$text", bson.D{{"$search", "\"serves 2\""}}}} // bu sadece serves 2 yı arar normalde bosluk or yerıne gecıyordu oradan kacmıs olduk
				|
				 bson.D{{"$text", bson.D{{"$search", "vegan -tofu"}}}}
				|
				filter := bson.D{{"$text", bson.D{{"$search", "vegetarian"}}}}
				sort := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
				projection := bson.D{{"name", 1}, {"description", 1}, {"score", bson.D{{"$meta", "textScore"}}}, {"_id", 0}}
				opts := options.Find().SetSort(sort).SetProjection(projection)
				|
				SetHint()  :
				|
				SetCollation()
				|
				bson.D{{"$inc", bson.D{{"sizes.$", -2}}}}
				|
				SetReturnDocument(options.After)
				|
				 bson.D{{"$unset", bson.D{{"styles.$[hotOptions]", ""}}}}
					SetArrayFilters(options.ArrayFilters{Filters: identifier}).
					SetReturnDocument(options.After)
				|
				bson.D{{"$mul", bson.D{{"sizes.$[]", 29.57}}}}
				|
				bson.D{{"$set", bson.D{{"species", "Ledebouria socialis"}, {"plant_id", 5}, {"height", 8.3}}}}
				|
				BulkWrite()
				|
				FindOneAndDelete
				FindOneAndUpdate
				FindOneAndReplace
				|
				Aggregation
				$group
					matchStage := bson.D{{"$match", bson.D{{"toppings", "milk foam"}}}}   // filtreleme yaptık
					unsetStage := bson.D{{"$unset", bson.A{"_id", "category"}}}           // belşirli alanları çıakrtık
					sortStage := bson.D{{"$sort", bson.D{{"price", 1}, {"toppings", 1}}}} // sıralama yaptık
					limitStage := bson.D{{"$limit", 2}}                                   // ilk iki belgeyi göstermek icin
				|
				indexModel := mongo.IndexModel{
				Keys: bson.D{{"title", 1}},
				}
				|
				indexModel := mongo.IndexModel{
					Keys: bson.D{
						{"fullplot", -1},
						{"title", 1},
					},
				}
				|
					cio := bson.D{{"key", bson.D{{"_id", 1}}}, {"unique", true}}
				opts := options.CreateCollection().SetClusteredIndex(cio)
				db.CreateCollection(context, "tea4", opts)
				|
				indexModel := mongo.IndexModel{
					Keys:    bson.D{{"theaterId", -1}},
					Options: options.Index().SetUnique(true), // Benzersiz bir dizin oluşturmak için, çoğaltılmasını engellemek istediğiniz alanı veya alan birleşimini belirtin ve uniqueseçeneği olarak ayarlayın true.
				}
				name, err := coll.Indexes().CreateOne(context, indexModel)
				|
				Transactions
				session, err := client.StartSession()
					if err != nil {
						log.Fatalf("Oturum başlatılamadı: %v", err)
					}
					defer session.EndSession(context.())
					   . Oturum Başlatma
					   MongoDB'de işlemler oturumlar içinde çalışır. Bir oturum başlatmak için istemciden StartSession() yöntemi çağrılır:
				txnOptions := options.Transaction().SetWriteConcern(writeconcern.Majority())
				result, err := session.WithTransaction(context.(), func(ctx mongo.SessionContext) (interface{}, error) {
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
				|
				|
				$lookup :
*/
