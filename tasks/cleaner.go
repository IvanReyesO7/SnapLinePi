package tasks

import (
	"fmt"
	"log"
	"os"
)

func CleanTempFiles() {
	log.Println("Starting cleaner task...")
	tmpPath := os.Getenv("PWD") + "/tmp" // specify the directory

	files, err := os.ReadDir(tmpPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			os.Remove(tmpPath + "/" + file.Name())
		}
	}
}
