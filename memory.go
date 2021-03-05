package blobfactory

import (
	"bytes"
	"io"
	"os"
	"sync"
)

// closeCallback is io.Writer with a Close() methods that does nothing
type closeCallback struct {
	io.Writer
	closeCallack func() error
}

func (nc *closeCallback) Close() error {
	if nc.closeCallack == nil {
		return nil
	}
	return nc.closeCallack()
}

// NewMemory returns a WriterFactory which is backed by bytes.Buffer.
func NewMemory() (map[string]*bytes.Buffer, WriterFactory, ReaderFactory) {
	res := map[string]*bytes.Buffer{}
	mutex := sync.Mutex{}
	wf := func(path string) (wc io.WriteCloser, err error) {
		mutex.Lock()
		defer mutex.Unlock()

		_, ok := res[path]
		if !ok {
			res[path] = bytes.NewBuffer(nil)
		}

		return &closeCallback{
			Writer: res[path],
			closeCallack: func() error {
				return nil
			},
		}, nil
	}

	rf := func(path string) (wc io.ReadCloser, err error) {
		mutex.Lock()
		defer mutex.Unlock()

		buffer, ok := res[path]
		if !ok {
			return nil, os.ErrNotExist
		}
		return io.NopCloser(buffer), nil
	}

	return res, wf, rf
}
