package persistency

import (
	"main/models"

	"gorm.io/gorm"
)

func SaveIntoDB(solarSystem *models.SolarSystem, db *gorm.DB) {
	predictions := make([]models.Predictions, 3651)

	for index, value := range solarSystem.Weather {
		predictions[index] = models.Predictions{WeatherDay: index, WeatherStatus: value.WeatherStatus}
	}

	db.Create(predictions[1:])
}
