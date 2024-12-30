package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Reservation represents a hotel reservation
type Reservation struct {
    ID                   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    FirstName            string        `json:"first_name"`
    LastName             string        `json:"last_name"`
    Date                 string        `json:"date"`
    ReservationID        int           `json:"reservation_id"`
    RoomArrivalDate      string        `json:"room_arrival_date"`
    RoomCurrencyCode     string        `json:"room_currencycode"`
    RoomDepartureDate    string        `json:"room_departure_date"`
    RoomID               primitive.ObjectID           `json:"room_id"`
    RoomMealPlan         string        `json:"room_meal_plan"`
    RoomNumberOfGuests   int           `json:"room_numberofguests"`
    Guests               []Guest       `json:"guests"`
    RoomPrices           []RoomPrice   `json:"room_prices"`
    RoomReservationID    int           `json:"room_roomreservation_id"`
    RoomTotalPrice       float64       `json:"room_totalprice"`
    Time                 string        `json:"time"`
}

// RoomPrice represents the price    // Here you can use the 'reservation' object in your application details for a room on a specific date
type RoomPrice struct {
    Date    string  `json:"date"`
    RateID  string  `json:"rate_id"`
    Value   float64 `json:"value"`
}
type Guest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Passport  string `json:"passport"`
	BirthDate string `json:"birth_date"`
}

/* 
func reservationFakeData() Reservation {
    // Example initialization
    reservation := Reservation{
        FirstName:          "Davids",
        LastName:           "Snyders",
        Date:               "2024-12-25",
        ReservationID:      325070178,
        RoomArrivalDate:    "2025-01-15",
        RoomCurrencyCode:   "USD",
        RoomDepartureDate:  "2025-01-20",
        RoomID:              primitive.NewObjectID(), 
        RoomMealPlan:       "Breakfast included",
        RoomNumberOfGuests: 2,
        RoomPrices: []RoomPrice{
            {Date: "2025-01-15", RateID: "1354450", Value: 150.0},
            {Date: "2025-01-16", RateID: "1354450", Value: 150.0},
        },
        RoomReservationID:  1799964999,
        RoomTotalPrice:     750.0,
        Time:               "14:30:00",
    }

	return reservation
}
*/