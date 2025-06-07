package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (productRouter *ProductRouter) InitProductRouter(router *gin.RouterGroup) {

	//public router
	publicRouter := router.Group("/product")
	{
		publicRouter.GET("/search")
		publicRouter.GET("/detail/:id")
	}
	//private router
}
