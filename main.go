package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Url struct {
	AppKey    string `json:"AppKey"`
	AppSecret string `json:"AppSecret"`
	Lurl      string `json:"Lurl"`
}

type Rtn struct {
	RtnCode string
	RtnMsg  string
	Data    string
}

func main() {
	router := gin.Default()

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	router.GET("/t/:path", func(c *gin.Context) {
		path := c.Param("path")
		fmt.Println(path)
		c.Redirect(http.StatusSeeOther, "http://www.baidu.com")
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.POST("/t/gen", func(c *gin.Context) {
		var url Url
		err := c.BindJSON(&url)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, gin.H{
			"rtnCode": "000000",
			"rtnMsg":  "成功",
			"url":     "/t/hhhhh",
		})
	})

	router.Run(":10000")
}
