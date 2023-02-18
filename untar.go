package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"strings"
)

func isTarGz(s string) bool {
	return strings.HasSuffix(s, ".tar.gz")
}

func untar(fileName string) {
	// Öffne das Archiv
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Entpacke die gzipped Datei
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	defer func(gzipReader *gzip.Reader) {
		err := gzipReader.Close()
		if err != nil {

		}
	}(gzipReader)

	// Entpacke das Tar-Archiv
	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break // Ende des Archivs
		}

		if err != nil {
			panic(err)
		}

		// Extrahiere die Datei
		switch header.Typeflag {
		case tar.TypeDir:
			// Verzeichnis erstellen
			if err := os.Mkdir(header.Name, os.FileMode(header.Mode)); err != nil {
				panic(err)
			}
		case tar.TypeReg:
			// Reguläre Datei extrahieren
			outFile, err := os.Create(header.Name)
			if err != nil {
				panic(err)
			}
			defer func(outFile *os.File) {
				err := outFile.Close()
				if err != nil {

				}
			}(outFile)
			if _, err := io.Copy(outFile, tarReader); err != nil {
				panic(err)
			}
		default:
			// Andere Typen ignorieren
		}
	}
}
