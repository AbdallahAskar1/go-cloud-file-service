package handlers

import (
	"github.com/AbdallahAskar1/go-cloud-file-service/services/storage"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	uploader   storage.Uploader
	downloader storage.Downloader
	list       storage.ListFiles
}

func NewFileHandler(uploader storage.Uploader, downloader storage.Downloader, lister storage.ListFiles) *FileHandler {
	return &FileHandler{
		uploader:   uploader,
		downloader: downloader,
		list:       lister,
	}
}

func (h *FileHandler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file provided"})
		return
	}
	defer file.Close()

	_, err = h.uploader.Upload(header.Filename, file, header.Header.Get("Content-Type"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upload file: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "File uploaded successfully", "key": header.Filename})
}

func (h *FileHandler) DownloadFile(c *gin.Context) {
	key := c.Param("key")

	reader, err := h.downloader.Download(key)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to download file"})
		return
	}

	c.DataFromReader(200, -1, "application/octet-stream", reader, nil)
}

func (h *FileHandler) ListFiles(c *gin.Context) {
	output, err := h.list.ListAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list files"})
		return
	}
	fileKeys := make([]string, 0, len(output))
	for key := range output {
		fileKeys = append(fileKeys, key)
	}

	c.JSON(200, gin.H{"files": fileKeys})
}
