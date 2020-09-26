package Routes

import (
	"github.com/MattRighetti/fdm-repository-backend/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("models", Controllers.GetUsers)
		grp1.GET("models/:id", Controllers.GetUserByID)
	}
	return r
}
