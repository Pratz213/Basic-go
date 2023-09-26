package main

import (
	"archive/zip"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	FILE_NAME = "Merce-Homepage.html"
	FILE_PATH = "/home/pratik/Desktop/"
	FILE_TYPE = ".zip"
)

func main() {

	// Ask the user for a URL
	url := os.Args[1]
	if len(os.Args) != 2 {
		fmt.Print("Invalid Url")
		os.Exit(1)
	}

	//Ignore SSL Verification by setting http.InsecureSkipVerify to true
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	urlData, err := http.Get(url)
	checkErrIsNil("Error fetching the URL:", err)

	defer urlData.Body.Close()

	// Create and save the HTML file
	htmlFile, err := os.Create(FILE_PATH + FILE_NAME)
	checkErrIsNil("Error creating HTML file:", err)
	
	defer htmlFile.Close()

	_, err = io.Copy(htmlFile, urlData.Body)
	checkErrIsNil("Error saving HTML file:", err)
	fmt.Println("HTML file saved as Merce-Homepage.html")

	// Opens the file
	data, err := os.Open(FILE_PATH + FILE_NAME)
	checkErrIsNil("Error reading data file:", err)
	defer data.Close()

	// Create the ZIP Archive
	archive, err := os.Create(FILE_PATH + FILE_NAME + FILE_TYPE)
	checkErrIsNil("Error creating ZIP file:", err)

	// Close the archive
	defer archive.Close()

	// Create a new ZIP writer
	zipWriter := zip.NewWriter(archive)

	// Create a ZIP entry for the file
	dataFile, err := zipWriter.Create(FILE_NAME)
	checkErrIsNil("Error creating ZIP entry file:", err)

	// Copy the data into the new dataFile
	_, err = io.Copy(dataFile, data)
	if err != nil {
		fmt.Println("Error compressing file:", err)
		return
	}

	// Close the ZIP writer
	zipWriter.Close()

	// Print original file size
	dataInfo, _ := data.Stat()
	fmt.Printf("Original file size: %d bytes\n", dataInfo.Size())

	// Print compressed file size
	archiveInfo, _ := os.Stat(FILE_PATH + FILE_NAME + FILE_TYPE)
	fmt.Printf("Compressed file size: %d bytes\n", archiveInfo.Size())
}

// Prints Error!
func checkErrIsNil(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
