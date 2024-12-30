package controllers

import (
	"context"
	"fmt"

	"github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getRoomID(name string) (primitive.ObjectID, error) {
	// MongoDB filter for available rooms with matched name
	filter := bson.M{
		"status": "Available",
	}
	if name != "" {
		filter["name"] = name
	}
	
	// Find all available rooms with the matched name
	cursor, err := roomCollection.Find(context.Background(), filter)
	if err != nil {
		return primitive.NewObjectID(), fmt.Errorf("failed to fetch rooms: %v", err)
	}
	defer cursor.Close(context.Background())

	fmt.Println(cursor)
	// Store room IDs in a slice
	var roomIDs []primitive.ObjectID
	for cursor.Next(context.Background()) {
		var room models.Room
		if err := cursor.Decode(&room); err != nil {
			return primitive.NewObjectID(), fmt.Errorf("failed to decode room: %v", err)
		}
		roomIDs = append(roomIDs, room.ID)
	}

	fmt.Println(roomIDs)
	// If no rooms found, return an error
	if len(roomIDs) == 0 {
		return primitive.NewObjectID(), fmt.Errorf("no available rooms found")
	}

	// Randomly select a room ID
	return roomIDs[0], nil
}

func updateRoom(ctx context.Context,roomId primitive.ObjectID, guestId string)  {
    result, err := roomCollection.UpdateOne(
        ctx,
        bson.M{"_id": roomId},
        bson.D{
            {Key: "$set", Value: bson.D{{Key: "status", Value: "Occupied"}}},
			{Key: "$set", Value: bson.D{{Key: "guest_id", Value: guestId}}},
        },
    )

    if err != nil {
		fmt.Println("MongoDb Update Error", err)
    }

    fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}