# Go-MongoDB Rehberi

Bu rehber, Go programlama dili ile MongoDB veritabanı kullanımını detaylı bir şekilde açıklamaktadır. MongoDB, NoSQL veritabanı olarak belge tabanlı bir yapıya sahiptir ve Go dili ile kullanımı oldukça yaygındır.

## İçindekiler

1. [Karşılaştırma Sorgu Operatörleri](#karşılaştırma-sorgu-operatörleri)
2. [Mantıksal Sorgu Operatörleri](#mantıksal-sorgu-operatörleri)
3. [Eleman Sorgu Operatörleri](#eleman-sorgu-operatörleri)
4. [Değerlendirme Sorgu Operatörleri](#değerlendirme-sorgu-operatörleri)
5. [Dizi Sorgu Operatörleri](#dizi-sorgu-operatörleri)
6. [Alan Güncelleme Operatörleri](#alan-güncelleme-operatörleri)
7. [Dizi Güncelleme Operatörleri](#dizi-güncelleme-operatörleri)
8. [Projeksiyon Operatörleri](#projeksiyon-operatörleri)
9. [Çeşitli Sorgu Operatörleri](#çeşitli-sorgu-operatörleri)
10. [Örnek Kodlar ve Kullanımlar](#örnek-kodlar-ve-kullanımlar)

## Karşılaştırma Sorgu Operatörleri

MongoDB'de veri filtreleme için kullanılan temel operatörlerdir. Bu operatörler, belgeleri belirli kriterlere göre filtrelemek için kullanılır.

### Operatörler ve Açıklamaları:

- `$eq` (Eşittir): Belirtilen değere tam olarak eşit olan değerleri bulur. Örneğin, yaşı 25 olan kullanıcıları bulmak için kullanılır.
- `$gt` (Büyüktür): Belirtilen değerden büyük olan değerleri bulur. Örneğin, fiyatı 100'den büyük olan ürünleri bulmak için kullanılır.
- `$gte` (Büyük Eşittir): Belirtilen değerden büyük veya eşit olan değerleri bulur. Örneğin, yaşı 18 ve üzeri olan kullanıcıları bulmak için kullanılır.
- `$in` (İçinde): Belirtilen değerler listesinden herhangi biriyle eşleşen değerleri bulur. Örneğin, belirli kategorilerdeki ürünleri bulmak için kullanılır.
- `$lt` (Küçüktür): Belirtilen değerden küçük olan değerleri bulur. Örneğin, stok miktarı 10'dan az olan ürünleri bulmak için kullanılır.
- `$lte` (Küçük Eşittir): Belirtilen değerden küçük veya eşit olan değerleri bulur. Örneğin, fiyatı 50 ve altı olan ürünleri bulmak için kullanılır.
- `$ne` (Eşit Değildir): Belirtilen değere eşit olmayan tüm değerleri bulur. Örneğin, aktif olmayan kullanıcıları bulmak için kullanılır.
- `$nin` (İçinde Değildir): Belirtilen değerler listesinde olmayan değerleri bulur. Örneğin, belirli kategorilerde olmayan ürünleri bulmak için kullanılır.

### Örnek Kullanımlar:

```go
// $eq örneği: Yaşı 25 olan kullanıcıları bul
filter := bson.D{{"age", bson.D{{"$eq", 25}}}}

// $gt ve $lt örneği: Fiyatı 100 ile 200 arasında olan ürünleri bul
filter := bson.D{{"price", bson.D{{"$gt", 100}, {"$lt", 200}}}}

// $in örneği: Belirli kategorilerdeki ürünleri bul
filter := bson.D{{"category", bson.D{{"$in", []string{"elektronik", "giyim", "kitap"}}}}}

// $nin örneği: Belirli durumlarda olmayan kullanıcıları bul
filter := bson.D{{"status", bson.D{{"$nin", []string{"silindi", "pasif"}}}}}
```

## Mantıksal Sorgu Operatörleri

Karmaşık sorgular oluşturmak için kullanılan operatörlerdir. Bu operatörler, birden fazla koşulu birleştirmek için kullanılır.

### Operatörler ve Açıklamaları:

- `$and`: Tüm koşulların aynı anda sağlanması gereken sorgular için kullanılır. Örneğin, hem yaşı 18'den büyük hem de aktif olan kullanıcıları bulmak için kullanılır.
- `$or`: Koşullardan herhangi birinin sağlanması yeterli olan sorgular için kullanılır. Örneğin, yaşı 18'den küçük veya aktif olmayan kullanıcıları bulmak için kullanılır.
- `$not`: Sorgu sonucunu tersine çeviren operatördür. Örneğin, yaşı 18'den büyük olmayan kullanıcıları bulmak için kullanılır.
- `$nor`: Hiçbir koşulun sağlanmaması gereken sorgular için kullanılır. Örneğin, ne silinmiş ne de pasif olan kullanıcıları bulmak için kullanılır.

### Örnek Kullanımlar:

```go
// $and örneği: Hem yaşı 18'den büyük hem de aktif olan kullanıcıları bul
filter := bson.D{{"$and", bson.A{
    bson.D{{"age", bson.D{{"$gt", 18}}}},
    bson.D{{"status", "active"}},
}}}

// $or örneği: Yaşı 18'den küçük veya aktif olmayan kullanıcıları bul
filter := bson.D{{"$or", bson.A{
    bson.D{{"age", bson.D{{"$lt", 18}}}},
    bson.D{{"status", bson.D{{"$ne", "active"}}}},
}}}

// $not örneği: Yaşı 18'den büyük olmayan kullanıcıları bul
filter := bson.D{{"age", bson.D{{"$not", bson.D{{"$gt", 18}}}}}}

// $nor örneği: Ne silinmiş ne de pasif olan kullanıcıları bul
filter := bson.D{{"$nor", bson.A{
    bson.D{{"status", "deleted"}},
    bson.D{{"status", "inactive"}},
}}}
```

## Eleman Sorgu Operatörleri

Belge yapısını kontrol etmek için kullanılan operatörlerdir. Bu operatörler, belgelerin yapısını ve içeriğini kontrol etmek için kullanılır.

### Operatörler ve Açıklamaları:

- `$exists`: Belirtilen alanın belgede var olup olmadığını kontrol eder. Örneğin, email alanı olan kullanıcıları bulmak için kullanılır.
- `$type`: Alanın veri tipini kontrol eder. Örneğin, yaş alanı sayı tipinde olan kullanıcıları bulmak için kullanılır.

### Örnek Kullanımlar:

```go
// $exists örneği: Email alanı olan kullanıcıları bul
filter := bson.D{{"email", bson.D{{"$exists", true}}}}

// $type örneği: Yaş alanı sayı tipinde olan kullanıcıları bul
filter := bson.D{{"age", bson.D{{"$type", "number"}}}}
```

## Değerlendirme Sorgu Operatörleri

Karmaşık değerlendirmeler için kullanılan operatörlerdir. Bu operatörler, belgeler üzerinde daha karmaşık işlemler yapmak için kullanılır.

### Operatörler ve Açıklamaları:

- `$expr`: Sorgu içinde ifadeler kullanmayı sağlar. Örneğin, fiyatı maliyetinden yüksek olan ürünleri bulmak için kullanılır.
- `$jsonSchema`: JSON şemasına göre doğrulama yapar. Örneğin, belirli bir şemaya uyan belgeleri bulmak için kullanılır.
- `$mod`: Modül işlemi yapar. Örneğin, yaşı çift olan kullanıcıları bulmak için kullanılır.
- `$regex`: Düzenli ifade ile eşleştirme yapar. Örneğin, ismi "A" ile başlayan kullanıcıları bulmak için kullanılır.
- `$where`: JavaScript ifadeleri kullanmayı sağlar. Örneğin, karmaşık hesaplamalar gerektiren sorgular için kullanılır.

### Örnek Kullanımlar:

```go
// $expr örneği: Fiyatı maliyetinden yüksek olan ürünleri bul
filter := bson.D{{"$expr", bson.D{{"$gt", []string{"$price", "$cost"}}}}}

// $regex örneği: İsmi "A" ile başlayan kullanıcıları bul
filter := bson.D{{"name", bson.D{{"$regex", "^A"}}}}

// $mod örneği: Yaşı çift olan kullanıcıları bul
filter := bson.D{{"age", bson.D{{"$mod", []int{2, 0}}}}}
```

## Dizi Sorgu Operatörleri

Dizi alanları üzerinde işlem yapmak için kullanılan operatörlerdir. Bu operatörler, dizi tipindeki alanlar üzerinde işlem yapmak için kullanılır.

### Operatörler ve Açıklamaları:

- `$all`: Belirtilen tüm öğeleri içeren dizileri bulur. Örneğin, hem "go" hem de "mongodb" etiketlerine sahip belgeleri bulmak için kullanılır.
- `$elemMatch`: Dizi içindeki öğelerin belirli koşulları sağlaması gereken belgeleri bulur. Örneğin, matematik dersinden 90'dan yüksek not alan öğrencileri bulmak için kullanılır.
- `$size`: Belirli boyuttaki dizileri bulur. Örneğin, 3 etikete sahip belgeleri bulmak için kullanılır.

### Örnek Kullanımlar:

```go
// $all örneği: Hem "go" hem de "mongodb" etiketlerine sahip belgeleri bul
filter := bson.D{{"tags", bson.D{{"$all", []string{"go", "mongodb"}}}}}

// $elemMatch örneği: Matematik dersinden 90'dan yüksek not alan öğrencileri bul
filter := bson.D{{"scores", bson.D{{"$elemMatch", bson.D{
    {"subject", "matematik"},
    {"score", bson.D{{"$gt", 90}}},
}}}}}

// $size örneği: 3 etikete sahip belgeleri bul
filter := bson.D{{"tags", bson.D{{"$size", 3}}}}
```

## Projeksiyon Operatörleri

Sorgu sonuçlarını şekillendirmek için kullanılan operatörlerdir. Bu operatörler, dönen belgelerin hangi alanlarını görmek istediğimizi belirlemek için kullanılır.

### Operatörler ve Açıklamaları:

- `$`: Sorgu koşuluyla eşleşen dizideki ilk öğeyi yansıtır. Örneğin, belirli bir koşula uyan ilk yorumu göstermek için kullanılır.
- `$elemMatch`: Belirtilen koşulla eşleşen dizideki ilk öğeyi yansıtır. Örneğin, belirli bir koşula uyan ilk skoru göstermek için kullanılır.
- `$meta`: Meta verileri yansıtır. Örneğin, metin arama sonuçlarında puanı göstermek için kullanılır.
- `$slice`: Diziden belirli sayıda öğe yansıtır. Örneğin, son 5 yorumu göstermek için kullanılır.

### Örnek Kullanımlar:

```go
// $slice örneği: Son 5 yorumu göster
projection := bson.D{{"comments", bson.D{{"$slice", 5}}}}

// $meta örneği: Metin arama sonuçlarında puanı göster
projection := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
```

## Çeşitli Sorgu Operatörleri

Özel durumlar için kullanılan operatörlerdir. Bu operatörler, özel durumlarda kullanılmak üzere tasarlanmıştır.

### Operatörler ve Açıklamaları:

- `$rand`: Rastgele değer üretir. Örneğin, rastgele belgeler seçmek için kullanılır.
- `$natural`: Doğal sıralama için kullanılır. Örneğin, belgeleri doğal sıralamada göstermek için kullanılır.

### Örnek Kullanımlar:

```go
// $rand örneği: Rastgele belgeler seç
filter := bson.D{{"$expr", bson.D{{"$gt", []interface{}{"$random", 0.5}}}}}

// $natural örneği: Belgeleri doğal sıralamada göster
opts := options.Find().SetHint(bson.D{{"$natural", 1}})
```

## Alan Güncelleme Operatörleri

Belge güncellemeleri için kullanılan operatörler:

- `$set`: Belge alanlarını günceller
- `$inc`: Sayısal değerleri artırır/azaltır
- `$mul`: Sayısal değerleri çarpar
- `$rename`: Alan adlarını değiştirir
- `$unset`: Alanları siler

### Örnek Kullanımlar:

```go
// $set örneği
update := bson.D{{"$set", bson.D{
    {"status", "updated"},
    {"lastModified", time.Now()},
}}}

// $inc örneği
update := bson.D{{"$inc", bson.D{
    {"viewCount", 1},
    {"score", 5},
}}}
```

## Dizi Güncelleme Operatörleri

Dizi alanlarını güncellemek için kullanılan operatörler:

- `$`: Eşleşen ilk öğeyi günceller
- `$[]`: Tüm eşleşen öğeleri günceller
- `$[<identifier>]`: Belirli koşullara göre öğeleri günceller
- `$addToSet`: Tekrarsız öğe ekler
- `$pop`: İlk veya son öğeyi kaldırır
- `$pull`: Koşula uyan öğeleri kaldırır
- `$push`: Yeni öğe ekler
- `$pullAll`: Belirtilen değerleri kaldırır
- `$each`: Birden fazla öğe ekler
- `$position`: Ekleme konumunu belirler
- `$slice`: Dizi boyutunu sınırlar
- `$sort`: Diziyi sıralar

### Örnek Kullanımlar:

```go
// $push ve $each örneği
update := bson.D{{"$push", bson.D{
    {"tags", bson.D{
        {"$each", []string{"yeni", "etiketler"}},
        {"$slice", 5},
    }},
}}}

// $pull örneği
update := bson.D{{"$pull", bson.D{{"tags", "eski"}}}}

// $addToSet örneği
update := bson.D{{"$addToSet", bson.D{{"categories", "yeni-kategori"}}}}
```

## Metin Arama ve İndeksleme

```go
// Metin indeksi oluşturma
indexModel := mongo.IndexModel{
    Keys: bson.D{{"description", "text"}},
}
name, err := collection.Indexes().CreateOne(context.TODO(), indexModel)

// Metin arama örneği
filter := bson.D{{"$text", bson.D{{"$search", "arama terimi"}}}}
```

## Agregasyon Pipeline

```go
pipeline := mongo.Pipeline{
    {{"$match", bson.D{{"status", "active"}}}},
    {{"$group", bson.D{
        {"_id", "$category"},
        {"total", bson.D{{"$sum", 1}}},
    }}},
    {{"$sort", bson.D{{"total", -1}}}},
    {{"$limit", 10}},
}
cursor, err := collection.Aggregate(context.TODO(), pipeline)
```

## İşlemler (Transactions)

```go
session, err := client.StartSession()
if err != nil {
    log.Fatal(err)
}
defer session.EndSession(context.TODO())

err = session.StartTransaction()
if err != nil {
    log.Fatal(err)
}

// İşlem içinde işlemler
_, err = collection.InsertOne(session, bson.D{{"name", "Ali"}})
if err != nil {
    session.AbortTransaction(session)
    log.Fatal(err)
}

err = session.CommitTransaction(session)
if err != nil {
    log.Fatal(err)
}
```

## Belge Sayma ve Benzersiz Değerler

```go
// Yaklaşık belge sayısı
count, err := collection.EstimatedDocumentCount(context.TODO())

// Tam belge sayısı
count, err := collection.CountDocuments(context.TODO(), filter)

// Benzersiz değerler
results, err := collection.Distinct(context.TODO(), "field", filter)
```

## ObjectId Dönüşümleri

MongoDB'de ObjectId ile string arasında dönüşüm yapmak için kullanılan metodlar:

```go
// String'den ObjectId'ye dönüşüm
objectId, err := primitive.ObjectIDFromHex("675d3a322979f406206c8341")
if err != nil {
    log.Fatal(err)
}

// ObjectId'den String'e dönüşüm
stringID := objectId.Hex()
```

## CountOptions Yapılandırması

Belge sayma işlemlerini özelleştirmek için kullanılan seçenekler:

```go
opts := options.Count().
    SetCollation(&options.Collation{Locale: "tr"}).  // Dil sıralama türü
    SetHint(bson.D{{"indexName", 1}}).              // Kullanılacak indeks
    SetLimit(100).                                   // Maksimum belge sayısı
    SetMaxTime(5 * time.Second).                     // Maksimum çalışma süresi
    SetSkip(10)                                      // Atlanacak belge sayısı

count, err := collection.CountDocuments(context.TODO(), filter, opts)
```

## Distinct Kullanımı

Belirli bir alandaki benzersiz değerleri almak için:

```go
// Belirli bir filtreye göre benzersiz değerleri alma
filter := bson.D{{"title", "Back to the Future"}}
results, err := collection.Distinct(context.TODO(), "year", filter)

// Filtresiz benzersiz değerleri alma
results, err := collection.Distinct(context.TODO(), "department", nil)
```

## Find Options

Sorgu sonuçlarını özelleştirmek için kullanılan seçenekler:

```go
// Sıralama
opts := options.Find().SetSort(bson.D{{"enrollment", 1}})  // 1: artan, -1: azalan

// Limit ve Skip
opts := options.Find().
    SetLimit(2).                    // İlk 2 belge
    SetSkip(1)                      // İlk belgeyi atla

// Projeksiyon (Hangi alanların döneceğini belirleme)
opts := options.Find().SetProjection(bson.D{
    {"course_id", 0},              // 0: hariç tut, 1: dahil et
    {"enrollment", 0},
})

// Birleşik kullanım
opts := options.Find().
    SetSort(bson.D{{"enrollment", 1}}).
    SetSkip(1).
    SetLimit(2).
    SetProjection(bson.D{{"title", 1}})
```

## Metin Arama Özellikleri

```go
// Metin indeksi oluşturma
model := mongo.IndexModel{
    Keys: bson.D{{"description", "text"}},
}
name, err := collection.Indexes().CreateOne(context.TODO(), model)

// Basit metin arama
filter := bson.D{{"$text", bson.D{{"$search", "SERVES fish"}}}}

// Tam ifade arama
filter := bson.D{{"$text", bson.D{{"$search", "\"serves 2\""}}}}

// Hariç tutma ile arama
filter := bson.D{{"$text", bson.D{{"$search", "vegan -tofu"}}}}

// Metin arama sonuçlarını sıralama ve puanlama
filter := bson.D{{"$text", bson.D{{"$search", "vegetarian"}}}}
sort := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
projection := bson.D{
    {"name", 1},
    {"description", 1},
    {"score", bson.D{{"$meta", "textScore"}}},
    {"_id", 0},
}
opts := options.Find().SetSort(sort).SetProjection(projection)
```

## Dizi Güncelleme Detayları

```go
// Belirli bir dizi öğesini güncelleme
update := bson.D{{"$inc", bson.D{{"sizes.$", -2}}}}

// Dizi filtreleri ile güncelleme
identifier := bson.D{{"hotOptions", bson.D{{"$gt", 100}}}}
update := bson.D{{"$unset", bson.D{{"styles.$[hotOptions]", ""}}}}
opts := options.Update().
    SetArrayFilters(options.ArrayFilters{Filters: identifier}).
    SetReturnDocument(options.After)

// Tüm dizi öğelerini güncelleme
update := bson.D{{"$mul", bson.D{{"sizes.$[]", 29.57}}}}
```

## Koleksiyon Oluşturma ve İndeksleme

```go
// Kümelenmiş indeks ile koleksiyon oluşturma
cio := bson.D{{"key", bson.D{{"_id", 1}}}, {"unique", true}}
opts := options.CreateCollection().SetClusteredIndex(cio)
db.CreateCollection(context.TODO(), "yeniKoleksiyon", opts)

// Benzersiz indeks oluşturma
indexModel := mongo.IndexModel{
    Keys:    bson.D{{"theaterId", -1}},
    Options: options.Index().SetUnique(true),
}
name, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
```

## Agregasyon Pipeline Detayları

MongoDB'de agregasyon pipeline'ı, belgeleri işlemek ve dönüştürmek için kullanılan güçlü bir araçtır. Pipeline, belgeleri bir dizi aşamadan geçirerek istenen sonuçları elde etmenizi sağlar.

### Pipeline Operatörleri ve Açıklamaları:

1. **$match**: Belgeleri filtrelemek için kullanılır. SQL'deki WHERE koşuluna benzer.
   - Belirli kriterlere uyan belgeleri seçer
   - Pipeline'ın başında kullanıldığında performansı artırır
   - Birden fazla koşul birleştirilebilir

2. **$unset**: Belirtilen alanları belgelerden kaldırır.
   - Belge yapısını sadeleştirmek için kullanılır
   - Birden fazla alan aynı anda kaldırılabilir
   - İç içe alanlar da kaldırılabilir

3. **$sort**: Belgeleri belirtilen alanlara göre sıralar.
   - 1: Artan sıralama (A'dan Z'ye)
   - -1: Azalan sıralama (Z'den A'ya)
   - Birden fazla alana göre sıralama yapılabilir

4. **$limit**: Dönen belge sayısını sınırlar.
   - Performansı artırmak için kullanılır
   - Genellikle $sort ile birlikte kullanılır
   - Sayfalama için kullanılabilir

5. **$group**: Belgeleri gruplar ve her grup için hesaplamalar yapar.
   - _id: Gruplama yapılacak alan
   - $sum: Toplam hesaplama
   - $avg: Ortalama hesaplama
   - $min: Minimum değer
   - $max: Maximum değer
   - $first: Gruptaki ilk değer
   - $last: Gruptaki son değer
   - $push: Değerleri dizi olarak toplar
   - $addToSet: Tekrarsız değerleri dizi olarak toplar

6. **$lookup**: Başka bir koleksiyonla birleştirme (join) yapar.
   - from: Birleştirilecek koleksiyon
   - localField: Mevcut koleksiyondaki alan
   - foreignField: Birleştirilecek koleksiyondaki alan
   - as: Sonuçların ekleneceği alan adı

### Örnek Kullanımlar:

```go
pipeline := mongo.Pipeline{
    // 1. Filtreleme: Sadece "milk foam" içeren ürünleri seç
    {{"$match", bson.D{{"toppings", "milk foam"}}}},
    
    // 2. Alanları çıkarma: _id ve category alanlarını kaldır
    {{"$unset", bson.A{"_id", "category"}}}},
    
    // 3. Sıralama: Önce fiyata göre artan, sonra toppings'e göre artan sırala
    {{"$sort", bson.D{{"price", 1}, {"toppings", 1}}}},
    
    // 4. Limit: İlk 2 belgeyi al
    {{"$limit", 2}},
    
    // 5. Gruplama: Kategoriye göre grupla ve istatistikler hesapla
    {{"$group", bson.D{
        {"_id", "$category"},
        {"total", bson.D{{"$sum", 1}}},           // Toplam belge sayısı
        {"avgPrice", bson.D{{"$avg", "$price"}}}, // Ortalama fiyat
        {"minPrice", bson.D{{"$min", "$price"}}}, // Minimum fiyat
        {"maxPrice", bson.D{{"$max", "$price"}}}, // Maximum fiyat
        {"firstProduct", bson.D{{"$first", "$name"}}}, // İlk ürün adı
        {"allProducts", bson.D{{"$push", "$name"}}},   // Tüm ürün adları
        {"uniqueToppings", bson.D{{"$addToSet", "$toppings"}}}, // Tekrarsız toppings
    }}},
    
    // 6. Lookup: Kategoriler koleksiyonu ile birleştir
    {{"$lookup", bson.D{
        {"from", "categories"},                    // Birleştirilecek koleksiyon
        {"localField", "categoryId"},             // Mevcut koleksiyondaki alan
        {"foreignField", "_id"},                  // Birleştirilecek koleksiyondaki alan
        {"as", "categoryInfo"},                   // Sonuçların ekleneceği alan
    }}},
    
    // 7. Projeksiyon: Sadece istenen alanları göster
    {{"$project", bson.D{
        {"_id", 0},                              // _id alanını gösterme
        {"category", 1},                         // category alanını göster
        {"total", 1},                           // total alanını göster
        {"avgPrice", 1},                        // avgPrice alanını göster
        {"categoryInfo.name", 1},               // categoryInfo içindeki name alanını göster
    }}},
    
    // 8. Skip: İlk 5 sonucu atla (sayfalama için)
    {{"$skip", 5}},
    
    // 9. Facet: Farklı hesaplamaları aynı anda yap
    {{"$facet", bson.D{
        {"priceStats", bson.A{
            bson.D{{"$group", bson.D{
                {"_id", nil},
                {"avgPrice", bson.D{{"$avg", "$price"}}},
                {"minPrice", bson.D{{"$min", "$price"}}},
                {"maxPrice", bson.D{{"$max", "$price"}}},
            }}},
        }},
        {"categoryStats", bson.A{
            bson.D{{"$group", bson.D{
                {"_id", "$category"},
                {"count", bson.D{{"$sum", 1}}},
            }}},
        }},
    }}},
}

// Pipeline'ı çalıştır
cursor, err := collection.Aggregate(context.TODO(), pipeline)
if err != nil {
    log.Fatal(err)
}
defer cursor.Close(context.TODO())

// Sonuçları işle
var results []bson.M
if err = cursor.All(context.TODO(), &results); err != nil {
    log.Fatal(err)
}

// Sonuçları yazdır
for _, result := range results {
    fmt.Printf("%+v\n", result)
}
```

### Diğer Önemli Pipeline Operatörleri:

1. **$project**: Belge yapısını yeniden şekillendirir.
   - Alanları yeniden adlandırma
   - Yeni alanlar oluşturma
   - Alanları kaldırma veya ekleme

2. **$skip**: Belirtilen sayıda belgeyi atlar.
   - Sayfalama için kullanılır
   - $limit ile birlikte kullanılabilir

3. **$facet**: Birden fazla agregasyon pipeline'ını paralel olarak çalıştırır.
   - Farklı istatistikler hesaplamak için kullanılır
   - Tek sorguda birden fazla sonuç döndürür

4. **$addFields**: Yeni alanlar ekler veya mevcut alanları günceller.
   - Hesaplanmış alanlar eklemek için kullanılır
   - Mevcut alanları değiştirmek için kullanılır

5. **$replaceRoot**: Belge yapısını değiştirir.
   - İç içe belgeleri düzleştirmek için kullanılır
   - Belge yapısını yeniden düzenlemek için kullanılır

### Performans İpuçları:

1. Pipeline'ın başında $match kullanın
2. Gereksiz alanları $unset ile kaldırın
3. İndeksleri doğru kullanın
4. $limit ve $skip'i doğru sırada kullanın
5. Büyük veri setlerinde $facet kullanırken dikkatli olun

Bu detaylı açıklamalar ve örnekler, MongoDB agregasyon pipeline'ını daha etkili kullanmanıza yardımcı olacaktır. Her operatörün ne zaman ve nasıl kullanılacağını anlamak, karmaşık sorguları daha verimli bir şekilde yazmanızı sağlar.

Bu rehber, Go ile MongoDB kullanımının tüm temel ve ileri düzey özelliklerini kapsamaktadır. Her operatör ve metod için detaylı açıklamalar ve örnekler eklenmiştir. Daha fazla bilgi için MongoDB resmi dokümantasyonunu inceleyebilirsiniz.
