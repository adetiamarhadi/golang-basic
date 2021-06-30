package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error:", "Bad Arguments")
		os.Exit(1)
	}

	fn := args[1]

	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, f)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
