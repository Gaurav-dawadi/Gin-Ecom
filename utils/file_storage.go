package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func FileSystemStorage(file multipart.File, uploadFile *multipart.FileHeader) string {
	filename := uploadFile.Filename
	out, err := os.Create(filepath.Join("media", filepath.Base(filename)))
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println(err)
	}

	filepath := "media/" + filename
	return filepath
}
