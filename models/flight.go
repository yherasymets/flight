package models

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	ID            uuid.UUID  `json:"id"`
	FromAirport   string     `json:"from"`
	ToAirport     string     `json:"to"`
	AirlineName   string     `json:"airlineName"`
	FligtNumber   string     `json:"fligtNumber"`
	DayOfWeek     []int      `json:"dayOfWeek,omitempty"`
	DepartureTime CustomTime `json:"departureTime"`
	ArrivalTime   CustomTime `json:"arrivalTime"`
	AircraftType  []string   `json:"aircraftType"`
	FlightTime    FlightTime `json:"flightTime,omitempty"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
}
