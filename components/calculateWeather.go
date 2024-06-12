package components

import (
	"main/models"
)

func Calculateweather() models.SolarSystem {
	solarSystem := models.SolarSystem{}.Init()

	for i := 1; i <= 3650; i++ {
		solarSystem.CalculateDayPosition(i)
	}

	solarSystem.CalculatePeakDays()

	return solarSystem
}
