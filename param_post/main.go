package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main(){
	r:=gin.Default()
	r.POST("/test",func(c *gin.Context){
		bodyByts,err := ioutil.ReadAll(c.Request.Body)
		if err != nil{
			c.String(http.StatusBadRequest,err.Error())
			c.Abort()
		}
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name","default_last_name")
		c.String(http.StatusOK,string(bodyByts),"%s,%s,%s",firstName,lastName,string(bodyByts))
	})
	r.Run(":8080")
}
