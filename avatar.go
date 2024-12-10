package avatargen

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math/rand"

	"golang.org/x/image/draw"
)

const arraySize int = 12

type Avatar struct {
	matrix [arraySize][arraySize]int
	color  color.RGBA
}

func New() *Avatar {
	avatar := &Avatar{}
	avatar.Generate()
	return avatar
}

func (a *Avatar) generateMatrix() {
	for row := 1; row < arraySize-1; row++ {
		for col := 1; col < arraySize/2; col++ {
			value := rand.Intn(2)
			a.matrix[row][col] = value
			a.matrix[row][arraySize-col-1] = value
		}
	}
}

func (a *Avatar) generateColor() {
	a.color = color.RGBA{
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		0xff,
	}
}

func (a *Avatar) Generate() {
	a.generateMatrix()
	a.generateColor()
}

func (a *Avatar) Print() {
	for row := 0; row < arraySize; row++ {
		for col := 0; col < arraySize; col++ {
			value := a.matrix[row][col]
			if value == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print("██")
			}
		}
		fmt.Println()
	}
}

func (a *Avatar) ToBuffer(size int) *bytes.Buffer {
	src := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{12, 12}})

	for row := 0; row < arraySize; row++ {
		for col := 0; col < arraySize; col++ {
			value := a.matrix[row][col]

			if value == 1 {
				src.SetRGBA(col, row, a.color)
			} else {
				src.Set(col, row, color.White)
			}
		}
	}

	resizedImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{size, size}})

	draw.NearestNeighbor.Scale(resizedImage, resizedImage.Rect, src, src.Bounds(), draw.Over, nil)

	buffer := new(bytes.Buffer)

	jpeg.Encode(buffer, resizedImage, &jpeg.Options{Quality: 100})

	return buffer
}
