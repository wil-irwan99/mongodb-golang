package main

import (
	"golang-mongodb/delivery"
)

// const mongodbUri = "mongodb://localhost:27017"

// type Student struct {
// 	Id       primitive.ObjectID `bson:"_id"`
// 	Name     string             `bson:"fullName"`
// 	Age      int                `bson:"age"`
// 	Gender   string             `bson:"gender"`
// 	JoinDate primitive.DateTime `bson:"joinDate"`
// 	Senior   bool               `bson:"senior"`
// }

func main() {
	delivery.NewServer().Run()
}

// uname := os.Getenv("DB_UNAME")
// pass := os.Getenv("DB_PASS")

// credential := options.Credential{
// 	AuthMechanism: "SCRAM-SHA-256",
// 	Username:      uname,
// 	Password:      pass,
// }

// clientOptions := options.Client()
// clientOptions.ApplyURI(mongodbUri).SetAuth(credential)

// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()

// connect, err := mongo.Connect(context.Background(), clientOptions)
// if err != nil {
// 	panic(err)
// } else {
// 	fmt.Println("Connected...")
// }

// defer func() {
// 	if err := connect.Disconnect(ctx); err != nil {
// 		panic(err)
// 	}
// }()

// //membuat sebuah db - collection
// db := connect.Database("enigma")
// //coll := db.Collection("student")
// coll := db.Collection("products")

//create
//insertMany
// jd01 := parseTime("2022-07-02 15:05:31")
// jd02 := parseTime("2022-07-03 15:05:31")
// jd03 := parseTime("2022-07-04 15:05:31")
// students := []interface{}{
// 	bson.D{
// 		{"name", "Jamal"},
// 		{"age", 21},
// 		{"gender", "M"},
// 		{"joinDate", primitive.NewDateTimeFromTime(jd01)},
// 		{"senior", false},
// 	},
// 	bson.D{
// 		{"name", "Tika"},
// 		{"age", 20},
// 		{"gender", "F"},
// 		{"joinDate", primitive.NewDateTimeFromTime(jd02)},
// 		{"senior", false},
// 	},

// 	bson.D{
// 		{"name", "Ayu"},
// 		{"age", 23},
// 		{"gender", "F"},
// 		{"joinDate", primitive.NewDateTimeFromTime(jd03)},
// 		//{"joinDate", primitive.NewDateTimeFromTime(parseTime("2022-07-03"))},
// 		{"senior", false},
// 	},
// }

// newId, err := coll.InsertMany(ctx, students)
// fmt.Printf("inserted document with ID %v\n", newId.InsertedIDs)

//insertOne
// newId, err := coll.InsertOne(ctx, bson.D{ //bson.D urutan data ngaruh ke mongodb
// 	{"name", "Jack"},
// 	{"age", 22},
// 	{"gender", "M"},
// 	{"senior", false},
// })

// newStudent := Student{
// 	Id:       primitive.NewObjectID(),
// 	Name:     "Doni",
// 	Age:      23,
// 	Gender:   "M",
// 	JoinDate: primitive.NewDateTimeFromTime(parseTime("2022-07-12 15:05:31")),
// 	Senior:   false,
// }

// newId, err := coll.InsertOne(ctx, newStudent)

// if err != nil {
// 	log.Println(err.Error())
// }

// fmt.Printf("inserted document with ID %v\n", newId.InsertedID) //InsertedID untuk insertOne, InsertedIDs untuk insertMany

//delete one
// filter := bson.D{{"fullName", bson.D{{"$eq", "Doni"}}}}
// //opts := options.Delete().SetHint(bson.D{{"_id", 1}})
// deleteDoc, err := coll.DeleteOne(ctx, filter)
// if err != nil {
// 	log.Println(err.Error())
// }

// fmt.Printf("Document Deleted: %v\n", deleteDoc.DeletedCount)

//update one
// filter := bson.D{{"fullName", bson.D{{"$eq", "Doni"}}}}
// update := bson.D{{"$set", bson.D{{"age", 100}}}}
// result, err := coll.UpdateOne(ctx, filter, update)
// if err != nil {
// 	panic(err)
// }
// fmt.Println(result.ModifiedCount)

//select * from student
// //cursor, err := coll.Find(ctx, bson.D{})
// //select name from student
// cursor, err := coll.Find(ctx, bson.D{}, options.Find().SetProjection(bson.D{{"name", 1}, {"_id", 0}}))
// if err != nil {
// 	log.Println(err.Error())
// }
// var students []bson.D
// err = cursor.All(ctx, &students)
// if err != nil {
// 	log.Println(err.Error())
// }

// for _, student := range students {
// 	fmt.Println(student)
// }

//logical, find name, age froms student where gender = "F" && age >= 21
// filterGenderAndAge := bson.D{
// 	{"$and", bson.A{
// 		bson.D{
// 			{"gender", "M"},
// 			{"age", bson.D{{"$gte", 20}}},
// 		},
// 	}},
// }

// projection := bson.D{
// 	{"_id", 0},
// 	{"fullName", 1},
// 	{"age", 1},
// }

// cursor, err := coll.Find(ctx, filterGenderAndAge, options.Find().SetProjection(projection))
// if err != nil {
// 	log.Println(err.Error())
// }
// var students []bson.D
// err = cursor.All(ctx, &students)
// if err != nil {
// 	log.Println(err.Error())
// }

// for _, student := range students {
// 	fmt.Println(student)
// }

//mapping result query ke struct
// filterGenderAndAgeResult := make([]*Student, 0)
// cursor, err := coll.Find(ctx, filterGenderAndAge, options.Find().SetProjection(projection))
// if err != nil {
// 	log.Println(err.Error())
// }

// for cursor.Next(ctx) {
// 	var student Student
// 	err := cursor.Decode(&student)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	filterGenderAndAgeResult = append(filterGenderAndAgeResult, &student)
// }

// for _, student := range filterGenderAndAgeResult {
// 	fmt.Println(student)
// }

//Agregation
// 	count, err := coll.CountDocuments(ctx, bson.D{})
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	fmt.Println("Product Total: ", count)

// 	//with filter
// 	count, err = coll.CountDocuments(ctx, bson.D{{"category", "food"}})
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	fmt.Println("Product Total with Category[food]: ", count)

// 	//match, group, sort, dll
// 	matchStage := bson.D{
// 		{"$match", bson.D{
// 			{"category", "food"},
// 		}},
// 	}

// 	groupStage := bson.D{
// 		{"$group", bson.D{
// 			{"_id", "$category"},
// 			{"Total", bson.D{{"$sum", 1}}},
// 		}},
// 	}

// 	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	var productCount []bson.M
// 	err = cursor.All(ctx, &productCount)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	for _, product := range productCount {
// 		fmt.Printf("Group[%v], Total[%v]\n", product["_id"], product["Total"])
// 	}

// }

// func parseTime(date string) time.Time {
// 	layoutFormat := "2006-01-02 15:04:05"
// 	parse, _ := time.Parse(layoutFormat, date)
// 	return parse
//}

/*
buat koneksi ke mongodb (url) -> mongodb://localhost:27017/
siapkan user Auth: username & password
*/
