package route

import (
	"github.com/gin-gonic/gin"
	"desafioGo/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define a rota GET /api/fruits/report-sugar
	r.GET("/api/fruits/report-sugar", controller.ReportSugar)

	return r
}
