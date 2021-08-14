package controller

import (
	"fileserver/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type DLCtrl struct {
	FileStore *service.FileStore
}

func SetupDL(r *gin.Engine, folder string) {
	ctrl := &DLCtrl{
		FileStore: &service.FileStore{
			Folder: folder,
		},
	}
	r.GET("/download/:file", ctrl.Download)
	r.GET("/downloadlarge/:file", ctrl.DownloadLargeFile)
}

func (c *DLCtrl) Download(ctx *gin.Context) {
	fileName := ctx.Param("file")
	filePath := c.FileStore.Folder + fileName
	file, err := os.Open(filePath)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	defer file.Close()
	ctx.Writer.Header().Add("Content-Type", "multipart/form-data")
	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
}

func (c *DLCtrl) DownloadLargeFile(ctx *gin.Context) {
	fileName := ctx.Param("file")
	filePath := c.FileStore.Folder + fileName
	err := c.FileStore.GenerateFile(filePath, 1*1024*1024*1024)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	file, err := os.Open(filePath)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	defer file.Close()
	ctx.Writer.Header().Add("Content-Type", "multipart/form-data")
	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
}
