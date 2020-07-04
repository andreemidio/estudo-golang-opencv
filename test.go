package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gocv.io/x/gocv"
)

func main() {
	// read image
	pizzaPath := filepath.Join("PnoWu.png")
	pizza := gocv.IMRead(pizzaPath, gocv.IMReadColor)
	if pizza.Empty() {
		fmt.Printf("Failed to read image: %s\n", pizzaPath)
		os.Exit(1)
	}
	// Convert BGR to HSV image (dont modify the original)
	hsvPizza := gocv.NewMat()
	gocv.CvtColor(pizza, &hsvPizza, gocv.ColorBGRToHSV)
	pizzaChannels, pizzaRows, pizzaCols := hsvPizza.Channels(), hsvPizza.Rows(), hsvPizza.Cols()
	// define HSV color upper and lower bound ranges
	lower := gocv.NewMatFromScalar(gocv.NewScalar(110.0, 50.0, 50.0, 0.0), gocv.MatTypeCV8UC3)
	upper := gocv.NewMatFromScalar(gocv.NewScalar(130.0, 255.0, 255.0, 0.0), gocv.MatTypeCV8UC3)
	// split HSV lower bounds into H, S, V channels
	lowerChans := gocv.Split(lower)
	lowerMask := gocv.NewMatWithSize(pizzaRows, pizzaCols, gocv.MatTypeCV8UC3)
	lowerMaskChans := gocv.Split(lowerMask)
	// split HSV lower bounds into H, S, V channels
	upperChans := gocv.Split(upper)
	upperMask := gocv.NewMatWithSize(pizzaRows, pizzaCols, gocv.MatTypeCV8UC3)
	upperMaskChans := gocv.Split(upperMask)
	// copy HSV values to upper and lower masks
	for c := 0; c < pizzaChannels; c++ {
		for i := 0; i < pizzaRows; i++ {
			for j := 0; j < pizzaCols; j++ {
				lowerMaskChans[c].SetUCharAt(i, j, lowerChans[c].GetUCharAt(0, 0))
				upperMaskChans[c].SetUCharAt(i, j, upperChans[c].GetUCharAt(0, 0))
			}
		}
	}
	gocv.Merge(lowerMaskChans, &lowerMask)
	gocv.Merge(upperMaskChans, &upperMask)
	// global mask
	mask := gocv.NewMat()
	gocv.InRange(hsvPizza, lowerMask, upperMask, &mask)
	// cut out pizza mask
	pizzaMask := gocv.NewMat()
	gocv.Merge([]gocv.Mat{mask, mask, mask}, &pizzaMask)
	// cut out the pizza and convert back to BGR
	gocv.BitwiseAnd(hsvPizza, pizzaMask, &hsvPizza)
	gocv.CvtColor(hsvPizza, &hsvPizza, gocv.ColorHSVToBGR)
	// write image to filesystem
	outPizza := "no_pizza.jpeg"
	if ok := gocv.IMWrite(outPizza, hsvPizza); !ok {
		fmt.Printf("Failed to write image: %s\n", outPizza)
		os.Exit(1)
	}
	// write pizza mask to filesystem
	outPizzaMask := "no_pizza_mask.jpeg"
	if ok := gocv.IMWrite(outPizzaMask, mask); !ok {
		fmt.Printf("Failed to write image: %s\n", outPizza)
		os.Exit(1)
	}
}
