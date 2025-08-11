// internal/handlers/files.go
package handlers

import (
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

func UploadFile(c *gin.Context) {
	issueID := c.Param("issue_id")

	file, header, err := c.Request.FormFile("file")
	if file == nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se proporcionó ningún archivo"})
		return
	}
	defer file.Close()

	dir := filepath.Join(UploadsDir, issueID)
	if err := os.MkdirAll(dir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el directorio"})
		return
	}

	ext := filepath.Ext(header.Filename)
	randomName := fmt.Sprintf("%x%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(dir, randomName)

	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el archivo"})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el archivo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            "success",
		"filename":          randomName,
		"path":              fmt.Sprintf("/uploads/%s/%s", issueID, randomName),
		"original_filename": header.Filename,
		"size":              header.Size,
	})
}

func DownloadFile(c *gin.Context) {
	issueID := c.Param("issue_id")
	fileName := c.Param("file_name")

	filePath := filepath.Join(UploadsDir, issueID, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found or not accessible"})
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	c.Header("Content-Type", contentType)

	c.File(filePath)
}
