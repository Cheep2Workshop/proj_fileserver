package service

import (
	"fmt"
	"time"
)

type FileStore struct {
	Folder string
}

func DefaultFileStore() *FileStore {
	return &FileStore{
		Folder: defaultFolder,
	}
}

const defaultFolder = "files/"

func (store *FileStore) VarifyFileName(filename string) string {
	stamp := time.Now().Unix()
	return fmt.Sprintf("%s%d_%s", store.Folder, stamp, filename)
}
