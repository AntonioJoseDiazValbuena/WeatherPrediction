package components

import (
	"main/models"

	"gorm.io/gorm"
)

func PredictionsExist(db *gorm.DB) bool {
	var allPredictions []models.Predictions

	allPredictionsResult := db.Find(&allPredictions)

	return allPredictionsResult.RowsAffected == 0
}
