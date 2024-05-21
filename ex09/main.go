package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wget <url>")
		return
	}

	url := os.Args[1]
	err := downloadFile("downloaded_page.html", url)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}

	fmt.Println("Downloaded page saved as downloaded_page.html")
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
