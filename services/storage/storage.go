package storage

import (
	"io"
)
type Uploader interface {
	Upload(fileName string, file io.Reader, ContentType string) (string, error)
}

type Downloader interface {
	Download(key string) (io.Reader, error)
}
type ListFiles interface {
	ListAll() (map[string]string, error)
}

type FileStorage interface {
	Uploader
	Downloader
	ListFiles
}
