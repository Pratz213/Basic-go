package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

const (
	FILE_PATH = "/home/pratik/Desktop/"
	FILE_NAME = "Merce-Homepage.html"
	FILE_TYPE = ".zip"
)

func main() {
	// Opens the file
	data, err := os.Open(FILE_PATH + FILE_NAME)
	checkErrIfNil("Error reading data file:", err)
	defer data.Close()

	// Create the ZIP Archive
	archive, err := os.Create(FILE_PATH + FILE_NAME + FILE_TYPE)
	checkErrIfNil("Error creating ZIP file:", err)

	// Close the archive
	defer archive.Close()

	// Create a new ZIP writer
	zipWriter := zip.NewWriter(archive)

	// Create a ZIP entry for the file
	dataFile, err := zipWriter.Create(FILE_NAME)
	checkErrIfNil("Error creating ZIP entry file: ", err)

	// Copy the data into the new dataFile
	_, err = io.Copy(dataFile, data)
	checkErrIfNil("Error compressing file: ", err)

	// Closing the ZIP writer
	zipWriter.Close()

	// Print original file size
	dataInfo, err := data.Stat()
	checkErrIfNil("Error Fetching Orignal file: ", err)
	fmt.Printf("Original file size: %d bytes\n", dataInfo.Size())

	// Print compressed file size
	archiveInfo, err := os.Stat(FILE_PATH + FILE_NAME + FILE_TYPE)
	checkErrIfNil("Error Fetching Orignal file: ", err)
	fmt.Printf("Compressed file size: %d bytes\n", archiveInfo.Size())
}

func checkErrIfNil(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		return
	}
}
