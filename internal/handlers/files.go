package handlers

import (
	"files-api/internal/config"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const UploadsDir = "uploads"
const PublicDir = "public"

func UploadFile(c *gin.Context) {
	foldeName := c.Param("folder_name")

	file, header, err := c.Request.FormFile("file")
	if file == nil || err != nil {
		c.JSON(http.StatusBadRequest, config.SignResponse{
			Success: false,
			Message: "No se proporcionó ningún archivo",
			Error:   err.Error(),
		})
		return
	}
	defer file.Close()

	dir := filepath.Join(UploadsDir, foldeName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, config.SignResponse{
			Success: false,
			Message: "No se pudo crear el directorio",
			Error:   err.Error(),
		})
		return
	}

	ext := filepath.Ext(header.Filename)
	randomName := fmt.Sprintf("%x%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(dir, randomName)

	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.SignResponse{
			Success: false,
			Message: "Error al crear el archivo",
			Error:   err.Error(),
		})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.SignResponse{
			Success: false,
			Message: "Error al guardar el archivo",
			Error:   err.Error(),
		})
		return
	}

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = http.DetectContentType([]byte(ext))
	}

	// Success response
	c.JSON(http.StatusOK, config.SignResponse{
		Success: true,
		Message: "Archivo subido correctamente",
		Data: config.FileData{
			Filename:         randomName,
			Path:             fmt.Sprintf("/uploads/%s/%s", foldeName, randomName),
			OriginalFilename: header.Filename,
			Size:             header.Size,
			MimeType:         mimeType,
		},
	})
}

func UploadFileToPublic(c *gin.Context) {
	foldeName := c.Param("folder_name")

	file, header, err := c.Request.FormFile("file")
	if file == nil || err != nil {
		c.JSON(http.StatusBadRequest, config.SignResponse{
			Success: false,
			Message: "No se proporcionó ningún archivo",
			Error:   err.Error(),
		})
		return
	}
	defer file.Close()

	dir := filepath.Join(PublicDir, UploadsDir, foldeName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, config.SignResponse{
			Success: false,
			Message: "No se pudo crear el directorio",
			Error:   err.Error(),
		})
		return
	}

	ext := filepath.Ext(header.Filename)
	randomName := fmt.Sprintf("%x%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(dir, randomName)

	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.SignResponse{
			Success: false,
			Message: "Error al crear el archivo",
			Error:   err.Error(),
		})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.SignResponse{
			Success: false,
			Message: "Error al guardar el archivo",
			Error:   err.Error(),
		})
		return
	}

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = http.DetectContentType([]byte(ext))
	}

	// Success response
	c.JSON(http.StatusOK, config.SignResponse{
		Success: true,
		Message: "Archivo subido correctamente al directorio público",
		Data: config.FileData{
			Filename:         randomName,
			Path:             fmt.Sprintf("/uploads/%s/%s", foldeName, randomName),
			OriginalFilename: header.Filename,
			Size:             header.Size,
			MimeType:         mimeType,
		},
	})
}

func DownloadFile(c *gin.Context) {
	foldeName := c.Param("folder_name")
	fileName := c.Param("file_name")

	filePath := filepath.Join(UploadsDir, foldeName, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, config.SignResponse{
			Success: false,
			Message: "Archivo no encontrado",
			Error:   err.Error(),
		})
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	c.Header("Content-Type", contentType)

	c.File(filePath)
}
