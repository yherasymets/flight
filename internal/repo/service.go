package repo

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/yherasymets/flight/models"
)

type repo struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &repo{db}
}

type Service interface {
	Create(flight *models.Flight) error
	Get(id uuid.UUID) (*models.Flight, error)
	Update(id uuid.UUID, flight *models.Flight) error
	Delete(id uuid.UUID) error
}

func (r *repo) Create(flight *models.Flight) error {
	query := `INSERT INTO flight (from_airport, to_airport, airline_name, fligt_number, day_of_week, departure_time, 
				     arrival_time, aircraft_type, flight_time, created_at, updated_at)
		      VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
			  RETURNING id`
	flight.CreatedAt = time.Now()
	flight.UpdatedAt = time.Now()
	return r.db.QueryRow(query,
		flight.FromAirport,
		flight.ToAirport,
		flight.AirlineName,
		flight.FligtNumber,
		pq.Array(flight.DayOfWeek),
		parseTime(flight.DepartureTime),
		parseTime(flight.ArrivalTime),
		pq.Array(flight.AircraftType),
		flight.FlightTime,
		flight.CreatedAt,
		flight.UpdatedAt).Scan(&flight.ID)
}

func (r *repo) Get(id uuid.UUID) (*models.Flight, error) {
	flight := new(models.Flight)
	var dayOfWeek []sql.NullInt64
	query := `SELECT id, from_airport, to_airport, airline_name, fligt_number, day_of_week, departure_time, 
				     arrival_time, aircraft_type, flight_time, created_at, updated_at
			  FROM flight 
			  WHERE id = $1`
	if err := r.db.QueryRow(query, id).
		Scan(&flight.ID,
			&flight.FromAirport,
			&flight.ToAirport,
			&flight.AirlineName,
			&flight.FligtNumber,
			pq.Array(&dayOfWeek),
			&flight.DepartureTime,
			&flight.ArrivalTime,
			pq.Array(&flight.AircraftType),
			&flight.FlightTime,
			&flight.CreatedAt,
			&flight.UpdatedAt,
		); err != nil {
		return nil, err
	}
	for _, v := range dayOfWeek {
		flight.DayOfWeek = append(flight.DayOfWeek, int(v.Int64))
	}
	return flight, nil
}

func (r *repo) Update(id uuid.UUID, flight *models.Flight) error {
	query := `UPDATE flight
			  SET from_airport = $1, to_airport = $2, airline_name = $3, fligt_number = $4, day_of_week = $5, departure_time = $6, 
				  arrival_time = $7, aircraft_type = $8, flight_time = $9, updated_at = $10 
			  WHERE id = $11
			  RETURNING id`
	return r.db.QueryRow(query, flight.FromAirport,
		flight.ToAirport,
		flight.AirlineName,
		flight.FligtNumber,
		pq.Array(flight.DayOfWeek),
		parseTime(flight.DepartureTime),
		parseTime(flight.ArrivalTime),
		pq.Array(flight.AircraftType),
		flight.FlightTime,
		flight.UpdatedAt,
		id).Err()
}

func (r *repo) Delete(id uuid.UUID) error {
	query := `DELETE FROM flight WHERE id = $1`
	return r.db.QueryRow(query, id).Err()
}

func parseTime(customTime models.CustomTime) time.Time {
	time, err := time.Parse(time.TimeOnly, customTime.String())
	if err != nil {
		panic(err)
	}
	return time
}
