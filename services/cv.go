package services

import (
	"fmt"
	"os"
	"sync"
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

func TakeClip() (*string, *string, error) {
	deviceID := 0
	webcam, err := gocv.OpenVideoCapture(deviceID)

	currentTime := time.Now()
	timeFormated := currentTime.Format("20060102150405")
	preview := "tmp/" + timeFormated + ".jpg"
	tmpFile := "tmp/" + timeFormated + ".mp4"

	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", 0)
		return nil, nil, err
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return nil, nil, err
	}
	gocv.IMWrite(preview, img)
	writer, err := gocv.VideoWriterFile(tmpFile, "H264", 25, img.Cols(), img.Rows(), true)
	if err != nil {
		fmt.Printf("error opening video writer device: %v\n", tmpFile)
		return nil, nil, err
	}
	defer writer.Close()

	for i := 0; i < 100; i++ {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return nil, nil, err
		}
		if img.Empty() {
			continue
		}

		writer.Write(img)
	}
	return &tmpFile, &preview, nil
}

func DeleteFile(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("Error deleting file %s: %v\n", filePath, err)
	} else {
		fmt.Printf("File %s has been deleted successfully.\n", filePath)
	}
}
