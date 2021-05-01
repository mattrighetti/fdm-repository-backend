package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mattrighetti/fdm-repository-backend/controllers"
)

//SetupRouter Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	group := router.Group("/api")
	{
		group.GET("models", controllers.GetModels)
		group.GET("models/:id", controllers.GetModelByID)
		group.GET("missing-models", controllers.GetMissingModels)
		group.GET("filters", controllers.GetFilters)
	}

	return router
}
