package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"math/rand/v2"

	"time"
)

type properties struct {
	Timestamp        string   `json:"timestamp"`
	DeviceID         string   `json:"device_id"`
	Location         location `json:"location"`
	Metrics          metrics  `json:"metrics"`
	Battery_level    int      `json:"battery_level"`
	Signal_strength  int      `json:"signal_strength"`
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

var a51Location = location{
	Latitude:  37.233522,
	Longitude: -115.812720,
	Altitude:  1361.0,
	RoomID:    "Area51",
}

// getProperties responds as JSON.
func getProperties(c *gin.Context) {

	var respProperties = properties{
		Timestamp:        time.Now().Format(time.RFC3339),
		DeviceID:         "mock",
		Location:         a51Location,
		Metrics:          generateMockMetrics(),
		Battery_level:    100,
		Signal_strength:  100,
		Firmware_version: "mock",
	}

	c.IndentedJSON(http.StatusOK, respProperties)
}

func generateMockMetrics() metrics {

	var mock = metrics{
		Temperature:       float32(rand.IntN(30-10)) + 10.0 + rand.Float32(),
		Humidity:          rand.Float32(),
		Air_quality_index: float32(rand.IntN(500)),
		Co2_levels:        float32(rand.IntN(40000-400)) + 400.0,
		Light_intensity:   rand.Float32(),
		Occupancy:         false,
	}

	return mock
}

func main() {
	router := gin.Default()
	router.GET("/properties", getProperties)

	router.Run("localhost:8080")
}
