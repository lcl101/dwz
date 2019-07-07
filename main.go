package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lcl101/dwz/conf"
	mydraw "github.com/lcl101/dwz/draw"
	"github.com/lcl101/dwz/slot"
	"gopkg.in/yaml.v2"
)

// type Url struct {
// 	AppKey    string `json:"AppKey"`
// 	AppSecret string `json:"AppSecret"`
// 	Lurl      string `json:"Lurl"`
// }

// type Rtn struct {
// 	RtnCode string
// 	RtnMsg  string
// 	Data    string
// }

func test1() {
	file, err := os.Create("test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data1 := []float64{0.001, 0.011, 0.331, 0.364, 0.301, 0.272, 0.327, 0.326, 0.335, 0.403, 0.365, 0.344, 0.342, 0.385, 0.394, 0.319, 0.285, 0.378, 0.377, 0.366, 0.335, 0.326, 0.345, 0.339, 0.299, 0.363, 0.341, 0.344, 0.337, 0.333, 0.368, 0.340, 0.296, 0.338, 0.339, 0.283, 0.319, 0.382, 0.316, 0.341, 0.338, 0.305, 0.339, 0.283, 0.337, 0.341, 0.291, 0.301, 0.275, 0.379, 0.322, 0.349, 0.334, 0.412, 0.338, 0.359, 0.292, 0.412, 0.382, 0.369, 0.351, 0.393, 0.335, 0.296, 0.475, 0.817, 0.603, 0.690, 0.530, 0.526, 0.548, 0.557, 0.378, 0.809, 0.602, 0.594, 0.675, 0.510, 0.584, 0.672, 0.487, 0.748, 0.617, 0.651, 0.465, 0.494, 0.633, 0.566, 0.411, 0.505, 0.613, 0.447, 0.397, 0.612, 0.839, 0.442, 0.694, 0.679, 0.838, 0.677, 0.637, 0.578, 0.506, 0.600, 0.874, 0.806, 0.726, 0.727, 0.577, 0.622, 0.803, 0.686, 0.898, 0.566, 0.694, 0.476, 0.463, 0.490, 0.606, 0.387, 0.472, 0.453, 0.561, 0.379, 0.426, 0.518, 0.745, 0.709, 0.805, 0.777, 0.777, 0.561, 0.560, 0.583, 0.585, 0.343, 0.671, 0.580, 0.580, 0.578, 0.598, 0.625, 0.798, 0.489, 0.701, 0.671, 0.674, 0.493, 0.516, 0.659, 0.568, 0.480, 0.462, 0.508, 0.558, 0.432, 0.410, 0.615, 0.681, 0.444, 0.767, 0.784, 0.885, 0.575, 0.648, 0.562, 0.513, 0.384, 0.902, 0.661, 0.648, 0.706, 0.459, 0.646, 0.739, 0.577, 0.598, 0.602, 0.347, 0.353, 0.540, 0.597, 0.369, 0.402, 0.473, 0.560, 0.389, 0.446, 0.518, 0.832, 0.614, 0.854, 0.754, 0.747, 0.518, 0.696, 0.722, 0.911, 0.792, 0.704, 0.611, 0.711, 0.406, 0.440, 0.597, 0.839, 0.574, 0.722, 0.783, 0.659, 0.666, 0.775, 0.588, 0.502, 0.477, 0.555, 0.599, 0.731, 0.579, 0.684, 0.661, 0.691, 0.640, 0.863, 0.770, 0.723, 0.691, 0.816, 0.686, 0.678, 0.641, 0.835, 0.663, 0.697, 0.627, 0.564, 0.413, 0.691, 0.441, 0.772, 0.615, 0.964, 0.612, 0.523, 0.923, 0.961, 0.573, 0.643, 0.664, 0.963, 0.549, 0.413, 0.791, 0.876, 0.504, 0.878, 0.966, 0.963, 0.499, 0.739, 0.704, 0.568, 0.890, 0.681, 0.915, 0.751, 0.760, 0.828, 0.963, 0.640, 0.966, 0.759, 0.961, 0.596, 0.544, 0.783, 0.916, 0.717, 0.597, 0.485, 0.935, 0.466, 0.510, 0.725, 0.930, 0.621, 0.785, 0.554, 0.962, 0.564, 0.786, 0.906, 0.863, 0.759, 0.761, 0.677, 0.952, 0.622, 0.736, 0.837, 0.942, 0.645, 0.815, 0.770, 0.966, 0.560, 0.573, 0.915, 0.917, 0.624, 0.887, 0.579, 0.948, 0.490, 0.624, 0.743, 0.964, 0.498, 0.805, 0.821, 0.880, 0.584, 0.592, 0.965, 0.869, 0.657, 0.883, 0.692, 0.964, 0.656, 0.557, 0.909, 0.836, 0.541, 0.943, 0.733, 0.945, 0.640, 0.687, 0.806, 0.721, 0.619, 0.671, 0.907, 0.535, 0.544, 0.824, 0.886, 0.569, 0.867, 0.607, 0.960, 0.715, 0.606, 0.966, 0.918, 0.637, 0.932, 0.670, 0.921, 0.599, 0.683, 0.653, 0.918, 0.697, 0.818, 0.594, 0.928, 0.741, 0.543, 0.960, 0.849, 0.725, 0.819, 0.526, 0.966, 0.501, 0.560, 0.710, 0.775, 0.873, 0.873, 0.673, 0.948, 0.696, 0.881, 0.822, 0.964, 0.722, 0.940, 0.692, 0.963, 0.610, 0.728, 0.932, 0.955, 0.657, 0.927, 0.606, 0.855, 0.763, 0.775, 0.731, 0.850, 0.574, 0.706, 0.781, 0.961, 0.710, 0.657, 0.714, 0.867, 0.735, 0.917, 0.820, 0.909, 0.752, 0.745, 0.962, 0.901, 0.752, 0.952, 0.578, 0.620, 0.491, 0.520, 0.893, 0.557, 0.803, 0.809, 0.952, 0.572, 0.474, 0.798, 0.922, 0.835, 0.864, 0.539, 0.842, 0.451, 0.606, 0.769, 0.965, 0.656, 0.723, 0.527, 0.939, 0.471, 0.474, 0.834, 0.773, 0.652, 0.627, 0.527, 0.966, 0.488, 0.557, 0.900, 0.953, 0.716, 0.964, 0.732, 0.950, 0.726, 0.926, 0.808, 0.854, 0.818, 0.847, 0.708, 0.966, 0.688, 0.683, 0.788, 0.838, 0.461, 0.846, 0.443, 0.408, 0.363, 0.264, 0.230, 0.177, 0.144, 0.094, 0.061, 0.042, 0.034, 0.021, 0.013, 0.005, 0.003, 0.003, 0.002, 0.002, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001}
	// fmt.Println(data)

	d := mydraw.NewDrawRGBA()
	d.Computed = data1
	img := d.Draw()
	// img := mydraw.Draw(6, 1, 5, data1)
	// Encode results as PNG to stdout
	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}

func run() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &conf.Conf)
	if err != nil {
		log.Fatal(err)
	}
	slot.InitDB()

	router := gin.Default()

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	router.GET("/t/:path", func(c *gin.Context) {
		path := c.Param("path")
		url := slot.FindUrlBySlot(path)
		log.Printf("%#v\n", url)
		if url == nil {
			c.Redirect(http.StatusSeeOther, conf.Conf.Url.Home)
			return
		}
		if url.Expired() {
			c.Redirect(http.StatusTemporaryRedirect, conf.Conf.Url.Home)
			return
		}
		log.Println(url.Origin)
		c.Redirect(http.StatusMovedPermanently, url.Origin)
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.POST("/t/gen", func(c *gin.Context) {
		var uri slot.Uri
		err := c.BindJSON(&uri)
		if err != nil {
			log.Println(err)
			printRtn(c, "100000", fmt.Sprintf("参数错误:%s", err.Error()), "")
			return
		}
		url := slot.NewUrl(uri.Url, conf.Conf.Gen.Unique)
		if url == nil {
			printRtn(c, "100002", "创建失败", "")
			return
		}
		printRtn(c, "000000", "创建成", fmt.Sprintf(conf.Conf.Gen.Base, url.Count))
	})

	//画波形图
	router.POST("/t/draw", func(c *gin.Context) {
		var musicData mydraw.Music
		err := c.BindJSON(&musicData)
		if err != nil {
			log.Println(err)
			printRtn(c, "110000", fmt.Sprintf("参数错误:%s", err.Error()), "")
			return
		}
		if musicData.Data == "" || musicData.FileName == "" {
			printRtn(c, "110001", "data or filename is null", "")
			return
		}

		file, err := os.Create(musicData.FileName)
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
			printRtn(c, "110003", "创建png图片文件失败", "")
			return
		}
		printRtn(c, "000000", "图片生成成功", "")
	})

	//运行服务
	router.Run(fmt.Sprintf("%s:%d", conf.Conf.Http.IP, conf.Conf.Http.Port))
}

func str2float(data string) []float64 {
	tmp := data[1 : len(data)-1]
	arr := strings.Split(tmp, ",")
	f := make([]float64, len(arr))
	for i, s := range arr {
		fl, _ := strconv.ParseFloat(s, 64)
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

// func createHandle(w http.ResponseWriter, r *http.Request) {
// 	var response UrlResponse
// 	r.ParseForm()
// 	uri := fmt.Sprintf("%s://%s", r.FormValue("protocol"), r.FormValue("url"))
// 	if !IsUrl(uri) {
// 		response.Ok = false
// 		response.Tips = "url不符合规范或域名无法解析"
// 	} else {
// 		url := ifth.NewUrl(uri, config.Url.Unique)
// 		if url == nil {
// 			response.Ok = false
// 			response.Tips = "创建失败，请联系管理员"
// 		} else {
// 			response.Ok = true
// 			response.Url = fmt.Sprintf(config.Url.Base, url.Slot)
// 		}
// 	}
// 	w.Header().Set("Content-type", "text/html")
// 	t := template.Must(template.ParseFiles("./templates/index.html"))
// 	t.Execute(w, response)
// }

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

func main() {
	// test1()
	run()
}
