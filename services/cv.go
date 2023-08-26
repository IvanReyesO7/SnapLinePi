package services

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
)

func TakeSnapshot() (*string, error) {
	deviceID := 0
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", 0)
		return nil, err
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return nil, err
	}
	if img.Empty() {
		fmt.Printf("no image on device %v\n", deviceID)
		return nil, err
	}

	currentTime := time.Now()
	timeFormated := currentTime.Format("20060102150405")
	tmpFile := "tmp/" + timeFormated + ".jpg"

	gocv.IMWrite(tmpFile, img)
	return &tmpFile, nil
}
