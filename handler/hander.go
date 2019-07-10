package handler

import (
	"fmt"
	"image/png"

	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lcl101/dwz/conf"
	mydraw "github.com/lcl101/dwz/draw"
	"github.com/lcl101/dwz/log"
	"github.com/lcl101/dwz/slot"
)

// Turl 转换地址
func Turl(c *gin.Context) {
	path := c.Param("path")
	url := slot.FindUrlBySlot(path)
	log.Log.Debugf("%#v\n", url)
	if url == nil {
		c.Redirect(http.StatusSeeOther, conf.Conf.Url.Home)
		return
	}
	if url.Expired() {
		c.Redirect(http.StatusTemporaryRedirect, conf.Conf.Url.Home)
		return
	}
	log.Log.Debug(url.Origin)
	c.Redirect(http.StatusMovedPermanently, url.Origin)
}

// Gen 生成短地址
func Gen(c *gin.Context) {
	var uri slot.Uri
	err := c.BindJSON(&uri)
	if err != nil {
		log.Log.Debug(err)
		printRtn(c, "100000", fmt.Sprintf("参数错误:%s", err.Error()), "")
		return
	}
	url := slot.NewUrl(uri.Url, conf.Conf.Gen.Unique)
	if url == nil {
		printRtn(c, "100002", "创建失败", "")
		return
	}
	printRtn(c, "000000", "创建成", fmt.Sprintf(conf.Conf.Gen.Base, url.Slot))
}

// Draw 音乐波形图
func Draw(c *gin.Context) {
	var musicData mydraw.Music
	err := c.BindJSON(&musicData)
	if err != nil {
		log.Log.Debug(err)
		printRtn(c, "110000", fmt.Sprintf("参数错误:%s", err.Error()), "")
		return
	}
	if musicData.Data == "" || musicData.FileName == "" {
		printRtn(c, "110001", "data or filename is null", "")
		return
	}

	file, err := os.Create(musicData.FileName)
	if err != nil {
		log.Log.Debug(err)
		printRtn(c, "110002", fmt.Sprintf("创建图片文件%s失败", musicData.FileName), "")
		return
	}
	defer file.Close()
	data := str2float(musicData.Data)
	d := mydraw.NewDrawRGBA()
	d.Computed = data
	img := d.Draw()
	// img := mydraw.Draw(6, 1, 5, data1)
	// Encode results as PNG to stdout
	if err := png.Encode(file, img); err != nil {
		log.Log.Debug(err)
		printRtn(c, "110003", "创建png图片文件失败", "")
		return
	}
	printRtn(c, "000000", "图片生成成功", "")
}

func str2float(data string) []float64 {
	tmp := data[1 : len(data)-1]
	arr := strings.Split(tmp, ",")
	f := make([]float64, len(arr))
	for i, s := range arr {
		fl, _ := strconv.ParseFloat(strings.Trim(s, " "), 64)
		f[i] = fl
	}
	return f
}

func printRtn(c *gin.Context, rc, rm, url string) {
	c.JSON(200, gin.H{
		"rtnCode": rc,
		"rtnMsg":  rm,
		"data":    url,
	})
}

func IsUrl(url string) bool {
	//根据dns判断可用
	reg := regexp.MustCompile(`http[s]?:\/\/([\w\.-]+)`)
	match := reg.FindAllStringSubmatch(url, 2)
	if len(match) > 0 && len(match[0]) > 1 {
		domain := match[0][1]
		_, err := net.LookupIP(domain)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
