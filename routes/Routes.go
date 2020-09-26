package routes

import (
	"github.com/MattRighetti/fdm-repository-backend/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	group := r.Group("/api")
	{
		group.GET("models", controllers.GetUsers)
		group.GET("models/:id", controllers.GetUserByID)
		group.GET("missingModels", controllers.GetMissingModels)
	}
	return r
}
