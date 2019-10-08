package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/greatontime/goblogapi/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//connectionString eg : "mongodb://localhost:27017"
const connectionString = "mongodb://localhost:27017"

//database name
const dbName ="gobloghellogopher"

//collection name
const collectionname = "blog"

//collection object/instance
var collection *mongo.Collection

// create connection with mongo db

func init(){
	//Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(),clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB ...")

	collection = client.Database(dbName).Collection(collectionname)
	fmt.Println("Collection instance created !!")
}

//GetAllBlog from db
func GetAllBlog() []primitive.M{
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M 
	for cur.Next(context.Background()){
		var result bson.M 
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil{
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}

//InsertOneBlog One blog to mongodb
func InsertOneBlog(blog models.Blog){
	insertResult, err := collection.InsertOne(context.Background(),blog)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record",insertResult.InsertedID)
}
//BlogActive function for active blog
func BlogActive(blog string){
	fmt.Println(blog)
	id, _ := primitive.ObjectIDFromHex(blog)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"status":true}}
	result, err := collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:",result.ModifiedCount)
}
//BlogUnActive function for unactive blog
func BlogUnActive(blog string){
	fmt.Println(blog)
	id, _ := primitive.ObjectIDFromHex(blog)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"status":false}}
	result, err := collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:",result.ModifiedCount)
}
//DeleteOneBlog delete one blog by id
func DeleteOneBlog(blog string){
	fmt.Println(blog)
	id, _ := primitive.ObjectIDFromHex(blog)
	filter := bson.M{"_id":id}
	d, err := collection.DeleteOne(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Document",d.DeletedCount)
}
//DeleteAllBlog delete all blog
func DeleteAllBlog() int64{
	d,err := collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Document Many !!!",d.DeletedCount)
	return d.DeletedCount
}