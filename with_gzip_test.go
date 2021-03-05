package blobfactory

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type closer interface {
	Close() error
}

func TestWithGzipWriterFactory(t *testing.T) {
	_, rawWf, rawRf := NewMemory()
	gzipWf := rawWf.WithGzip()
	gzipRf := rawRf.WithGzip()

	testString := []byte("Yay, compression for the people")

	// Test 1 read back compressed stream
	writer, err := gzipWf("key1")
	require.NoError(t, err)
	_, err = writer.Write(testString)
	require.NoError(t, err)

	rr, err := rawRf("key1")
	require.NoError(t, err, "expected to be able to find written keys")
	buffer := bytes.Buffer{}
	_, err = buffer.ReadFrom(rr)
	require.NoError(t, err, "Expected to be able to read raw bytes")
	assert.NotEqual(t, testString, buffer, "gzipped content should not match uncompressed content")

	// Tets 2; Add another key
	writer, err = gzipWf("key2")
	require.NoError(t, err)
	_, err = writer.Write(testString)
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	rr, err = rawRf("key2")
	require.NoError(t, err, "expected to be able to find written keys")
	buffer = bytes.Buffer{}
	_, err = buffer.ReadFrom(rr)
	require.NoError(t, err, "Expected to be able to read raw bytes")
	assert.NotEqual(t, testString, buffer, "gzipped content should not match uncompressed content")

	// Test 3, read back gzip
	writer, err = gzipWf("key3")
	require.NoError(t, err, "should be able to create a writer")
	_, err = writer.Write(testString)
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	gzR, err := gzipRf("key3")
	require.NoError(t, err, "expected to be able to create gzip reader")
	buffer = bytes.Buffer{}
	_, err = buffer.ReadFrom(gzR)
	require.NoError(t, err, "expected to be able to find written keys in gzip reader")

	assert.EqualValues(t, buffer.Bytes(), testString)

}
