package client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Upload(url string, filename string) error {
	contextType, body, err := LoadFile(filename)
	if err != nil {
		return nil
	}
	resp, err := http.Post(url, contextType, body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Body)
	return nil
}

func LoadFile(filename string) (string, *bytes.Buffer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base("./"+filename))
	if err != nil {
		return "", nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", nil, err
	}
	contentType := writer.FormDataContentType()
	writer.Close()
	return contentType, body, nil
}

func Download(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
