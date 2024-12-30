package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/database"
	helper "github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/helpers"
	"github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var reservationCollection *mongo.Collection = database.OpenCollection(database.Client, "reservation")
var roomCollection *mongo.Collection = database.OpenCollection(database.Client, "rooms")

func CreateReservation() gin.HandlerFunc {
    return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var walkInRegistration models.WalkInRegistration

		//validate the request body
		if err := c.BindJSON(&walkInRegistration); err != nil {
			fmt.Println("Request Data Could not fetch")
			c.JSON(http.StatusBadRequest,  gin.H{"response": "Request Data Could not fetch"})
			return
		}

		roomId, err := getRoomID(walkInRegistration.RoomType)
		if err != nil {
			fmt.Println("Room Id Not Found")
			c.JSON(http.StatusBadRequest,  gin.H{"response": "Room Id Not Found"})
			return
		}

		fmt.Println(roomId)
		reservation := &models.Reservation{
			FirstName: walkInRegistration.FirstName,
			LastName: walkInRegistration.LastName,
			Date: time.Now().UTC().String(),
			ReservationID: int(primitive.NewObjectID()[2]),
			RoomArrivalDate: walkInRegistration.ArrivalDate,
			RoomCurrencyCode: "", // It will be claim after POS device integration
			RoomDepartureDate: walkInRegistration.DepartureDate,
			RoomID: roomId,
			RoomMealPlan: walkInRegistration.MealPlan,
			RoomNumberOfGuests: walkInRegistration.NumberOfGuests,
			Guests: walkInRegistration.Guests,
			RoomReservationID: int(primitive.NewObjectID()[4]),
			RoomTotalPrice: walkInRegistration.TotalPrice,
			Time: time.Time.UTC(time.Now()).String(),
		}

	    err = helper.CalculatePrices(reservation)
	    if err != nil {
			fmt.Println("Price Calc Fail")
			c.JSON(http.StatusBadRequest,  gin.H{"response": "Price Calc Fail"})
		    return
	    }
        
		res2B, _ := json.Marshal(reservation)
		fmt.Println(string(res2B))
		result, err := reservationCollection.InsertOne(ctx, reservation)
		if err != nil {
			fmt.Println("MongoDb Insert Error", err)
			c.JSON(http.StatusBadRequest,  gin.H{"response": "MongoDb Insert Error"})
			return
		}

		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
            updateRoom(ctx, roomId, oid.Hex())
        }

        c.JSON(http.StatusOK, gin.H{"response": result})
    }
}
