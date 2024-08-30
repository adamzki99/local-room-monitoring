package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	dbwrapper "github.com/adamzki99/local-room-monitoring/base/data-collector/src/packages"
	"github.com/go-zoox/fetch"
	"github.com/spf13/viper"
)

type properties struct {
	Timestamp        string   `json:"timestamp"`
	DeviceID         string   `json:"device_id"`
	Location         location `json:"location"`
	Metrics          metrics  `json:"metrics"`
	Battery_level    int32    `json:"battery_level"`
	Signal_strength  int32    `json:"signal_strength"`
	Firmware_version string   `json:"firmware_version"`
}

type location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Altitude  float32 `json:"altitude"`
	RoomID    string  `json:"room_id"`
}

type metrics struct {
	Temperature       float32 `json:"temperature"`
	Humidity          float32 `json:"humidity"`
	Air_quality_index float32 `json:"air_quality_index"`
	Co2_levels        float32 `json:"co2_levels"`
	Light_intensity   float32 `json:"light_intensity"`
	Occupancy         bool    `json:"occupancy"`
}

func readDevice(deviceUrl string) properties {

	response, err := fetch.Get(deviceUrl)
	if err != nil {
		panic(err)
	}

	body, err := response.JSON()
	if err != nil {
		fmt.Println(err)
	}

	data := properties{}
	json.Unmarshal([]byte(body), &data)

	return data
}

func ExtractProperties(prop *properties) (dbwrapper.DataRecording, dbwrapper.Device, dbwrapper.Location, error) {

	timeStamp, err := time.Parse(time.RFC3339, prop.Timestamp)

	if err != nil {
		return dbwrapper.DataRecording{}, dbwrapper.Device{}, dbwrapper.Location{}, err
	}

	dataRecording := dbwrapper.DataRecording{
		DeviceID:        prop.DeviceID,
		Timestamp:       timeStamp,
		Temperature:     prop.Metrics.Temperature,
		Humidity:        prop.Metrics.Humidity,
		AirQualityIndex: prop.Metrics.Air_quality_index,
		CO2Levels:       prop.Metrics.Co2_levels,
		LightIntensity:  prop.Metrics.Light_intensity,
		Occupancy:       prop.Metrics.Occupancy,
		SignalStrength:  prop.Signal_strength,
		BatteryLevel:    prop.Battery_level,
	}

	device := dbwrapper.Device{
		DeviceID:        prop.DeviceID,
		RoomID:          prop.Location.RoomID,
		FirmwareVersion: prop.Firmware_version,
	}

	location := dbwrapper.Location{
		RoomID:    prop.Location.RoomID,
		Latitude:  prop.Location.Latitude,
		Longitude: prop.Location.Longitude,
		Altitude:  prop.Location.Altitude,
	}

	return dataRecording, device, location, nil

}

func main() {

	// read env-variables for the database connection
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	databaseConfig := dbwrapper.DatabaseConfig{Host: viper.GetString("DATABASE_HOST"),
		Port:     viper.GetString("DATABASE_PORT"),
		User:     viper.GetString("DATABASE_USER"),
		Password: viper.GetString("DATABASE_PASSWORD"),
		Name:     viper.GetString("DATABASE_NAME"),
		Schema:   viper.GetString("DATABASE_SCHEMA"),
	}

	devices, err := dbwrapper.GetDevicesFromDatabase(databaseConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	for {

		for _, device := range devices {
			time.Sleep(time.Second * 5)

			deviceAddress := fmt.Sprintf("http://%s:8080/properties", device.Address)

			prop := readDevice(deviceAddress)
			dataRec, _, _, err := ExtractProperties(&prop)

			if err != nil {
				fmt.Println(err)
				continue
			}

			_, err = dbwrapper.WriteDataRecordingToDatabase(databaseConfig, &dataRec)

			if err != nil {
				fmt.Println(err)
				continue
			}

		}
	}

}
