package blobfactory

import (
	"io"
	"path"
)

// WithWriterPrefix adds a prefix to all writes conducted by the writerfactory
func WithWriterPrefix(wf WriterFactory, prefix string) WriterFactory {
	return func(writePath string) (io.WriteCloser, error) {
		return wf(path.Join(prefix, writePath))
	}
}

// WithReaderPrefix adds a prefix to all writes conducted by the writerfactory
func WithReaderPrefix(rf ReaderFactory, prefix string) ReaderFactory {
	return func(readPath string) (io.ReadCloser, error) {
		return rf(path.Join(prefix, readPath))
	}
}

// WithPrefix adds a prefix to all writes conducted by the writerfactory
func (wf WriterFactory) WithPrefix(prefix string) WriterFactory {
	return WithWriterPrefix(wf, prefix)
}

// WithPrefix adds a prefix to all writes conducted by the writerfactory
func (rf ReaderFactory) WithPrefix(prefix string) ReaderFactory {
	return WithReaderPrefix(rf, prefix)
}
