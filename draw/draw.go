package draw

import (
	"image"
	"image/color"
	"math"
)

// DrawRgb 画透明图
type DrawRgb struct {
	Wave
}

// NewDrawAlpha new一个实例
func NewDrawRGBA() *DrawRgb {
	d := &DrawRgb{}
	d.ImgY = 50
	d.ScaleX = 5
	d.ScaleY = 0.8
	d.Sharpness = 1
	d.LineWidth = 2
	d.Step = 5
	d.MColor = color.RGBA{0x66, 0x66, 0x66, 0xff}
	d.SColor = color.RGBA{0x99, 0x99, 0x99, 0xff}
	return d
}

// Draw 实现画图
func (d *DrawRgb) Draw() image.Image {
	waveX := d.waveX()
	waveY := d.ImgY
	img := image.NewRGBA(image.Rect(0, 0, waveX, waveY))
	bounds := img.Bounds()
	// Calculate halfway point of Y-axis for image
	imgHalfY := bounds.Max.Y / 2

	// Values to be used for repeated computations
	var scaleComputed, halfScaleComputed int
	intBoundY := int(bounds.Max.Y)
	f64BoundY := float64(bounds.Max.Y)
	// intSharpness := int(sharpness)

	// Begin iterating all computed values
	x := 0
	for n := 0; n < d.waveLen(); n++ {
		scaleComputed = int(math.Floor(d.Computed[n*d.Step] * f64BoundY * d.ScaleY))
		halfScaleComputed = scaleComputed / 2

		// 设置透明度
		for y := 0; y < intBoundY; y++ {
			// If X-axis is being scaled, draw background over several X coordinates
			for i := 0; i < d.sxld(); i++ {
				img.Set(x+i, y, color.RGBA{0x0, 0x0, 0x0, 0})
			}
		}

		for y := imgHalfY - halfScaleComputed; y < scaleComputed+(imgHalfY-halfScaleComputed); y++ {
			for i := 0; i < d.ScaleX; i++ {
				d.drawLine(x, y, i, imgHalfY, img)
			}
		}

		x += d.sxld()
	}
	return img
}

// Draw takes a slice of computed values and generates
// a waveform image from the input.
func Draw(scaleX, scaleY int, step int, computed []float64) image.Image {
	// Store integer scale values
	intScaleX := int(scaleX)
	intScaleY := int(scaleY)

	imgYDefault := 50
	scaleDefault := 0.8
	scaleClipping := false
	sharpness := 1
	// Calculate maximum n, x, y, where:
	//  - n: number of computed values
	//  - x: number of pixels on X-axis
	//  - y: number of pixels on Y-axis
	maxN := len(computed) / step
	maxX := maxN * intScaleX
	maxY := imgYDefault * intScaleY

	// Create output, rectangular image
	img := image.NewRGBA(image.Rect(0, 0, maxX, maxY))
	bounds := img.Bounds()

	// Calculate halfway point of Y-axis for image
	imgHalfY := bounds.Max.Y / 2

	// Calculate a peak value used for smoothing scaled X-axis images
	peak := int(math.Ceil(float64(scaleX))/2) / 2

	// Calculate scaling factor, based upon maximum value computed by a SampleReduceFunc.
	// If option ScaleClipping is true, when maximum value is above certain thresholds
	// the scaling factor is reduced to show an accurate waveform with less clipping.
	imgScale := scaleDefault
	if scaleClipping {
		// Find maximum value from input slice
		var maxValue float64
		for _, c := range computed {
			if c > maxValue {
				maxValue = c
			}
		}

		// For each 0.05 maximum increment at 0.30 and above, reduce the scaling
		// factor by 0.25.  This is a rough estimate and may be tweaked in the future.
		for i := 0.30; i < maxValue; i += 0.05 {
			imgScale -= 0.25
		}
	}

	// Values to be used for repeated computations
	var scaleComputed, halfScaleComputed, adjust int
	intBoundY := int(bounds.Max.Y)
	f64BoundY := float64(bounds.Max.Y)
	intSharpness := int(sharpness)

	// Begin iterating all computed values
	x := 0
	for n := 0; n < len(computed)/step; n++ {
		// Scale computed value to an integer, using the height of the image and a constant
		// scaling factor
		scaleComputed = int(math.Floor(computed[n*step] * f64BoundY * imgScale))

		// Calculate the halfway point for the scaled computed value
		halfScaleComputed = scaleComputed / 2

		// Draw background color down the entire Y-axis
		for y := 0; y < intBoundY; y++ {
			// If X-axis is being scaled, draw background over several X coordinates
			for i := 0; i < intScaleX; i++ {
				img.Set(x+i, y, color.RGBA{0x0, 0x0, 0x0, 0})
			}
		}

		// Iterate image coordinates on the Y-axis, generating a symmetrical waveform
		// image above and below the center of the image
		for y := imgHalfY - halfScaleComputed; y < scaleComputed+(imgHalfY-halfScaleComputed); y++ {
			// If X-axis is being scaled, draw computed value over several X coordinates
			for i := 0; i < intScaleX/2; i++ {
				// When scaled, adjust computed value to be lower on either side of the peak,
				// so that the image appears more smooth and less "blocky"
				if i < peak {
					// Adjust downward
					adjust = (i - peak) * intSharpness
				} else if i == peak {
					// No adjustment at peak
					adjust = 0
				} else {
					// Adjust downward
					adjust = (peak - i) * intSharpness
				}

				// On top half of the image, invert adjustment to create symmetry between
				// top and bottom halves
				if y < imgHalfY {
					adjust = -1 * adjust
				}

				// Retrieve and apply color function at specified computed value
				// count, and X and Y coordinates.
				// The output color is selected using the function, and is applied to
				// the resulting image.
				img.Set(x+i, y+adjust, color.RGBA{0x32, 0x32, 0x32, 255})
			}
		}

		// Increase X by scaling factor, to continue drawing at next loop
		x += intScaleX
	}

	// Return generated image
	return img
}
