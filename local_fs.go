package blobfactory

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// NewLocal returns a writer factory that creates local files "basePath".
// If basePath == "" it defaults to current working dir ("./")
func NewLocal(basePath string, permMode fs.FileMode, fileFlags int) (WriterFactory, ReaderFactory) {
	if basePath == "" {
		basePath = "./"
	}

	wf := func(path string) (wc io.WriteCloser, err error) {
		os.MkdirAll(filepath.Dir(filepath.Join(basePath, path)), permMode)
		return os.OpenFile(basePath+path, fileFlags, permMode)
	}

	rf := func(path string) (r io.ReadCloser, err error) {
		os.MkdirAll(filepath.Dir(filepath.Join(basePath, path)), permMode)
		return os.Open(basePath + path)
	}

	return wf, rf
}

// NewDefaultLocal creates a new writers factory using
// NewLocal(basePath, os.ModePerm, os.O_CREATE | os.O_APPEND | os.O_RDWR)
func NewDefaultLocal(basePath string, permMode fs.FileMode, fileFlags int) (WriterFactory, ReaderFactory) {
	return NewLocal(basePath, os.ModePerm, os.O_CREATE|os.O_APPEND|os.O_RDWR)
}
