package routers

import "github.com/gin-gonic/gin"

func AddRouters(superRouter *gin.RouterGroup) {
	peopleRouters(superRouter)
}
