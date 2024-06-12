package models

type WeatherDetail struct {
	WeatherStatus WeatherStatus
	Perimeter     float64
}

type WeatherStatus string

type Coordinates struct {
	X float64
	Y float64
}

const (
	Rainy   WeatherStatus = "Rainy"
	Peak    WeatherStatus = "Peak"
	Normal  WeatherStatus = "Normal"
	Drought WeatherStatus = "Drought"
	Optimum WeatherStatus = "Optimum"
)

type Predictions struct {
	WeatherDay    int           `gorm:"column:weather_day"`
	WeatherStatus WeatherStatus `gorm:"column:weather_status"`
}
