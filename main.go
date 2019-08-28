package main

import (
	"github.com/gin-gonic/gin"
	"github.com/HexChrx/testgokafka/routers"
)

func main()  {
	r := gin.Default()
	routers.Init(r)
	r.Run()
}


