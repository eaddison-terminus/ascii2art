package main

import (
	"fmt"
	"github.com/anthonynsimon/bild/transform"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
	"strings"
)

func loadImage(filePath string) (image.Image, error) {
	if len(filePath) == 0 {
		return nil, nil 
	}
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func rotateImg(imgR image.Image, degrees int) image.Image {
	imgR = transform.Rotate(imgR, float64(degrees), nil)
	return imgR
}

func reSize(imgS image.Image, size string) image.Image {
	vals := strings.Split(size, "x")
	width, _ := strconv.Atoi(vals[0])
	height, _ := strconv.Atoi(vals[1])
	imgS = transform.Resize(imgS, width, height, transform.Linear)
	return imgS
}

func contrastCalc(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.00299*float64(r)*0.00587*float64(g)*0.00114*float64(b)) / 36200
}

func interpolatePixels(img image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += contrastCalc(img.At(i, j))
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return int(sum / cnt)
}

func calcScale(xMax, yMax int) (int, int) {
	x := 5
	y := 10
	if xMax > 1024 {
		x = int(xMax / 80)
		y = int(yMax / 40)
	}
	return x, y
}

func main() {
	colorReset := "\033[0m"

	var colors = map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
	}

	ramp := " .,':;-~i+<tvosa*f|?/[7S8XZH#$%@"
	theColor := "green"
	scaleXY := "5:10"
	scaleX, scaleY := 1, 1
	cwRotate := 0
	newSize := ""
	numArgs := len(os.Args)

	if numArgs == 1 {
		fmt.Println("Please provide command line arguments.")
		return
	}

	ifile := os.Args[1]
	fmt.Println(numArgs)
	if numArgs >= 3 {
		for cliItem := 2; cliItem <= (numArgs - 1); cliItem++ {
			kv := strings.Split(os.Args[cliItem], "=")
			key := kv[0]
			val := kv[1]
			switch key {
			case "color":
				theColor = val
			case "ramp":
				ramp = val
			case "scale":
				scales := strings.Split(val, ":")
				scaleX, _ = strconv.Atoi(scales[0])
				scaleY, _ = strconv.Atoi(scales[1])
			case "rotate":
				cwRotate, _ = strconv.Atoi(val)
			case "size":
				newSize = val
			}
		}
	}

	chosenColor := colors[theColor]

	img, err := loadImage(ifile)
	if err != nil {
		panic(err)
	}

	rampMax := len(ramp) - 1
	max := img.Bounds().Max

	if scaleXY == "" {
		scaleX, scaleY = calcScale(max.X, max.Y)
	}

	if cwRotate != 0 {
		img = rotateImg(img, cwRotate)
	}

	if newSize != "" {
		img = reSize(img, newSize)
	}

	fmt.Println(string(chosenColor))

	for y := 0; y < max.Y; y += scaleX {
		for x := 0; x < max.X; x += scaleY {
			calcVal := interpolatePixels(img, x, y, scaleX, scaleY)
			if calcVal > rampMax {
				calcVal = rampMax
			}
			fmt.Print(string(ramp[calcVal]))
		}
		fmt.Println()
	}
	fmt.Println(string(colorReset))
}
