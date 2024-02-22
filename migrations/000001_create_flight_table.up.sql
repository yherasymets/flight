CREATE TABLE IF NOT EXISTS flight (
    id serial PRIMARY KEY,
    from_airport varchar(3) NOT NULL,
    to_airport varchar(3) NOT NULL,
    airline_name varchar(35) NOT NULL,
    fligt_number varchar(10) NOT NULL,
    day_of_week integer[] NOT NULL, 
    departure_time timestamp,
    arrival_time timestamp,
    aircraft_type text[] NOT NULL,
    flight_time text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);