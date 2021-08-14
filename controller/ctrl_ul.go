package controller

import (
	"fileserver/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ULCtrl struct {
	FileStore *service.FileStore
}

func SetupULCtrl(r *gin.Engine, folder string) {
	ctrl := &ULCtrl{
		FileStore: &service.FileStore{
			Folder: folder,
		},
	}
	r.POST("/upload", ctrl.UploadFile)
}

func (c *ULCtrl) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	log.Printf("File:%s\n", file.Filename)

	// file name
	dst := c.FileStore.VarifyFileName(file.Filename)
	log.Println(dst)
	// file save
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "File upload succeed.")
}
