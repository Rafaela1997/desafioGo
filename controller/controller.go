package controller

import (
	"desafioGo/database"
	"desafioGo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportSugar(c *gin.Context) {
	var fruits []model.Fruit

	// Busca todas as frutas do banco
	database.DB.Find(&fruits)

	// Classifica frutas
	highSugar := []gin.H{}
	lowSugar := []gin.H{}

	for _, fruit := range fruits {
		data := gin.H{"id": fruit.ID, "name": fruit.Name}
		if fruit.Sugar >= 10 {
			highSugar = append(highSugar, data)
		} else {
			lowSugar = append(lowSugar, data)
		}
	}

	// Resposta JSON
	c.JSON(http.StatusOK, gin.H{
		"high_sugar":       highSugar,
		"low_sugar":        lowSugar,
		"total_high_sugar": len(highSugar),
		"total_low_sugar":  len(lowSugar),
	})
}
