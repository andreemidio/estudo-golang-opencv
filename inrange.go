//https://stackoverflow.com/questions/49241490/gocv-how-to-cut-out-an-image-from-blue-background-using-opencv
//referência

package main

import (
	"fmt"
	_ "fmt"
	"os"
	"path/filepath"

	"gocv.io/x/gocv"
)

func main() {
	file := filepath.Join("PnoWu.png")
	image := gocv.IMRead(file, gocv.IMReadAnyColor)
	window := gocv.NewWindow("Aaaooba -  Teste InRange !!!")

	if image.Empty() {
		fmt.Printf("\nErro ao carregar a imagem\n", file)
		os.Exit(1)
	}

	img := gocv.NewMat()

	gocv.CvtColor(image, &img, gocv.ColorBGRToHSV)
	imageChannels, imageRows, imageCols := img.Channels(), img.Rows(), img.Cols()

	lower := gocv.NewMatFromScalar(gocv.NewScalar(110.0, 50.0, 50.0, 0.0), gocv.MatTypeCV8UC3)
	upper := gocv.NewMatFromScalar(gocv.NewScalar(130.0, 255.0, 255.0, 0.0), gocv.MatTypeCV8UC3)

	lowerChans := gocv.Split(lower)
	lowerMask := gocv.NewMatWithSize(imageRows, imageCols, gocv.MatTypeCV8UC3)
	lowerMaskChans := gocv.Split(lowerMask)

	upperChans := gocv.Split(upper)
	upperMask := gocv.NewMatWithSize(imageRows, imageCols, gocv.MatTypeCV8UC3)
	upperMaskChans := gocv.Split(upperMask)

	for c := 0; c < imageChannels; c++ {
		for i := 0; i < imageRows; i++ {
			for j := 0; j < imageCols; j++ {
				lowerMaskChans[c].SetUCharAt(i, j, lowerChans[c].GetUCharAt(0, 0))
				upperMaskChans[c].SetUCharAt(i, j, upperChans[c].GetUCharAt(0, 0))
			}
		}
	}

	gocv.Merge(lowerMaskChans, &lowerMask)
	gocv.Merge(upperMaskChans, &upperMask)

	mask := gocv.NewMat()
	gocv.InRange(img, lowerMask, upperMask, &mask)

	pmask := gocv.NewMat()
	gocv.Merge([]gocv.Mat{mask, mask, mask}, &pmask)

	gocv.BitwiseAnd(img, pmask, &img)
	gocv.CvtColor(img, &img, gocv.ColorHSVToBGR)

	//outp := "Aoooba.jpg"

	for {
		window.IMShow(image)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
