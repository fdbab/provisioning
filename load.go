package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var numBytes int64 = 0

func load(fileUrl string, fileName string) {

	// Datei erstellen
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// HTTP-Request erstellen
	req, err := http.NewRequest("GET", fileUrl, nil)
	if err != nil {
		panic(err)
	}

	// HTTP-Client erstellen
	client := &http.Client{}

	// HTTP-Response lesen
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Datei-Inhalt in Datei schreiben
	written, err := io.Copy(file, io.TeeReader(resp.Body, &progressWriter{}))
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nHerunterladen abgeschlossen: %d Bytes heruntergeladen\n", written)
}

type progressWriter struct{}

func (pw *progressWriter) Write(p []byte) (int, error) {
	numBytes += int64(len(p))
	fmt.Printf("\rHerunterladen... %d Bytes heruntergeladen", bytesToMegabytes(numBytes))
	return len(p), nil
}

func bytesToMegabytes(bytes int64) string {
	return fmt.Sprintf("%.2f MB", float64(bytes)/1024/1024)
}
