package models

type WalkInRegistration struct {
	FirstName     string  `json:"first_name" binding:"required"`
	Guests        []Guest `json:"guests" binding:"required"`
	LastName      string  `json:"last_name" binding:"required"`
	ArrivalDate   string  `json:"arrival_date" binding:"required"`
	DepartureDate string  `json:"departure_date" binding:"required"`
	RoomType      string  `json:"room_type" binding:"required"`
	MealPlan      string  `json:"meal_plan"`
	NumberOfGuests int     `json:"number_of_guests" binding:"required"`
	TotalPrice    float64 `json:"total_price"`
}