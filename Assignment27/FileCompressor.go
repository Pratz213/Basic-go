package main

import (
	"archive/zip"
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	FILE_NAME = "Merce-Homepage.html"
	FILE_PATH = "/home/pratik/Desktop/"
	FILE_TYPE = ".zip"
)

func main() {
	// Ask the user for a URL
	fmt.Print("Enter URL: ")
	reader := bufio.NewReader(os.Stdin)
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	//Ignore SSL Verification by setting http.InsecureSkipVerify to true
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer resp.Body.Close()

	// Create and save the HTML file
	htmlFile, err := os.Create(FILE_PATH + FILE_NAME)
	if err != nil {
		fmt.Println("Error creating HTML file:", err)
		return
	}
	defer htmlFile.Close()

	_, err = io.Copy(htmlFile, resp.Body)
	if err != nil {
		fmt.Println("Error saving HTML file:", err)
		return
	}

	fmt.Println("HTML file saved as ", FILE_NAME)

	// Opens the file
	data, err := os.Open(FILE_PATH + FILE_NAME)
	if err != nil {
		fmt.Println("Error reading data file:", err)
		return
	}
	defer data.Close()

	// Create the ZIP Archive
	archive, err := os.Create(FILE_PATH + FILE_NAME + FILE_TYPE)
	if err != nil {
		fmt.Println("Error creating ZIP file:", err)
		return
	}
	defer archive.Close()

	// Create a new ZIP writer
	zipWriter := zip.NewWriter(archive)

	// Create a ZIP entry for the file
	dataFile, err := zipWriter.Create(FILE_NAME)
	if err != nil {
		fmt.Println("Error creating ZIP entry file:", err)
		return
	}

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
