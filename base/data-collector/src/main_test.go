package main

import (
	"testing"
	"time"

	dbwrapper "github.com/adamzki99/local-room-monitoring/base/data-collector/src/packages"
)

func TestExtractProperties(t *testing.T) {

	testMetrics := metrics{Temperature: 0.0,
		Humidity:          1.0,
		Air_quality_index: 2.0,
		Co2_levels:        3.0,
		Light_intensity:   4.0,
		Occupancy:         false,
	}

	testLocation := location{Latitude: 5.0,
		Longitude: 6.0,
		Altitude:  7.0,
		RoomID:    "room0",
	}

	testProperties := properties{
		Timestamp:        "1999-11-27T12:00:00+02:00",
		DeviceID:         "device0",
		Location:         testLocation,
		Metrics:          testMetrics,
		Battery_level:    0,
		Signal_strength:  0,
		Firmware_version: "firmware0",
	}

	result1, result2, result3, err := ExtractProperties(&testProperties)

	expectTime, _ := time.Parse(time.RFC3339, "1999-11-27T12:00:00+02:00")
	expected1 := dbwrapper.DataRecording{
		DeviceID:        "device0",
		Timestamp:       expectTime,
		Temperature:     0.0,
		Humidity:        1.0,
		AirQualityIndex: 2.0,
		CO2Levels:       3.0,
		LightIntensity:  4.0,
		Occupancy:       false,
		SignalStrength:  0,
		BatteryLevel:    0,
	}

	expected2 := dbwrapper.Device{
		DeviceID:        "device0",
		RoomID:          "room0",
		FirmwareVersion: "firmware0",
	}

	expected3 := dbwrapper.Location{
		RoomID:    "room0",
		Latitude:  5.0,
		Longitude: 6.0,
		Altitude:  7.0,
	}

	if err != nil {
		t.Error("TestExtractProperties function test failed. General failure")
	}

	// DataRecording

	if result1.DeviceID != expected1.DeviceID || result2.DeviceID != expected2.DeviceID {
		t.Error("TestExtractProperties function test failed. Failed to extract 'DeviceID'")
	}

	if result1.Timestamp != expected1.Timestamp {
		t.Errorf("TestExtractProperties function test failed. Failed timestamp conversion")
	}

	if result1.Temperature != expected1.Temperature {
		t.Errorf("TestExtractProperties function test failed. Expected %f, but got %f", expected1.Temperature, result1.Temperature)
	}

	if result1.Humidity != expected1.Humidity {
		t.Errorf("TestExtractProperties function test failed. Expected %f, but got %f", expected1.Humidity, result1.Humidity)
	}

	if result1.AirQualityIndex != expected1.AirQualityIndex {
		t.Errorf("TestExtractProperties function test failed. Expected %f, but got %f", expected1.AirQualityIndex, result1.AirQualityIndex)
	}

	if result1.CO2Levels != expected1.CO2Levels {
		t.Errorf("TestExtractProperties function test failed. Expected %f, but got %f", expected1.CO2Levels, result1.CO2Levels)
	}

	if result1.LightIntensity != expected1.LightIntensity {
		t.Errorf("TestExtractProperties function test failed. Expected %f, but got %f", expected1.LightIntensity, result1.LightIntensity)
	}

	if result1.Occupancy != expected1.Occupancy {
		t.Errorf("TestExtractProperties function test failed. Expected %t, but got %t", expected1.Occupancy, result1.Occupancy)
	}

	if result1.SignalStrength != expected1.SignalStrength {
		t.Errorf("TestExtractProperties function test failed. Expected %d, but got %d", expected1.SignalStrength, result1.SignalStrength)
	}

	if result1.BatteryLevel != expected1.BatteryLevel {
		t.Errorf("TestExtractProperties function test failed. Expected %d, but got %d", expected1.BatteryLevel, result1.BatteryLevel)
	}

	// Device

	if result2.RoomID != expected2.RoomID {
		t.Errorf("TestExtractProperties function test failed. Expected %s, but got %s", expected2.RoomID, result2.RoomID)
	}

	if result2.FirmwareVersion != expected2.FirmwareVersion {
		t.Errorf("TestExtractProperties function test failed. Expected %s, but got %s", expected2.FirmwareVersion, result2.FirmwareVersion)
	}

	if result2.Address != expected2.Address {
		t.Errorf("TestExtractProperties function test failed. Expected %s, but got %s", expected2.Address, result2.Address)
	}

	// Location

	if result3.RoomID != expected3.RoomID {
		t.Errorf("TestExtractProperties function test failed. Expected %s, but got %s", expected3.RoomID, result3.RoomID)
	}

	if result3.Latitude != expected3.Latitude {
		t.Errorf("TestExtractProperties function test failed. Expected %f, but got %f", expected3.Latitude, result3.Latitude)
	}

	if result3.Longitude != expected3.Longitude {
		t.Errorf("TestExtractProperties function test failed. Expected %f, but got %f", expected3.Longitude, result3.Longitude)
	}

	if result3.Altitude != expected3.Altitude {
		t.Errorf("TestExtractProperties function test failed. Expected %f, but got %f", expected3.Altitude, result3.Altitude)
	}

}
