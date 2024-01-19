package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func fileSizeInBytes(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}
	return fileInfo.Size()
}

func main() {

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	filePath := os.Args[1]

	originalSize := fileSizeInBytes(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	resizedImg := resize.Resize(1000, 0, img, resize.Lanczos3)

	out, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = png.Encode(out, resizedImg)
	if err != nil {
		panic(err)
	}

	resizedSize := fileSizeInBytes(filePath)

	fmt.Printf("Tamanho original: %d bytes\n", originalSize)
	fmt.Printf("Tamanho após redimensionamento: %d bytes\n", resizedSize)
}
