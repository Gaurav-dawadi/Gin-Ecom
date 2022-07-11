package utils

import (
	"go-practice/utils/logger"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Generate Random file name for a file
func generate_filename(filename string) string {
	fileExtension := filepath.Ext(filename)
	random_filename := strconv.FormatInt(time.Now().UnixNano(), 10) + fileExtension
	return random_filename
}

func FileSystemStorage(file multipart.File, uploadFile *multipart.FileHeader) string {
	// Check if directory to store all files exists. If not, first create them.
	dirInfo, err := os.Stat("media")
	if os.IsNotExist(err) {
		if err := os.Mkdir(dirInfo.Name(), 0777); err != nil {
			logger.LogError(err, "Error occured during creation of director")
		}
	}

	filename := filepath.Base(uploadFile.Filename)
	new_filename := generate_filename(filename)

	out, err := os.Create(filepath.Join(dirInfo.Name(), new_filename))
	if err != nil {
		logger.LogError(err, "error occured during creation of file")
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		logger.LogError(err, "error during copying of file")
	}

	filepath := dirInfo.Name() + new_filename
	return filepath
}
