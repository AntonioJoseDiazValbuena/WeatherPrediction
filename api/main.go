package api

import (
	"main/components"
	"main/models"
	"main/persistency"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	r := gin.Default()

	var solarSystem models.SolarSystem

	dsn := "root:getsugatenshou@tcp(127.0.0.1:3306)/weather_prediction?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	r.GET("/configurate", func(c *gin.Context) {
		solarSystem = components.Calculateweather()
		if components.PredictionsExist(db) {
			persistency.SaveIntoDB(&solarSystem, db)
		}
		c.JSON(200, solarSystem.WeatherStatusCount)
	})

	r.GET("/weather", func(c *gin.Context) {
		if components.PredictionsExist(db) {
			c.JSON(404, gin.H{
				"message": "you have to call configurate first",
			})
			return
		}

		dayStr := c.Query("day")
		day, err := strconv.ParseInt(dayStr, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "the day param must be an int",
			})
			return
		}
		if day < 1 || day > 3650 {
			c.JSON(400, gin.H{
				"message": "the day param must be between 1 and 3650",
			})
			return
		}

		var predictions models.Predictions

		db.First(&predictions, day)

		c.JSON(200, gin.H{
			"weather": predictions.WeatherStatus,
		})

	})
	r.Run("localhost:8080") // listen and serve on localhost:8080
}
