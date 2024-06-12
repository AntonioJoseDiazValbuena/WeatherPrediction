package models

import (
	"main/utils"
	"math"
)

type Planet struct {
	Speed       int
	Name        string
	Distance    int
	Coordinates Coordinates
}

func (planet *Planet) ChangePositionByDay(day int) {
	angle := float64(day * planet.Speed)

	planet.Coordinates.X = utils.Round(float64(planet.Distance) * math.Cos(angle))
	planet.Coordinates.Y = utils.Round(float64(planet.Distance) * math.Sin(angle))
}

func (planet Planet) HasSameSlope(secondPlanet, thirdPlanet Planet) bool {
	slope1 := utils.FindSlope(planet.Coordinates.X, secondPlanet.Coordinates.X, planet.Coordinates.Y, secondPlanet.Coordinates.Y)
	slope2 := utils.FindSlope(planet.Coordinates.X, thirdPlanet.Coordinates.X, planet.Coordinates.Y, thirdPlanet.Coordinates.Y)

	return slope1 == slope2
}
