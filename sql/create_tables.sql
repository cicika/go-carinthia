CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  email text UNIQUE,
  password text,
  name text NOT NULL,
  auth_token text,
  type text NOT NULL,
  beacon_identifier text
);

CREATE TABLE IF NOT EXISTS beacons (
  id serial PRIMARY KEY,
  identifier text UNIQUE,
  user_id int REFERENCES users,
  transport_type text NOT_NULL,
  provider_id,
  is_stationary boolean
);

CREATE TABLE IF NOT EXISTS providers (
  id serial PRIMARY KEY,
  name text NOT NULL,
  transport_type text NOT NULL
);

CREATE TABLE IF NOT EXISTS trips (
  id serial PRIMARY KEY,
  user_id int REFERENCES users,
  started_at bigint,
  ended_at bigint,
  total_price decimal
);

CREATE TABLE IF NOT EXISTS trip_segments (
  id serial PRIMARY KEY,
  user_id int REFERENCES users,
  beacon_identifier text,
  trip_id int REFERENCES trips,
  check_in_type int,
  transport_station_id int,
  latitude numeric,
  longitude numeric,
  passed_at bigint
);

CREATE TABLE IF NOT EXISTS transport_stations (
  id serial PRIMARY KEY,
  name text,
  providers_identifier int,
  latitude numeric,
  longitude numeric,
  type text,
  zone_id int
);
