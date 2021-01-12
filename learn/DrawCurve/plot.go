package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	const (
		dx = 256
		dy = 256
	)

	// 需要保存的文件
	imgcounter := 123
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
	defer imgfile.Close()
	// 新建一个 指定大小的 RGBA位图
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
//golang 这个坐标系是反转的
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			/* 测试平方的图片显示
			if x*x>y*255{
			//if x*x>y*255{
				img.Set(x,y, color.RGBA{100, 100, 0, 255})
			}*/
			if(x == 5){
				img.Set(x, y, color.RGBA{uint8(x), 0, 0, 255})
			}

			//draw uv type
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})

			/*if x%8 == 0 {
				// 设置某个点的颜色，依次是 RGBA
				img.Set(x, y, color.RGBA{uint8(x % 256), 255, 255, 255})
			}*/
		}
	}
	//反转图片 横向
	dst := image.NewRGBA(img.Bounds())
	for x := 0; x < dst.Bounds().Dx(); x++ {
		for y := 0; y < dst.Bounds().Dx(); y++ {
			//flip horizontal
			dst.Set(x, y, img.At(img.Bounds().Dx()-x, y))
			//flip vertical
			dst.Set(x, y, img.At(x, img.Bounds().Dx()-y))
		}
	}
	// 以PNG格式保存文件
	err := png.Encode(imgfile, dst)
	if err != nil {
		log.Fatal(err)
	}

}
