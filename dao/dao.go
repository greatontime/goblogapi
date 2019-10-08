package dao

import (
	"context"
	"fmt"
	"log"

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

//getAllBlog from db

func getAllBlog() []primitive.M{
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