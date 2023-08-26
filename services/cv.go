package services

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
)

func TakeSnapshot() {
	deviceID := 0
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", 0)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return
	}
	if img.Empty() {
		fmt.Printf("no image on device %v\n", deviceID)
		return
	}

	currentTime := time.Now()
	timeFormated := currentTime.Format("20060102150405")
	tmpFile := "tmp/" + timeFormated + ".jpg"

	gocv.IMWrite(tmpFile, img)
}
