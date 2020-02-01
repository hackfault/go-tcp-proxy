package main

import (
	"fmt"
	"log"
	"os"
	"io"
)

type FooReader struct{}

func (f *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (f *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data\n")
	}
}
