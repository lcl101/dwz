package log

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lcl101/dwz/conf"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// var timeFormat = "02/Jan/2006:15:04:05 -0700"
var timeFormat = "20060102 15:04:05"

func InitLog() {
	Log = logrus.New()

	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	Log.Out = src
	level := conf.Conf.LogLevel
	switch level {
	case "debug":
		setNull()
		Log.SetLevel(logrus.DebugLevel)
		// Log.SetOutput(os.Stderr)
	case "info":
		setNull()
		Log.SetLevel(logrus.InfoLevel)
	case "warn":
		setNull()
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		setNull()
		Log.SetLevel(logrus.ErrorLevel)
	default:
		setNull()
		Log.SetLevel(logrus.InfoLevel)
	}

	apiLogPath := "dlog.log"
	logWriter, err := rotatelogs.New(
		apiLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(apiLogPath),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	// lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{TimestampFormat: timeFormat})
	Log.AddHook(lfHook)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}
		entry := Log.WithFields(logrus.Fields{
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			// msg := fmt.Sprintf("%s %s %s %d %d %s %s (%dms)", clientIP, c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			if statusCode > 499 {
				entry.Error()
			} else if statusCode > 399 {
				entry.Warn()
			} else {
				entry.Info()
			}
		}
	}
}

func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	logrus.SetOutput(writer)
}
