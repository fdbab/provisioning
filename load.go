package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func load(fileUrl string, fileName string) {

	// Create the file
	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	// Get the data
	resp, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Write the data to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
}
