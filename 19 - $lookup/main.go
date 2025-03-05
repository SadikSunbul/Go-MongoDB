package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Kullanıcı
type User struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

// Profil (1'e 1)
type Profile struct {
	UserID  string `bson:"user_id"`
	Age     int    `bson:"age"`
	Address string `bson:"address"`
}

// Gönderi (1'e çok)
type Post struct {
	ID      string `bson:"_id"`
	UserID  string `bson:"user_id"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
}

// Grup
type Group struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

// Ara Koleksiyon (User-Groups, Many-to-Many)
type UserGroup struct {
	UserID  string `bson:"user_id"`
	GroupID string `bson:"group_id"`
}

var client *mongo.Client
var usersCollection *mongo.Collection
var profilesCollection *mongo.Collection
var postsCollection *mongo.Collection
var groupsCollection *mongo.Collection
var user_groupsCollection *mongo.Collection

func connectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	usersCollection = client.Database("testdb").Collection("users")
	profilesCollection = client.Database("testdb").Collection("profiles")
	postsCollection = client.Database("testdb").Collection("posts")
	groupsCollection = client.Database("testdb").Collection("groups")
	user_groupsCollection = client.Database("testdb").Collection("user_groups")
	fmt.Println("MongoDB'ye bağlanıldı!")
}

func addUser(user User) {

	_, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Kullanıcı eklendi:", user.Name)
}

func addProfile(profile Profile) {

	_, err := profilesCollection.InsertOne(context.TODO(), profile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Profil eklendi:", profile.UserID)
}

func addPost(post Post) {

	_, err := postsCollection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Post eklendi:", post.Title)
}

func addGroup(group Group) {

	_, err := groupsCollection.InsertOne(context.TODO(), group)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Grup eklendi:", group.Name)
}

func addUserToGroup(userGroup UserGroup) {

	_, err := user_groupsCollection.InsertOne(context.TODO(), userGroup)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Kullanıcı gruba eklendi:", userGroup.UserID, "->", userGroup.GroupID)
}

func Insert() {
	// Kullanıcılar ekleyelim
	users := []User{
		{ID: "1", Name: "Ahmet", Email: "ahmet@example.com"},
		{ID: "2", Name: "Mehmet", Email: "mehmet@example.com"},
		{ID: "3", Name: "Zeynep", Email: "zeynep@example.com"},
	}
	for _, user := range users {
		addUser(user)
	}

	// Profiller ekleyelim
	profiles := []Profile{
		{UserID: "1", Age: 25, Address: "İstanbul"},
		{UserID: "2", Age: 30, Address: "Ankara"},
		{UserID: "3", Age: 22, Address: "İzmir"},
	}
	for _, profile := range profiles {
		addProfile(profile)
	}

	// Postlar ekleyelim
	posts := []Post{
		{ID: "101", UserID: "1", Title: "MongoDB Öğreniyorum", Content: "Bugün MongoDB çalıştım!"},
		{ID: "102", UserID: "1", Title: "GoLang Harika", Content: "Go dilini öğrenmeye başladım!"},
		{ID: "103", UserID: "2", Title: "Fiber Framework", Content: "Go'da Fiber kullanarak API geliştirdim."},
		{ID: "104", UserID: "3", Title: "Golang Concurrency", Content: "Goroutine ve Channel'ları inceledim."},
	}
	for _, post := range posts {
		addPost(post)
	}

	// Gruplar ekleyelim
	groups := []Group{
		{ID: "201", Name: "Golang Severler"},
		{ID: "202", Name: "MongoDB Kullanıcıları"},
	}
	for _, group := range groups {
		addGroup(group)
	}

	// Kullanıcıları gruplara ekleyelim (Çoktan çoğa ilişkiler)
	userGroups := []UserGroup{
		{UserID: "1", GroupID: "201"},
		{UserID: "1", GroupID: "202"},
		{UserID: "2", GroupID: "201"},
		{UserID: "3", GroupID: "202"},
	}
	for _, ug := range userGroups {
		addUserToGroup(ug)
	}
}

func main() {
	connectDB()

	query9()
}
func query1() {
	// Ahmet’in katıldığı tüm grupları getir.

	pipeline := mongo.Pipeline{
		{{"$match", bson.M{"name": "Ahmet"}}}, // "Ahmet" adlı kullanıcıyı bul
		{{"$lookup", bson.M{
			"from":         "user_groups", // user_groups koleksiyonuyla ilişkilendir
			"localField":   "_id",         // users koleksiyonundaki _id alanı
			"foreignField": "user_id",     // user_groups koleksiyonundaki user_id alanı
			"as":           "user_groups", // Elde edilen sonuçları "user_groups" dizisine ekle
		}}},
		{{Key: "$unwind", Value: "$user_groups"}}, // user_groups dizisini açarak her üyeliği ayrı bir satır yap
		{{Key: "$lookup", Value: bson.M{
			"from":         "groups",               // groups koleksiyonunu bağla
			"localField":   "user_groups.group_id", // user_groups içindeki group_id ile eşleştir
			"foreignField": "_id",                  // groups koleksiyonundaki _id ile eşleştir
			"as":           "group_info",           // Elde edilen sonuçları "group_info" dizisine ekle
		}}},
		{{Key: "$unwind", Value: "$group_info"}}, // group_info dizisini açarak her grubu ayrı bir satır yap
		{{Key: "$project", Value: bson.M{
			"_id":   0,                  // _id değerini göstermeyelim
			"group": "$group_info.name", // Sonuç olarak sadece grup adını al
		}}},
	}

	cursor, err := usersCollection.Aggregate(context.TODO(), pipeline) // Pipeline'ı çalıştır ve cursor al
	if err != nil {
		log.Fatal(err) // Hata varsa çık
	}
	defer cursor.Close(context.TODO()) // İşlem bitince cursor'u kapat

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil { // Sonuçları results dizisine aktar
		log.Fatal(err)
	}

	fmt.Println("Ahmet ", "adlı kullanıcının katıldığı gruplar:")
	for _, result := range results { // Sonuçları ekrana yazdır
		fmt.Println(result["group"])
	}
}

func query2() {
	// "Golang Severler" grubuna üye olan tüm kullanıcıları listele.

	// group
	pipeline := mongo.Pipeline{{
		{"$match", bson.M{"name": "Golang Severler"}}},
		{{"$lookup", bson.M{
			"from":         "user_groups",
			"localField":   "_id",
			"foreignField": "group_id",
			"as":           "group_user",
		}}},
		{{Key: "$unwind", Value: "$group_user"}},
		{{"$lookup", bson.M{
			"from":         "users",
			"localField":   "group_user.user_id",
			"foreignField": "_id",
			"as":           "group_users_list",
		}}},
		{{Key: "$unwind", Value: "$group_users_list"}},
	}

	cursor, err := groupsCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		fmt.Errorf("errr:", err)
		return
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil { // Sonuçları results dizisine aktar
		log.Fatal(err)
	}
	fmt.Println("-----------------------")
	for _, result := range results { // Sonuçları ekrana yazdır
		fmt.Println(result["group_users_list"])
	}
	fmt.Println("-----------------------")
}

func query3() {
	// Ahmet’in yazdığı tüm gönderileri getir.

	// user ve post arasında kullanılıcak

	// user uszerınden ilerlenicek

	piplene := mongo.Pipeline{{
		{"$match", bson.M{"name": "Ahmet"}}},
		{{"$lookup", bson.M{
			"from":         "posts",
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "users_posts",
		}}},
		{{Key: "$unwind", Value: "$users_posts"}},
	}

	cursor, err := usersCollection.Aggregate(context.TODO(), piplene)
	if err != nil {
		fmt.Errorf("Err:", err)
		return
	}

	var result []bson.M

	if err = cursor.All(context.TODO(), &result); err != nil {
		fmt.Errorf("err:", err)
		return
	}

	fmt.Println("------------------------")
	for _, val := range result {
		res, _ := json.Marshal(val["users_posts"])
		fmt.Printf("%s \n", res)
	}

	fmt.Println("------------------------")

}

func query4() {
	// "MongoDB Kullanıcıları" grubunda kaç kişi var?

	// groups -> user_groups

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"name": "MongoDB Kullanıcıları"}}},
		bson.D{{
			Key: "$lookup",
			Value: bson.M{
				"from":         "user_groups",
				"localField":   "_id",
				"foreignField": "group_id",
				"as":           "select_group",
			},
		}},
		bson.D{{Key: "$unwind", Value: "$select_group"}},
		bson.D{{
			Key: "$group",
			Value: bson.M{
				"_id":         "$select_group.name", // Burayı uygun bir ID'ye çevirebilirsin
				"total_count": bson.M{"$sum": 1},
			},
		}},
		bson.D{{"$project", bson.M{"_id": 0}}},
	}

	cursor, err := groupsCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("count: %v \n", result)
}

func query5() {
	// Yaşı 25’ten büyük kullanıcıların gönderilerini getir.
	// users  --> profiles(yaş) --> posts
	pipeline := mongo.Pipeline{
		bson.D{{"$lookup", bson.M{
			"from":         "profiles",
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "user_profile",
		}}},
		bson.D{{"$unwind", bson.M{"path": "$user_profile", "preserveNullAndEmptyArrays": false}}},
		bson.D{{"$match", bson.M{"user_profile.age": bson.M{"$gt": 25}}}},
		bson.D{{"$lookup", bson.M{
			"from":         "posts",
			"localField":   "user_profile.user_id",
			"foreignField": "user_id",
			"as":           "users_posts",
		}}},
		bson.D{{"$unwind", bson.M{"path": "$users_posts", "preserveNullAndEmptyArrays": false}}},
		bson.D{{"$project", bson.M{"users_posts.title": 1, "users_posts.content": 1, "users_posts.user_id": 1}}},
	}

	cursor, err := usersCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var result []bson.M

	if err = cursor.All(context.TODO(), &result); err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("post:%s \n", result)
}

func query6() {
	// Her kullanıcının kaç tane gönderisi olduğunu listele.

	// users -> posts

	pipeline := mongo.Pipeline{
		bson.D{{"$lookup", bson.M{
			"from":         "posts",
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "user_post",
		}}},
		bson.D{{"$unwind", bson.M{"path": "$user_post", "preserveNullAndEmptyArrays": false}}},
		bson.D{{"$group", bson.D{
			{"_id", "$user_post.user_id"},
			{"total_post", bson.M{"$sum": 1}},
		}}},
	}

	cursor, err := usersCollection.Aggregate(context.TODO(), pipeline)

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var result []bson.M

	if err = cursor.All(context.TODO(), &result); err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("data:%s \n", result)

}

func query7() {
	//Ahmet’in yazdığı en son gönderiyi getir.

	// users (aheti bul) -> posts (en son olusturulan)

	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.M{"name": "Ahmet"}}},
		bson.D{{"$lookup", bson.M{
			"from":         "posts",
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "user_profile", // post orneklerini taşır
		}}},
		bson.D{{"$unwind", "$user_profile"}},
		bson.D{{"$sort", bson.M{"user_profile._id": -1}}}, // Eğer posts._id bir ObjectID ise, MongoDB'de ObjectID'ler oluşturulma zamanına göre sıralanır.
		bson.D{{"$limit", 1}},
		bson.D{{"$project", bson.M{"user_profile": 1}}},
	}

	cursor, err := usersCollection.Aggregate(context.TODO(), pipeline)

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var result []bson.M

	if err = cursor.All(context.TODO(), &result); err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("result: %s \n", result)

}

func query8() {
	// En çok gönderi yazan kullanıcıyı getir.

	// user -> gönderi
	pipeline := mongo.Pipeline{
		bson.D{{"$lookup", bson.M{
			"from":         "posts",
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "user_posts",
		}}},
		bson.D{{"$unwind", "$user_posts"}},
		bson.D{{"$group", bson.D{
			{"_id", "$user_posts.user_id"},
			{"total_posts", bson.D{{"$sum", 1}}},
		}}},
		bson.D{{"$sort", bson.M{"total_posts": -1}}},
		bson.D{{"$limit", 1}},
		bson.D{{"$project", bson.M{"user_posts.total_posts": 1, "users.name": 1}}},
	}

	cursro, err := usersCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var result []bson.M

	if err = cursro.All(context.TODO(), &result); err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("data:%s \n", result)
}

func query9() {
	// En çok gönderi yazan kullanıcıyı getir.
	pipeline := mongo.Pipeline{
		bson.D{{"$facet", bson.M{
			"user_posts": bson.A{
				bson.D{{"$lookup", bson.M{
					"from":         "posts",
					"localField":   "_id",
					"foreignField": "user_id",
					"as":           "posts",
				}}},
				bson.D{{"$unwind", "$posts"}},
				bson.D{{"$group", bson.D{
					{"_id", "$_id"},
					{"total_posts", bson.D{{"$sum", 1}}},
					{"user_name", bson.D{{"$first", "$name"}}},
				}}},
				bson.D{{"$sort", bson.M{"total_posts": -1}}},
				bson.D{{"$limit", 1}},
			},
		}}},
	}

	cursor, err := usersCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("data: %s \n", result)
}

func query10() {
	// Hiç gönderi yazmamış kullanıcıları getir.

}

/*
30 yaşından büyük olan kullanıcıların gruplarını getir.
Her grubun kaç üyesi olduğunu listele.
En çok gruba üye olan kullanıcıyı getir.
Hiç gruba üye olmamış kullanıcıları getir.
"MongoDB Öğreniyorum" gönderisini yazan kişinin yaşını bul.
Kullanıcıları yazdıkları gönderi sayısına göre sıralı listele.
2’den fazla gruba üye olan kullanıcıları getir.
"Ankara" şehrinde yaşayan kullanıcıların gönderilerini getir.
Kullanıcıların yaş ortalamasını hesapla.
Kullanıcıların gönderi başına ortalama kelime sayısını hesapla.
En son hangi kullanıcı gruba üye oldu?
*/
