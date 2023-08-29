package tasks

import (
	"log"
	"os"
	"time"
)

func CleanTempDir() {
	time.Sleep(time.Second)
	RemoveTempFiles()
	log.Println("[TASK] Sleeping for 30 minutes")
	ticker := time.NewTicker(30 * time.Minute)
	for range ticker.C {
		RemoveTempFiles()
		log.Println("[TASK] Sleeping for 30 minutes")
	}
}

func RemoveTempFiles() {
	log.Println("[TASK] Starting cleaner task...")
	tmpPath := os.Getenv("PWD") + "/tmp"

	files, err := os.ReadDir(tmpPath)
	if err != nil {
		log.Println("[TASK] Error reading directory:", err)
		return
	}

	if len(files) == 0 {
		log.Println("[TASK] tmp directory is empty")
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fileName := file.Name()
			os.Remove(tmpPath + "/" + fileName)
			log.Printf("[TASK] %s file removed succesfully", fileName)
		}
	}
	return
}
