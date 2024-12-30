package helper

import (
	"fmt"
	"time"

	"github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/models"
)

func CalculatePrices(reservation *models.Reservation) error {
	arrivalDate, err := time.Parse("2006-01-02", reservation.RoomArrivalDate)
	if err != nil {
		return fmt.Errorf("invalid arrival date format: %v", err)
	}

	departureDate, err := time.Parse("2006-01-02", reservation.RoomDepartureDate)
	if err != nil {
		return fmt.Errorf("invalid departure date format: %v", err)
	}

	if !departureDate.After(arrivalDate) {
		return fmt.Errorf("departure date must be after arrival date")
	}

	duration := int(departureDate.Sub(arrivalDate).Hours() / 24) // Number of days
	dailyPrice := reservation.RoomTotalPrice / float64(duration)    // Calculate daily price
	var prices []models.RoomPrice

	for i := 0; i < duration; i++ {
		currentDate := arrivalDate.AddDate(0, 0, i)
		prices = append(prices, models.RoomPrice{
			Date:   currentDate.Format("2006-01-02"),
			RateID: fmt.Sprintf("rate-%d", i+1),
			Value:  dailyPrice,
		})
	}

	reservation.RoomPrices = prices
	return nil
}