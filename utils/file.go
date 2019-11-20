package utils

import (
	"mime/multipart"
	"os"

	"github.com/micro/go-micro/util/log"
)

//GetFileMsg ...
func GetFileMsg(file *multipart.FileHeader) (string, []byte, error) {
	fileBytes := make([]byte, file.Size)

	src, err := file.Open()
	if err != nil {
		return file.Filename, fileBytes, err
	}
	defer src.Close()

	n, err := src.Read(fileBytes)
	if err != nil {
		return file.Filename, fileBytes, err
	}
	log.Info("filename: ", file.Filename, ", size: ", n)

	return file.Filename, fileBytes, err
}

//SaveFile ...
func SaveFile(filePath string, content []byte) error {
	newFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	n, err := newFile.Write(content)
	if err != nil {
		return err
	}

	log.Info("save file: ", filePath, ", len:", n)

	return nil
}
