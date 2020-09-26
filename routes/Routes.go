package routes

import (
	"github.com/MattRighetti/fdm-repository-backend/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	group := router.Group("/api")
	{
		group.GET("models", controllers.GetUsers)
		group.GET("models/:id", controllers.GetUserByID)
		group.GET("missingModels", controllers.GetMissingModels)
	}

	return router
}
