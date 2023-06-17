package main

import (
	"image/color"
	"math/cmplx"
)

func main() {
	//const (
	//	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	//	width, height          = 1024, 1024
	//)
	//img := image.NewRGBA(image.Rect(0, 0, width, height))
	//for py := 0; py < height; py++ { //每一行 从上到下 逐行扫描 1024行 1024列 一共1024*1024个点 也就是1024*1024个像素点 也就是1024*1024个复数 也就是1024*1024个虚数 也就是1024*1024个实数
	//	y := float64(py)/height*(ymax-ymin) + ymin //每一行的y值 从上到下 逐行扫描
	//	for px := 0; px < width; px++ {            //每一列 从左到右 逐列扫描
	//		x := float64(px)/width*(xmax-xmin) + xmin //每一列的x值 从左到右 逐列扫描
	//		z := complex(x, y)                        //每一个点的复数值
	//		// 图像点(px, py)表示复数值z
	//		img.Set(px, py, mandelbrot(z)) //每一个点的复数值对应的像素点的颜色
	//	}
	//
	//}
	//png.Encode(os.Stdout, img) // NOTE: ignoring errors

}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ { //迭代次数
		v = v*v + z           //迭代公式
		if cmplx.Abs(v) > 2 { //如果迭代公式的结果的绝对值大于2
			return color.Gray{255 - contrast*n} //返回颜色
		}
	}

	return color.Black //返回颜色
}
