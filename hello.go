package main

import (
	"gocv.io/x/gocv"
)

go func main() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Aaaooba -  Primeiro teste com OpenCV e GO !!!")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
