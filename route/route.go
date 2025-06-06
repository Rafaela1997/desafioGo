package route

import (
	"github.com/gin-gonic/gin"
	"desafioGo/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/fruits/report-sugar", controller.ReportSugar)

	return r
}
