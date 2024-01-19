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
	// Verificar se um argumento foi fornecido
	if len(os.Args) < 2 {
		fmt.Println("Uso: <programa> <caminho_para_imagem>")
		os.Exit(1)
	}

	// Pegar o primeiro argumento da linha de comando
	filePath := os.Args[1]

	// Obter o tamanho do arquivo original
	originalSize := fileSizeInBytes(filePath)

	// Carregar imagem do disco
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	// Redimensionar a imagem
	resizedImg := resize.Resize(1000, 0, img, resize.Lanczos3)

	out, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Salvar a imagem em formato PNG
	err = png.Encode(out, resizedImg)
	if err != nil {
		panic(err)
	}

	// Obter o tamanho do arquivo redimensionado
	resizedSize := fileSizeInBytes(filePath)

	// Imprimir os tamanhos dos arquivos
	fmt.Printf("Tamanho original: %d bytes\n", originalSize)
	fmt.Printf("Tamanho após redimensionamento: %d bytes\n", resizedSize)
}
