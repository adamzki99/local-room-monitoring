-- DROP SCHEMA smisk;

CREATE SCHEMA smisk AUTHORIZATION postgres;


-- Permissions

GRANT ALL ON SCHEMA smisk TO postgres;

-- Drop table

-- DROP TABLE smisk.data_recordings;

CREATE TABLE smisk.data_recordings (
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
	CONSTRAINT data_recordings_devices_fk FOREIGN KEY (device_id) REFERENCES smisk.devices(device_id)
);

-- Permissions

ALTER TABLE smisk.data_recordings OWNER TO postgres;
GRANT ALL ON TABLE smisk.data_recordings TO postgres;

-- Drop table

-- DROP TABLE smisk.devices;

CREATE TABLE smisk.devices (
	device_id varchar(256) NOT NULL,
	room_id varchar(256) NOT NULL,
	firmware_version varchar(256) NOT NULL,
	address varchar(2048) NOT NULL,
	CONSTRAINT devices_pk PRIMARY KEY (device_id),
	CONSTRAINT devices_locations_fk FOREIGN KEY (room_id) REFERENCES smisk.locations(room_id)
);

-- Permissions

ALTER TABLE smisk.devices OWNER TO postgres;
GRANT ALL ON TABLE smisk.devices TO postgres;

-- Drop table

-- DROP TABLE smisk.locations;

CREATE TABLE smisk.locations (
	room_id varchar(256) NOT NULL,
	latitude float4 NOT NULL,
	longitude float4 NOT NULL,
	altitude float4 NOT NULL,
	CONSTRAINT locations_pk PRIMARY KEY (room_id)
);

-- Permissions

ALTER TABLE smisk.locations OWNER TO postgres;
GRANT ALL ON TABLE smisk.locations TO postgres;