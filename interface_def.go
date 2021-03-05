package blobfactory

import "io"

// WriterFactory yeilds a new WriteCloser under the given path.
type WriterFactory func(path string) (wc io.WriteCloser, err error)

// ReaderFactory returns a new Reader to a given path.
// if the path does not exists it should return os.ErrNotExist
type ReaderFactory func(path string) (wc io.ReadCloser, err error)
