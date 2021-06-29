package Routes

import (
	"MS1/Controller"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() {
	r := gin.Default()
	grp1 := r.Group("/v1")
	{
		grp1.GET("user", Controller.GetUsers)
		grp1.POST("user", Controller.CreateUser)
	}

	r.Run(":8000")
}
