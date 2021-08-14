package service

import (
	"fmt"
	"os"
	"path/filepath"
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

func (store *FileStore) DeleteAll() error {
	d, err := os.Open(store.Folder)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(store.Folder, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (store *FileStore) GenerateFile(path string, size int64) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := f.Truncate(size); err != nil {
		return err
	}
	return nil
}
