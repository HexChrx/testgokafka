package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Print 一个简单的使用gin 中间件的示例
func Print(ctx *gin.Context) {
	fmt.Println("你好，我在业务执行前运行了")
	ctx.Next()
	fmt.Println("你好，我在业务执行后运行了")
}
