#!/bin/bash

# Exit script on any error
set -e

# Check for .env file and load environment variables
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo ".env file not found!"
  exit 1
fi

# Define the file location for init.sql
file_location="$(pwd)/database/init.sql"

# Remove old file if it exists
if [ -f "$file_location" ]; then
  rm "$file_location"
fi

# Create a new init.sql file
touch "$file_location"

# Function to append SQL commands to the file
append_to_file() {
  local content="$1"
  echo "$content" >> "$file_location"
  echo "" >> "$file_location"
}

# Append commands for schema creation
append_to_file "-- Create database SCHEMA
CREATE SCHEMA $DATABASE_SCHEMA AUTHORIZATION $POSTGRES_USER;
GRANT ALL ON SCHEMA $DATABASE_SCHEMA TO $POSTGRES_USER;"

# Append commands for locations table
append_to_file "-- Create locations TABLE
CREATE TABLE $DATABASE_SCHEMA.locations (
  room_id varchar(256) NOT NULL,
  latitude float4 NOT NULL,
  longitude float4 NOT NULL,
  altitude float4 NOT NULL,
  CONSTRAINT locations_pk PRIMARY KEY (room_id)
);
ALTER TABLE $DATABASE_SCHEMA.locations OWNER TO $POSTGRES_USER;
GRANT ALL ON TABLE $DATABASE_SCHEMA.locations TO $POSTGRES_USER;"

# Append commands for devices table
append_to_file "-- Create devices TABLE
CREATE TABLE $DATABASE_SCHEMA.devices (
  device_id varchar(256) NOT NULL,
  room_id varchar(256) NOT NULL,
  firmware_version varchar(256) NOT NULL,
  address varchar(2048) NOT NULL,
  CONSTRAINT devices_pk PRIMARY KEY (device_id),
  CONSTRAINT devices_locations_fk FOREIGN KEY (room_id) REFERENCES $DATABASE_SCHEMA.locations(room_id)
);
ALTER TABLE $DATABASE_SCHEMA.devices OWNER TO $POSTGRES_USER;
GRANT ALL ON TABLE $DATABASE_SCHEMA.devices TO $POSTGRES_USER;"

# Append commands for data_recordings table
append_to_file "-- Create data_recordings TABLE
CREATE TABLE $DATABASE_SCHEMA.data_recordings (
  device_id varchar(256) NOT NULL,
  \"timestamp\" timestamp NOT NULL,
  temperature float4 NULL,
  humidity float4 NULL,
  air_quality_index float4 NULL,
  co2_levels float4 NULL,
  light_intensity float4 NULL,
  occupancy bool NULL,
  signal_strength int4 NULL,
  battery_level int4 NULL,
  CONSTRAINT data_recordings_devices_fk FOREIGN KEY (device_id) REFERENCES $DATABASE_SCHEMA.devices(device_id)
);
ALTER TABLE $DATABASE_SCHEMA.data_recordings OWNER TO $POSTGRES_USER;
GRANT ALL ON TABLE $DATABASE_SCHEMA.data_recordings TO $POSTGRES_USER;"

# Append commands for creating a user for Grafana
append_to_file "-- Create grafana USER
CREATE USER $DATABASE_GRAFANA_USER WITH PASSWORD '$DATABASE_GRAFANA_PASSWORD';
GRANT USAGE ON SCHEMA $DATABASE_SCHEMA TO $DATABASE_GRAFANA_USER;
GRANT SELECT ON $DATABASE_SCHEMA.data_recordings TO $DATABASE_GRAFANA_USER;"

# Print a message indicating success
echo "Created file init.sql in directory $(pwd)/database/"
