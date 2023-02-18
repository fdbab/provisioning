package main

import (
	"log"
	"os"
)

func main() {
	fileUrl := os.Getenv("DOWNLOAD_URL")
	fileName := os.Getenv("DOWNLOAD_FILE")

	load(fileUrl, fileName)
	if isTarGz(fileName) {
		log.Println("Entpacke tar ball")
		untar(fileName)
	} else {
		errorLogger := log.New(os.Stderr, "", 0)
		errorLogger.Println("Datei ist kein tar ball")
	}

}
