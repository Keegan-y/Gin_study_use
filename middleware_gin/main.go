package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)
//C:\GOPATH\src\github.com\Keegen-y\gin_test_project\middleware_gin
func main(){
	f,_ := os.Create("./middleware_gin/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	r:=gin.New()
	r.Use(gin.Logger(),gin.Recovery())//,gin.Recovery()
	r.GET("/test",func(c *gin.Context){
		name := c.DefaultQuery("name","default_name")
		panic("test panic===模拟服务器挂掉了===")
		c.String(200,"%s",name)
	})
	r.Run()
}
