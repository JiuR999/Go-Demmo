package img

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func ChangeColor(path string) {
	// 打开原始图片
	inputFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// 解码图片
	img, err := jpeg.Decode(inputFile)
	if err != nil {
		panic(err)
	}

	// 创建一个新的灰度图像
	bounds := img.Bounds()
	sketch := image.NewGray(bounds)

	// 将图片素描化
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// 获取原始图像的颜色
			originalColor := img.At(x, y)
			// 转换为灰度
			r, g, b, a := originalColor.RGBA()
			fmt.Println(r, g, b, a)
			gray := (r + g + b) / 3
			// 设置素描图像的灰度值
			sketch.Set(x, y, color.Gray{uint8(gray >> 8)})
		}
	}

	// 保存素描图像
	outputFile, err := os.Create("sketch.jpg")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// 编码并保存为JPEG格式
	jpeg.Encode(outputFile, sketch, nil)

}

// 判断r四周是否有任意两个连通白块，如果没有，则认为是噪点
func imgConnect(ru, rd, rl, rr uint32) bool {
	if (ru < 0x7788 && (rl < 0x7788 || rr < 0x7788)) || (rd < 0x7788 && (rl < 0x7788 || rr < 0x7788)) || (rl < 0x7788 && (ru < 0x7788 || rd < 0x7788)) || (rr < 0x7788 && (ru < 0x7788 || rd < 0x7788)) {
		return true
	}
	return false
}

// 二值化
func ImgBarbarization(path string) image.Image {
	// 打开原始图片
	inputFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	// 解码图片
	img, err := jpeg.Decode(inputFile)
	binImg := image.NewGray16(img.Bounds())
	draw.Draw(binImg, binImg.Bounds(), img, img.Bounds().Min, draw.Over)
	rect := binImg.Bounds()
	// 遍历点像素点
	for x := 0; x < rect.Dx(); x++ {
		for y := 0; y < rect.Dy(); y++ {
			// 获取颜色
			r, _, _, _ := binImg.At(x, y).RGBA()
			ru, _, _, _ := img.At(x, y+1).RGBA()
			rd, _, _, _ := img.At(x, y-1).RGBA()
			rl, _, _, _ := img.At(x-1, y).RGBA()
			rr, _, _, _ := img.At(x+1, y).RGBA()
			if r < 0x7788 && imgConnect(ru, rd, rl, rr) {
				binImg.Set(x, y, color.White)
			} else {
				binImg.Set(x, y, color.Black)
			}
		}
	}
	// 保存素描图像
	outputFile, err := os.Create("sketch.jpg")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// 编码并保存为JPEG格式
	jpeg.Encode(outputFile, binImg, nil)
	return binImg
}

func Pre(path string) {
	// 打开图像文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 解码图像
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// 创建一个新的图像用于存储分割结果
	bounds := img.Bounds()
	result := image.NewGray(bounds)

	// 自适应阈值分割
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// 获取当前像素的颜色
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			// 计算阈值（可以根据需要调整）
			//threshold := uint8(128) // 固定阈值
			//TODO 自适应阈值
			//半径 r 常量 C
			threshold := getHold(img, x, y)
			if c.Y > threshold {
				result.SetGray(x, y, color.Gray{255}) // 白色
			} else {
				result.SetGray(x, y, color.Gray{0}) // 黑色
			}
		}
	}

	// 保存结果图像
	outFile, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	if err := jpeg.Encode(outFile, result, nil); err != nil {
		panic(err)
	}
}

func getHold(img image.Image, x int, y int) uint8 {
	r := 3
	c := 12
	var sum uint8 = 0
	//var sr, sg, sb, sa uint8
	for i := x - r; i < x+r; i++ {
		for j := y - r; j < y+r; j++ {
			if i > 0 && i < img.Bounds().Max.X &&
				j > 0 && j < img.Bounds().Max.Y {

				c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
				sum += c.Y
				/*				r, g, b, a := img.At(x, y).RGBA()
								sr += uint8(r >> 8)
								sg += uint8(g >> 8)
								sb += uint8(b >> 8)
								sa += uint8(a >> 8)*/
			}
		}
	}
	return sum/uint8(2*r) - uint8(c)
}
