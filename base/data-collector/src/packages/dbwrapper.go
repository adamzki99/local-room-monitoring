package dbwrapper

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Schema   string
}

type DB struct {
	conn *sql.DB
}

type DataRecording struct {
	DeviceID        string
	Timestamp       time.Time
	Temperature     float32
	Humidity        float32
	AirQualityIndex float32
	CO2Levels       float32
	LightIntensity  float32
	Occupancy       bool
	SignalStrength  int32
	BatteryLevel    int32
}

type Device struct {
	DeviceID        string
	RoomID          string
	FirmwareVersion string
	Address         string
}

type Location struct {
	RoomID    string
	Latitude  float32
	Longitude float32
	Altitude  float32
}

// newDatabaseConnection initializes the connection to the database and returns a DB struct.
func newDatabaseConnection(config DatabaseConfig) (*DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Check if the connection is alive.
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{conn: db}, nil
}

// Close closes the database connection.
func (db *DB) closeDatabaseConnection() error {
	return db.conn.Close()
}

// QueryRow executes a query that returns a single row.
func (db *DB) queryRow(query string, args ...interface{}) *sql.Row {
	return db.conn.QueryRow(query, args...)
}

// Query executes a query that returns multiple rows.
func (db *DB) query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

// Exec executes a query that doesn't return rows (e.g., INSERT, UPDATE, DELETE).
func (db *DB) exec(query string, args ...interface{}) (sql.Result, error) {
	return db.conn.Exec(query, args...)
}

// GetDataRecordingsFromDatabase retrieves all data recordings from the "data_recordings" table in the database specified by the provided config.
// It returns a slice of DataRecording and an error if any occurs during the process.
// The function establishes a database connection, executes a query to select all records, scans each row into a DataRecording struct, and appends it to the result slice.
// If an error occurs at any point, the database connection is closed and the error is returned.
func GetDataRecordingsFromDatabase(config DatabaseConfig) ([]DataRecording, error) {

	var dataRecordings []DataRecording

	connection, err := newDatabaseConnection(config)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM %s.data_recordings", config.Schema)

	rows, err := connection.query(query)

	if err != nil {
		connection.closeDatabaseConnection()
		return nil, err
	}

	for rows.Next() {

		var rec DataRecording

		if err := rows.Scan(rec.DeviceID,
			rec.Timestamp,
			rec.Temperature,
			rec.Humidity,
			rec.AirQualityIndex,
			rec.CO2Levels,
			rec.LightIntensity,
			rec.Occupancy,
			rec.SignalStrength,
			rec.BatteryLevel); err != nil {

			connection.closeDatabaseConnection()
			return dataRecordings, err
		}
		dataRecordings = append(dataRecordings, rec)
	}

	connection.closeDatabaseConnection()
	return dataRecordings, nil

}

// WriteDataRecordingToDatabase inserts a new DataRecording into the "devices" table in the database specified by the provided config.
// It returns a sql.Result indicating the outcome of the insert operation and an error if any occurs during the process.
// The function establishes a database connection, constructs an SQL INSERT query using the data from the DataRecording struct,
// executes the query, and closes the database connection. A TODO comment notes the need to change to a prepared statement to avoid SQL injection risks.
func WriteDataRecordingToDatabase(config DatabaseConfig, dataRecording *DataRecording) (sql.Result, error) {

	connection, err := newDatabaseConnection(config)

	if err != nil {
		return nil, err
	}

	timestamp := dataRecording.Timestamp

	// TODO: change to a prepared statement to avoid sql-injection
	query := fmt.Sprintf("INSERT INTO %s.data_recordings (device_id, timestamp, temperature, humidity, air_quality_index, co2_levels, light_intensity, occupancy, signal_strength, battery_level) values ('%s', '%s', '%f', '%f', '%f', '%f', '%f', '%t', '%d', '%d')",
		config.Schema,
		dataRecording.DeviceID,
		timestamp.Format(time.RFC3339),
		dataRecording.Temperature,
		dataRecording.Humidity,
		dataRecording.AirQualityIndex,
		dataRecording.CO2Levels,
		dataRecording.LightIntensity,
		dataRecording.Occupancy,
		dataRecording.SignalStrength,
		dataRecording.BatteryLevel,
	)

	result, err := connection.exec(query)

	if err != nil {
		connection.closeDatabaseConnection()
		return nil, err
	}

	connection.closeDatabaseConnection()
	return result, err
}

// GetDevicesFromDatabase retrieves all devices from the "devices" table in the database specified by the provided config.
// It returns a slice of Device structs and an error if any occurs during the process.
// The function establishes a database connection, executes a query to select all records from the "devices" table, scans each row into a Device struct, and appends it to the result slice.
// If an error occurs at any point, the database connection is closed and the error is returned.
func GetDevicesFromDatabase(config DatabaseConfig) ([]Device, error) {

	var devices []Device

	connection, err := newDatabaseConnection(config)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM %s.devices", config.Schema)

	rows, err := connection.query(query)

	if err != nil {
		connection.closeDatabaseConnection()
		return nil, err
	}

	for rows.Next() {

		var dev Device

		if err := rows.Scan(&dev.RoomID, &dev.DeviceID, &dev.FirmwareVersion, &dev.Address); err != nil {
			connection.closeDatabaseConnection()
			return devices, err
		}
		devices = append(devices, dev)
	}

	connection.closeDatabaseConnection()
	return devices, nil

}

// WriteDeviceToDatabase inserts a new Device into the "devices" table in the database specified by the provided config.
// It returns a sql.Result indicating the outcome of the insert operation and an error if any occurs during the process.
// The function establishes a database connection, constructs an SQL INSERT query using the data from the Device struct,
// executes the query, and closes the database connection. A TODO comment highlights the need to switch to a prepared statement to mitigate SQL injection risks.
func WriteDeviceToDatabase(config DatabaseConfig, device *Device) (sql.Result, error) {

	connection, err := newDatabaseConnection(config)

	if err != nil {
		return nil, err
	}

	// TODO: change to a prepared statement to avoid sql-injection
	query := fmt.Sprintf("INSERT INTO %s.devices (device_id, room_id, firmware_version, address) values ('%s', '%s', '%s', '%s')",
		config.Schema,
		device.DeviceID,
		device.RoomID,
		device.FirmwareVersion,
		device.Address,
	)

	result, err := connection.exec(query)

	if err != nil {
		connection.closeDatabaseConnection()
		return nil, err
	}

	connection.closeDatabaseConnection()
	return result, err
}

// GetLocationsFromDatabase retrieves all locations from the "locations" table in the database specified by the provided config.
// It returns a slice of Location structs and an error if any occurs during the process.
// The function establishes a database connection, executes a query to select all records from the "locations" table, scans each row into a Location struct, and appends it to the result slice.
// If an error occurs at any point, the database connection is closed and the error is returned.
func GetLocationsFromDatabase(config DatabaseConfig) ([]Location, error) {

	var locations []Location

	connection, err := newDatabaseConnection(config)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM %s.locations", config.Schema)

	rows, err := connection.query(query)

	if err != nil {
		connection.closeDatabaseConnection()
		return nil, err
	}

	for rows.Next() {

		var loc Location

		if err := rows.Scan(&loc.RoomID, &loc.Latitude, &loc.Longitude, &loc.Altitude); err != nil {
			connection.closeDatabaseConnection()
			return locations, err
		}
		locations = append(locations, loc)
	}

	connection.closeDatabaseConnection()
	return locations, nil

}

// WriteLocationToDatabase inserts a new Location into the "locations" table in the database specified by the provided config.
// It returns a sql.Result indicating the outcome of the insert operation and an error if any occurs during the process.
// The function establishes a database connection, constructs an SQL INSERT query using the data from the Location struct,
// executes the query, and closes the database connection. A TODO comment highlights the need to switch to a prepared statement to prevent SQL injection risks.
func WriteLocationToDatabase(config DatabaseConfig, location *Location) (sql.Result, error) {

	connection, err := newDatabaseConnection(config)

	if err != nil {
		return nil, err
	}

	// TODO: change to a prepared statement to avoid sql-injection
	query := fmt.Sprintf("INSERT INTO %s.locations (room_id, latitude, longitude, altitude) values ('%s', %f, %f, %f)",
		config.Schema,
		location.RoomID,
		location.Latitude,
		location.Longitude,
		location.Altitude,
	)

	result, err := connection.exec(query)

	if err != nil {
		connection.closeDatabaseConnection()
		return nil, err
	}

	connection.closeDatabaseConnection()
	return result, err

}
