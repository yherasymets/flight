package models

import (
	"time"
)

type Flight struct {
	ID            int        `json:"id"`
	FromAirport   string     `json:"from"`
	ToAirport     string     `json:"to"`
	AirlineName   string     `json:"airlineName"`
	FligtNumber   string     `json:"fligtNumber"`
	DayOfWeek     []int      `json:"dayOfWeek,omitempty"`
	DepartureTime time.Time  `json:"departureTime"`
	ArrivalTime   time.Time  `json:"arrivalTime"`
	AircraftType  []string   `json:"aircraftType"`
	FlightTime    FlightTime `json:"flightTime,omitempty"`
	CreatedAt     time.Time  `json:"createdAt"`
}
