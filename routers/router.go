package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/chrx/testgokafka/middlewares"
	"github.com/chrx/testgokafka/controllers"
)

// Init  web路由初始化
func Init(s *gin.Engine) {
	s.Use(middlewares.Print)
	s.GET("/", controllers.Index)
}
