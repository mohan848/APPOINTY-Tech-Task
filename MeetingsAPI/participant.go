// Code to insert participant collection with document
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
	ParticipantID int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	RSVP          string `json:"rsvp"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("MeetingsAPI").Collection("Participant_Details")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	oneDoc := MongoFields{
		ParticipantID: 1002,
		Name:          "dishanth",
		Email:         "dishanth123@gmail.com",
		RSVP:          "YES",
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
