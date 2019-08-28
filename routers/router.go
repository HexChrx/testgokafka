package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/HexChrx/testgokafka/middlewares"
	"github.com/HexChrx/testgokafka/controllers"
)

// Init  web路由初始化
func Init(s *gin.Engine) {
	s.Use(middlewares.Print)
	s.GET("/", controllers.Index)
}
