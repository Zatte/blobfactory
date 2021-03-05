# WriterFactory
Methods to create (index/keyed) io.Writer(s) for various storage options.


[![Go Report Card](https://goreportcard.com/badge/github.com/zatte/blobfactory?style=flat-square)](https://goreportcard.com/report/github.com/zatte/blobfactory)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/zatte/blobfactory)

## License
Apache 2.0

## Forked
Forked and rewritten & cleaned version of [github.com/kvanticoss/goutils/blobfactory](https://github.com/kvanticoss/goutils/tree/master/blobfactory)
(remove most external deps and some cleanup)

## Why?
Most of the factories are very short and simple with some sane defaults. Used by span writers (split writes into files ABC00001.json.gz to ABC99999.json.gz where each is at most 100Mb) with pluggable backends.

## Usage
```golang

  // local files
  readerFactory, writerFactory := NewDefaultLocal("/tmp/interimfiles")
  _, readerFactory, writerFactory := NewMem() // first arg is the map[string]*bytes.Buffer containing the raw data without mutex protection


  //create a writer under "file.txt", and another one under "file2.txt"
  f1, err := writerFactory("file.txt")
  f2, err := writerFactory("file2.txt")

  f1.Write(...)
  f2.Write(...)
  f1.Close()
  f2.Close()

  r, err := readerFactory("file.txt")
  r.Read(...)
```
