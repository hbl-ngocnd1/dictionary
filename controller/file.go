package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hblab-ngocnd/csv-demo/csv_reder"
	"github.com/labstack/echo/v4"
)

type FileHandler struct {
	fileService interface{}
}

func NewFileHandler() *FileHandler {
	return &FileHandler{
		fileService: nil,
	}
}
func (f *FileHandler) UploadFiles(c echo.Context) error {
	return c.Render(http.StatusOK, "upload.html", map[string]interface{}{"router": "upload"})
}
func (f *FileHandler) ApiUpload(c echo.Context) error {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	fileList := make([]string, 0, len(form.File["files"]))
	for _, file := range form.File["files"] {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		defer src.Close()
		// Destination
		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		fileList = append(fileList, file.Filename)
	}
	defer func() {
		for _, file := range fileList {
			os.Remove(file)
		}
	}()
	if len(fileList) < 2 {
		fmt.Println(fileList)
		return c.JSON(http.StatusForbidden, map[string]string{"error": "required more file"})
	}
	result, err := csv_reder.DecodeData(fileList[0], fileList[1])
	if err != nil {
		log.Fatal(err)
	}
	prettyPrint(result)
	return c.JSON(http.StatusOK, result)
}

func prettyPrint(in map[string]string) {
	b, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b))
}
