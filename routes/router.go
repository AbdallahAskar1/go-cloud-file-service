package router

import (
	handlers "github.com/AbdallahAskar1/go-cloud-file-service/handler"
	"github.com/AbdallahAskar1/go-cloud-file-service/services/storage"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	s := storage.NewS3Storage()

	fileHandler := handlers.NewFileHandler(s, s, s)

	router.POST("/upload", fileHandler.UploadFile)
	router.GET("/download/:key", fileHandler.DownloadFile)
	router.GET("/files", fileHandler.ListFiles)
}
