package repo

import (
	"database/sql"
	"time"

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
	Get(id int64) (*models.Flight, error)
	Update(id int64, flight *models.Flight) error
	Delete(id int64) error
}

func (r *repo) Create(flight *models.Flight) error {
	var id int
	query := `INSERT INTO flight (from_airport, to_airport, airline_name, fligt_number, day_of_week, departure_time, 
							      arrival_time, aircraft_type, flight_time, created_at)
		      VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			  RETURNING id`
	if err := r.db.QueryRow(query,
		flight.FromAirport,
		flight.ToAirport,
		flight.AirlineName,
		flight.FligtNumber,
		pq.Array(flight.DayOfWeek),
		flight.DepartureTime,
		flight.ArrivalTime,
		pq.Array(flight.AircraftType),
		flight.FlightTime,
		time.Now()).
		Scan(&id); err != nil {
		return err
	}
	flight.ID = id
	return nil
}

func (r *repo) Get(id int64) (*models.Flight, error) {
	flight := new(models.Flight)
	var dayOfWeek []sql.NullInt64
	query := `SELECT id, from_airport, to_airport, airline_name, fligt_number, day_of_week, departure_time, 
					arrival_time, aircraft_type, flight_time, created_at
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
		); err != nil {
		return nil, err
	}
	for _, v := range dayOfWeek {
		flight.DayOfWeek = append(flight.DayOfWeek, int(v.Int64))
	}
	return flight, nil
}

func (r *repo) Update(id int64, flight *models.Flight) error {
	query := `UPDATE flight
					SET from_airport = $1, to_airport = $2, airline_name = $3, fligt_number = $4, day_of_week = $5, departure_time = $6, 
					arrival_time = $7, aircraft_type = $8, flight_time = $9 
					WHERE id = $10
					RETURNING id`
	_, err := r.db.Exec(query, flight.FromAirport,
		flight.ToAirport,
		flight.AirlineName,
		flight.FligtNumber,
		pq.Array(flight.DayOfWeek),
		flight.DepartureTime,
		flight.ArrivalTime,
		pq.Array(flight.AircraftType),
		flight.FlightTime,
		id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) Delete(id int64) error {
	query := `DELETE FROM flight WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
