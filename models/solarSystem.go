package models

import (
	"main/utils"
	"math"
)

type SolarSystem struct {
	Ferengi            Planet
	Vulcano            Planet
	Betazoide          Planet
	Sun                Planet
	Weather            map[int]WeatherDetail
	MaxPerimeter       float64
	WeatherStatusCount map[WeatherStatus]int
}

func (solarSystem SolarSystem) Init() SolarSystem {
	return SolarSystem{
		Ferengi: Planet{
			Speed:    1,
			Name:     "Ferengi",
			Distance: 500,
			Coordinates: Coordinates{
				X: 500,
			},
		},
		Vulcano: Planet{
			Speed:    5,
			Name:     "Vulcano",
			Distance: 1000,
			Coordinates: Coordinates{
				X: 1000,
			},
		},
		Betazoide: Planet{
			Speed:    3,
			Name:     "Betazoide",
			Distance: 2000,
			Coordinates: Coordinates{
				X: 2000,
			},
		},
		Sun: Planet{
			Name: "Sun",
		},
		Weather: map[int]WeatherDetail{},
		WeatherStatusCount: map[WeatherStatus]int{
			Rainy:   0,
			Peak:    0,
			Normal:  0,
			Drought: 0,
			Optimum: 0,
		},
	}
}

func (solarSystem *SolarSystem) CalculateDayPosition(day int) {
	solarSystem.Ferengi.ChangePositionByDay(day)
	solarSystem.Vulcano.ChangePositionByDay(day)
	solarSystem.Betazoide.ChangePositionByDay(day)

	currentPerimeter := solarSystem.GetPerimeter()

	if solarSystem.Ferengi.HasSameSlope(solarSystem.Vulcano, solarSystem.Betazoide) {
		if solarSystem.Ferengi.HasSameSlope(solarSystem.Sun, solarSystem.Betazoide) {
			solarSystem.Weather[day] = WeatherDetail{WeatherStatus: Drought, Perimeter: currentPerimeter}
			return
		}
		solarSystem.Weather[day] = WeatherDetail{WeatherStatus: Optimum, Perimeter: currentPerimeter}
		return
	}

	if currentPerimeter > solarSystem.MaxPerimeter {
		solarSystem.MaxPerimeter = currentPerimeter
	}

	if solarSystem.SunIsInTriangle() {
		solarSystem.Weather[day] = WeatherDetail{WeatherStatus: Rainy, Perimeter: currentPerimeter}
		return
	}

	solarSystem.Weather[day] = WeatherDetail{WeatherStatus: Normal, Perimeter: currentPerimeter}
}

func (solarSystem SolarSystem) GetPerimeter() float64 {
	// P1 = Ferengi
	// P2 = Vulcano
	// P3 = Betazoide

	perimeter := math.Sqrt(
		// X1 & X2
		math.Pow(solarSystem.Vulcano.Coordinates.X-solarSystem.Ferengi.Coordinates.X, 2)+
			// Y1 & Y2
			math.Pow(solarSystem.Vulcano.Coordinates.Y-solarSystem.Ferengi.Coordinates.Y, 2),
	) +
		math.Sqrt(
			// X2 & X3
			math.Pow(solarSystem.Betazoide.Coordinates.X-solarSystem.Vulcano.Coordinates.X, 2)+
				// Y2 & Y3
				math.Pow(solarSystem.Betazoide.Coordinates.Y-solarSystem.Vulcano.Coordinates.Y, 2),
		) +
		math.Sqrt(
			// X1 & X3
			math.Pow(solarSystem.Betazoide.Coordinates.X-solarSystem.Ferengi.Coordinates.X, 2)+
				// Y1 & Y3
				math.Pow(solarSystem.Betazoide.Coordinates.Y-solarSystem.Ferengi.Coordinates.Y, 2),
		)

	return utils.Round(perimeter)
}

func (solarSystem SolarSystem) SunIsInTriangle() bool {
	a, b, c, p := solarSystem.Ferengi.Coordinates, solarSystem.Vulcano.Coordinates, solarSystem.Betazoide.Coordinates, solarSystem.Sun.Coordinates

	dx, dy := b.X-a.X, b.Y-a.Y

	ex, ey := c.X-a.X, c.Y-a.Y

	w1 := (ex*(a.Y-p.Y) + ey*(p.X-a.X)) / (dx*ey - dy*ex)

	w2 := (p.Y - a.Y - w1*dy) / ey

	return (w1 >= 0) && (w2 >= 0) && ((w1 + w2) <= 1)
}

func (solarSystem *SolarSystem) CalculatePeakDays() {
	maxPerimeter := solarSystem.MaxPerimeter

	for day, weatherDetail := range solarSystem.Weather {
		if weatherDetail.Perimeter != maxPerimeter {
			solarSystem.WeatherStatusCount[weatherDetail.WeatherStatus]++

			continue
		}

		weatherDetail.WeatherStatus = Peak

		solarSystem.WeatherStatusCount[weatherDetail.WeatherStatus]++

		solarSystem.Weather[day] = weatherDetail
	}
}
