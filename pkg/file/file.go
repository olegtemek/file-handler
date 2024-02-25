package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

func SaveFile(file *multipart.File, handler *multipart.FileHeader) (string, error) {
	currentTime := time.Now().Format(time.RFC3339Nano)
	filepath := fmt.Sprintf("uploads/%s--%s", currentTime, handler.Filename)
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, *file)
	return filepath, nil
}

func DeleteFile(filepath string) error {

	err := os.Remove(filepath)
	if err != nil {
		return err
	}

	return nil
}
