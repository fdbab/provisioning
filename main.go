package main

import "os"

func main() {
	fileUrl := os.Getenv("DOWNLOAD_URL")
	fileName := os.Getenv("DOWNLOAD_FILE")
	load(fileUrl, fileName)
}
