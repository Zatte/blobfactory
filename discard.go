package blobfactory

import (
	"io"
)

// NewDiscarder returns a WriterFactory which is discoards all writes
func NewDiscarder() WriterFactory {
	return func(path string) (wc io.WriteCloser, err error) {
		return &closeCallback{
			Writer: io.Discard,
			closeCallack: func() error {
				return nil
			},
		}, nil
	}
}
