package blobfactory_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zatte/blobfactory"
)

// TestNewMemory memory back writerfactory (and reader factory)
func TestNewMemory(t *testing.T) {
	m, wf, rf := blobfactory.NewMemory()

	testCase1Path := "subpath"
	testCase1Content := []byte("testing string")

	writer, err := wf(testCase1Path)
	assert.NoError(t, err, "MemoryWriterFactory should never yield errors")
	writer.Write(testCase1Content)
	_, err = rf(testCase1Path)
	require.NoError(t, err)
	assert.Equal(t, testCase1Content, m[testCase1Path].Bytes(), "MemWriter factory should store the results in buffers")

	testCase2Path := "subpath2"
	testCase2Content := []byte("testing string2")
	writer, err = wf(testCase2Path)
	assert.NoError(t, err, "MemoryWriterFactory should never yield errors")
	writer.Write(testCase2Content)
	_, err = rf(testCase2Path)
	require.NoError(t, err)
	assert.Equal(t, testCase2Content, m[testCase2Path].Bytes(), "MemWriter factory should store the results in buffers")

	_, err = rf("this path does not exists")
	assert.Error(t, err, os.ErrNotExist)
}
