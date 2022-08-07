package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/stackpath/backend-developer-tests/rest-service/controllers"
)

func peopleRouters(superRouter *gin.RouterGroup) {
	peopleRouter := superRouter.Group("/people")
	{
		peopleRouter.GET("", controllers.GetPeople)
		peopleRouter.GET("/:id", controllers.GetPeopleByID)

	}
}
