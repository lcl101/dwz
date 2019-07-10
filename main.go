package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/lcl101/dwz/conf"
	"github.com/lcl101/dwz/handler"
	"github.com/lcl101/dwz/log"
	"github.com/lcl101/dwz/slot"
	"gopkg.in/yaml.v2"
)

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		log.Log.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}

func run() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(data, &conf.Conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.InitLog()
	slot.InitDB()
	slot.InitSlotGenerator()
	handler.InitCache()
	log.Log.Debug("开始运行")

	// router := gin.Default()
	router := gin.New()
	// router.Use(Logger(), gin.Recovery())
	router.Use(log.Logger(), gin.Recovery())

	// 转换短地址
	router.GET("/t/:path", handler.Turl)

	// 产生短地址
	router.POST("/t/gen", handler.Gen)

	//画波形图
	router.POST("/t/draw", handler.Draw)

	//运行服务
	router.Run(fmt.Sprintf("%s:%d", conf.Conf.Http.IP, conf.Conf.Http.Port))
}

func main() {
	// test1()
	run()
}
