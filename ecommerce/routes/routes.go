package routes

import (
	"github.com/Forha-D/ecommerce/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes( incomingRouter *gin.Engine){
	incomingRouter.POST("/users/signup",     controller.Signup())
	incomingRouter.POST("/users/login",      controller.Login())
	incomingRouter.POST("/admin/addproduct", controller.ProductViewAdmin())
	incomingRouter.GET("/users/productview", controller.SearchProduct())
	incomingRouter.GET("/users/search",      controller.SearchProductByQuery())

}