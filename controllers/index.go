package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func Index(ctx *gin.Context)  {
	name := ctx.DefaultQuery("name", "Gin")
	fmt.Println("业务执行")
	ctx.String(http.StatusOK, "Hello %s", name)
}
