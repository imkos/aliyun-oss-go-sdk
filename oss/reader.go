package oss

import "io"

// with size Reader
type SReader struct {
	io.Reader
	Size int64
}

func NewSReader(r io.Reader, size int64) *SReader {
	return &SReader{
		Reader: r,
		Size:   size,
	}
}
