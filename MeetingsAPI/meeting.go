// Code to insert meeting collection with document
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
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Participants [3]int    `json:"participants"`
	StartTime    time.Time `json:"starttime"`
	EndTime      time.Time `json:"endtime"`
	TimeStamp    time.Time `json:"Timestamp"`
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("MeetingsAPI").Collection("Meeting_Details")
	oneDoc := MongoFields{
		ID:           1,
		Title:        "General Meeting",
		Participants: [3]int{1000, 1001, 1002},
		StartTime:    time.Date(2020, 10, 18, 20, 34, 58, 65777888, time.UTC),
		EndTime:      time.Date(2020, 10, 18, 20, 34, 58, 65777888, time.UTC),
		TimeStamp:    time.Now(),
	}

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr !=
		nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		os.Exit(1)
	} else {
		newID := result.InsertedID
		fmt.Println("InsertOne() newID:", newID)
		fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
	}
}
