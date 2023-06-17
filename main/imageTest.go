package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0, 0, 25, 0xFF},
	color.RGBA{255, 0, 25, 0xFF},
}

const (
	whiteIndex = 0
	blackIndex = 1
	RedIndex   = 2
)

func Lissajous(out io.Writer) {
	const (
		cycles  = 5     //完整的x振荡器变化的个数 , 几个周期的波
		res     = 0.001 //角度分辨率
		size    = 100   //图像画布包含[-size..+size]
		nframes = 64    //动画中的帧数
		delay   = 8     //以10ms为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes} // gif.GIF 是一个结构体 gif包里面的 里面有一个LoopCount字段 用来指定动画循环的次数 帧数
	phase := 0.0                        //相位
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) //image.Rect 是一个函数 返回一个image.Rectangle类型的结构体
		img := image.NewPaletted(rect, palette)      //image.NewPaletted 是一个函数 返回一个image.Paletted类型的结构体
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Cos(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), RedIndex) //设置黑点
		}
		phase += 0.1                           //相位差
		anim.Delay = append(anim.Delay, delay) //动画的帧间延迟
		anim.Image = append(anim.Image, img)   //动画的帧
	}
	gif.EncodeAll(out, &anim) //注意：忽略编码错误
}

//
//func main() {
//	rand.Seed(time.Now().UTC().UnixNano())
//	lissajous(os.Stdout)
//}
