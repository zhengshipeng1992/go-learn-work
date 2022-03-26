package main

import (
	"io"
	"os"
)

type ErrWriter struct {
	w io.Writer
	err error
}

func (ew *ErrWriter) Write(buf []byte)  {
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.Write(buf)
}

func Test() error {
	p0 := []byte("123 ")
	p1 := []byte("hello world")
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	ew := ErrWriter{w: file}
	ew.Write(p0)
	ew.Write(p1)
	if ew.err != nil {
		return ew.err
	}
	return nil
}
