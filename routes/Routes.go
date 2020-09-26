package routes

import (
	"github.com/MattRighetti/fdm-repository-backend/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("models", controllers.GetUsers)
		grp1.GET("models/:id", controllers.GetUserByID)
	}
	return r
}
