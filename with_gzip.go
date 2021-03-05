package blobfactory

import (
	"io"

	"github.com/zatte/gzipcloser"
)

// WithGzip adds gzip compresison to the writers that is returned by the underlying WriterFactory
func WithGzip(wf WriterFactory) WriterFactory {
	return func(path string) (io.WriteCloser, error) {
		w, err := wf(path)
		if err != nil {
			return nil, err
		}
		return gzipcloser.NewWriter(w), err
	}
}

// WithGzip adds gzip compresison to the writers that is returned by the underlying WriterFactory
func (wf WriterFactory) WithGzip() WriterFactory {
	return WithGzip(wf)
}

// WithDecompressGzip unpacks gzip streams from the reader that is returned by the underlying Reader Factory
func WithDecompressGzip(rf ReaderFactory) ReaderFactory {
	return func(path string) (r io.ReadCloser, err error) {
		r, err = rf(path)
		if err != nil {
			return nil, err
		}
		return gzipcloser.NewReader(r)
	}
}

// WithGzip adds gzip compresison to the writers that is returned by the underlying WriterFactory
func (rf ReaderFactory) WithGzip() ReaderFactory {
	return WithDecompressGzip(rf)
}
