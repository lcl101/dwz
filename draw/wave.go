package draw

import (
	"image"
	"image/color"
)

// DravWave 画波形图接口
type DrawWave interface {
	Draw() image.Image
}

type Music struct {
	Data     string `json:"data"`
	FileName string `json:"filename"`
}

// Wave 波形图结构
type Wave struct {
	ImgY      int        //图片高度
	ScaleY    float64    //线条高度比例0.8
	ScaleX    int        //间隔像素
	LineWidth int        //线条宽度
	Step      int        //采样数据
	Sharpness int        //圆弧化像素
	Computed  []float64  //波形图浮点数据
	MColor    color.RGBA //主颜色
	SColor    color.RGBA //次颜色
}

func (w *Wave) waveLen() int {
	return len(w.Computed) / w.Step
}

func (w *Wave) waveX() int {
	return (w.ScaleX + w.LineWidth) * w.waveLen()
}

func (w *Wave) sxld() int {
	return w.ScaleX + w.LineWidth
}

func (w *Wave) drawLine(x, y, i, halfY int, img *image.RGBA) {
	c := w.MColor
	if i == 0 || i == w.ScaleX-1 {
		c = w.SColor
	}
	peak := w.ScaleX / 2

	adjust := 0
	if i < peak {
		// Adjust downward
		adjust = (i - peak) * w.Sharpness
	} else if i == peak {
		// No adjustment at peak
		adjust = 0
	} else {
		// Adjust downward
		adjust = (peak - i) * w.Sharpness
	}
	// On top half of the image, invert adjustment to create symmetry between
	// top and bottom halves
	if y < halfY {
		adjust = -1 * adjust
	}
	img.Set(x+i, y+adjust, c)
}
