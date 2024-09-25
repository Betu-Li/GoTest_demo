package httptest_demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Params 请求参数
type Params struct {
	Name string `json:"name"`
}

// helloHandler /hello请求处理函数，
func helloHandler(c *gin.Context) {
	var p Params
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "we need a name"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("hello %s", p.Name)})
}

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/hello", helloHandler)
	return r
}
