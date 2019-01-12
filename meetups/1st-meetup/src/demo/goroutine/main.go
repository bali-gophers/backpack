package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func buildFileName(fullURL string) (string, error) {
	fileURL, err := url.Parse(fullURL)
	if err != nil {
		return "", err
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	return segments[len(segments)-1], nil
}

func downloadFile(url string) (string, error) {
	fmt.Printf("Downloading %s ...\n", url)
	fileName, err := buildFileName(url)
	if err != nil {
		return "", err
	}
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	httpClient := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("Download \"%s\" successfully\n", fileName)
	return fileName, nil
}

func multipleErrorsToString(errs []error) string {
	var errString string
	for _, err := range errs {
		errString = errString + ", " + err.Error()
	}
	return errString
}

func downloadMultipleFiles(urls []string) ([]string, error) {
	result := make(chan string, len(urls))
	errch := make(chan error, len(urls))
	for _, url := range urls {
		go func(url string) {
			fileName, err := downloadFile(url)
			if err != nil {
				errch <- err
				result <- ""
				return
			}
			result <- fileName
			errch <- nil
		}(url)
	}
	results := make([]string, 0)
	var errs []error
	for i := 0; i < len(urls); i++ {
		results = append(results, <-result)
		if err := <-errch; err != nil {
			errs = append(errs, err)
		}
	}
	var err error
	if len(errs) > 0 {
		err = errors.New(multipleErrorsToString(errs))
	}
	return results, err
}

func main() {
	urls := []string{
		"https://upload.wikimedia.org/wikipedia/commons/7/77/Pocket_gopher.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/c/cb/Pocket-Gopher_Ano-Nuevo-SP.jpg",
		"https://independent.media.clients.ellingtoncms.com/img/croppedphotos/2016/05/16/gopher-drawing-1_t958.jpg",
	}
	files, err := downloadMultipleFiles(urls)
	if err != nil {
		fmt.Printf("Errors: %v \n", err)
	}
	fmt.Printf("Finished downloading %d files \n", len(files))
}
