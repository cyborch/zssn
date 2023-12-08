package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"cyborch.com/apocalypse/api"
	"cyborch.com/apocalypse/docs"
	"cyborch.com/apocalypse/pkg/db"
)

// @BasePath /api/v1
func main() {
	godotenv.Load()
	DatabaseUri := os.Getenv("DB_URI")
	database := db.Connect(DatabaseUri)

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", func(c *gin.Context) {
				api.Register(c, database)
			})
			user.PUT("/:id/location", func(c *gin.Context) {
				api.UpdateLocation(c, database)
			})
			user.POST("/:id/flag", func(c *gin.Context) {
				api.Flag(c, database)
			})
			user.POST("/trade", func(c *gin.Context) {
				api.Trade(c, database)
			})
		}
		report := v1.Group("/report")
		{
			report.GET("/percentage", func(c *gin.Context) {
				api.ReportFlaggedUserPercentage(c, database)
			})
			report.GET("/averages", func(c *gin.Context) {
				api.ReportAverages(c, database)
			})
			report.GET("/lost", func(c *gin.Context) {
				api.ReportLostValue(c, database)
			})
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
