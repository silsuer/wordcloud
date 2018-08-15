package main

import (
	"image/color"
	"gitee.com/bangbaoshi/wordcloud"
	"os"
	"strings"
)

func renderNow(textList []string, path string, ttfPath string, backPath string) {
	//angles := []int{0, 90}
	angles := []int{0}
	colors := []*color.RGBA{
		&color.RGBA{0xbb, 0xbb, 0xbb, 0xbb},
		&color.RGBA{0x33, 0x33, 0x33, 0x33},
		//&color.RGBA{0x73, 0x73, 0x0, 0xff},
	}
	//"./fonts/heiti.TTF" ./imgs/wordcloud.png
	render := wordcloud_go.NewWordCloudRender(60, 20,
		ttfPath,
		backPath, textList, angles, colors, path)
	// 输出现在的时间
	render.Render()
}

func main() {

	// boot/boot 1/2/3/5/6  /images/img ttfPath backPath

	renderNow(strings.Split(os.Args[1], "/"), os.Args[2], os.Args[3], os.Args[4])
}
