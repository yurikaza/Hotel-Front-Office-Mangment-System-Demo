package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MealPlan struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Room struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string     `json:"name"`
	GuestId     primitive.ObjectID `json:"guest_id,omitempty"`
	Description string     `json:"description"`
	Capacity    int        `json:"capacity"`
	Amenities   []string   `json:"amenities"`
	MealPlans   []MealPlan `json:"mealPlans"`
	BasePrice   float64    `json:"basePrice"`
	Status      string     `json:"status"`
}