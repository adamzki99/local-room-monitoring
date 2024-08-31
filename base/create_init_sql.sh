#!/bin/bash

# Load the environment variables from the .env file
export $(grep -v '^#' .env | xargs)

# Set the file location for init.sql 
fileLocation="$(pwd)/database/init.sql"

# Remove file if old instance already exists
rm $fileLocation

# Create the init.sql file
touch $fileLocation

# Append commands required for schema creation

echo "-- Create database SCHEMA" >> "$fileLocation"
output_string="CREATE SCHEMA $DATABASE_SCHEMA AUTHORIZATION $POSTGRES_USER;
GRANT ALL ON SCHEMA $DATABASE_SCHEMA TO $POSTGRES_USER;"
echo "$output_string" >> "$fileLocation"
echo "" >> "$fileLocation" 
echo "" >> "$fileLocation" 

# Append commands required for locations table

echo "-- Create locations TABLE" >> "$fileLocation"
output_string="CREATE TABLE $DATABASE_SCHEMA.locations (
	room_id varchar(256) NOT NULL,
	latitude float4 NOT NULL,
	longitude float4 NOT NULL,
	altitude float4 NOT NULL,
	CONSTRAINT locations_pk PRIMARY KEY (room_id)
);"
echo "$output_string" >> "$fileLocation"
echo "" >> "$fileLocation"

output_string="ALTER TABLE $DATABASE_SCHEMA.locations OWNER TO $POSTGRES_USER;
GRANT ALL ON TABLE $DATABASE_SCHEMA.locations TO $POSTGRES_USER;"
echo "$output_string" >> "$fileLocation"
echo "" >> "$fileLocation"
echo "" >> "$fileLocation"

# Append commands required for devices table

echo "-- Create devices TABLE" >> "$fileLocation"
output_string="CREATE TABLE $DATABASE_SCHEMA.devices (
	device_id varchar(256) NOT NULL,
	room_id varchar(256) NOT NULL,
	firmware_version varchar(256) NOT NULL,
	address varchar(2048) NOT NULL,
	CONSTRAINT devices_pk PRIMARY KEY (device_id),
	CONSTRAINT devices_locations_fk FOREIGN KEY (room_id) REFERENCES $DATABASE_SCHEMA.locations(room_id)
);"
echo "$output_string" >> "$fileLocation"
echo "" >> "$fileLocation"

output_string="ALTER TABLE $DATABASE_SCHEMA.devices OWNER TO $POSTGRES_USER;
GRANT ALL ON TABLE $DATABASE_SCHEMA.devices TO $POSTGRES_USER;"
echo "$output_string" >> "$fileLocation"
echo "" >> "$fileLocation"
echo "" >> "$fileLocation"

# Append commands required for data_recordings table

echo "-- Create data_recordings TABLE" >> "$fileLocation"
output_string="CREATE TABLE $DATABASE_SCHEMA.data_recordings (
	device_id varchar(256) NOT NULL,
	"timestamp" timestamp NOT NULL,
	temperature float4 NULL,
	humidity float4 NULL,
	air_quality_index float4 NULL,
	co2_levels float4 NULL,
	light_intensity float4 NULL,
	occupancy bool NULL,
	signal_strength int4 NULL,
	battery_level int4 NULL,
	CONSTRAINT data_recordings_devices_fk FOREIGN KEY (device_id) REFERENCES $DATABASE_SCHEMA.devices(device_id)
);"
echo "$output_string" >> "$fileLocation"
echo "" >> "$fileLocation"

output_string="ALTER TABLE $DATABASE_SCHEMA.data_recordings OWNER TO $POSTGRES_USER;
GRANT ALL ON TABLE $DATABASE_SCHEMA.data_recordings TO $POSTGRES_USER;"
echo "$output_string" >> "$fileLocation"
echo "" >> "$fileLocation"
echo "" >> "$fileLocation"

# Append commands required for creation of a user for grafana

echo "-- Create grafana USER" >> "$fileLocation"
output_string="CREATE USER $DATABASE_GRAFANA_USER WITH PASSWORD '$DATABASE_GRAFANA_PASSWORD';
GRANT USAGE ON SCHEMA $DATABASE_SCHEMA TO $DATABASE_GRAFANA_USER;
GRANT SELECT ON $DATABASE_SCHEMA.data_recordings TO $DATABASE_GRAFANA_USER;"
echo "$output_string" >> "$fileLocation"
echo "" >> "$fileLocation"
echo "" >> "$fileLocation"

# Print a message indicating success
echo "Created file init.sql in directory $(pwd)/database/"