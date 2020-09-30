package routes

import (
	"github.com/MattRighetti/fdm-repository-backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	group := router.Group("/api")
	{
		group.GET("models", controllers.GetModels)
		group.GET("models/:id", controllers.GetUserByID)
		group.GET("missing-models", controllers.GetMissingModels)
		group.GET("filters", controllers.GetFilters)
	}

	return router
}
