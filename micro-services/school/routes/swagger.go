package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "vidyalaya.com/micro-services/school/docs" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
)

func (r routes) addSwagerRoutes(rg *gin.RouterGroup) {
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
