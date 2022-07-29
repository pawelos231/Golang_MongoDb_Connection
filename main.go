package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoFields struct {
	FieldStr  string `json: "Field Str"`
	FieldInt  int    `json: "Field Int"`
	FieldBool bool   `json: Field Bool`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE: ", reflect.TypeOf(clientOptions), '\n')
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("First_Database").Collection("First Collection")
	fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")
	oneDoc := MongoFields{
		FieldStr:  "This is our first data and its very important",
		FieldInt:  24214212,
		FieldBool: false,
	}
	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("IntertOne Error", insertErr)
	} else {
		fmt.Println("insertOne() result type", reflect.TypeOf(result))
		fmt.Println("insertOne() api result type", result)
		newID := result.InsertedID
		fmt.Println("InsertedOne(), newID ", newID)
		fmt.Println("insertedOne(), newID type: ", reflect.TypeOf(newID))
	}
}
