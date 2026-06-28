package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func copyData(src io.Reader, dst io.Writer) error {
	_, err := io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}

func copyWithLimit(src io.Reader, dst io.Writer, limit int64) error {
	tempBuf := make([]byte, 1)

	_, err := io.Copy(dst, io.LimitReader(src, limit))
	if err != nil {
		return err
	}

	n, _ := src.Read(tempBuf)
	if n > 0 {
		return fmt.Errorf("limit exceeded")
	}

	return nil
}

func teeDemo() {
	st := strings.NewReader("Go is great!")
	var buffer bytes.Buffer

	io.Copy(io.Discard, io.TeeReader(st, &buffer))
	fmt.Println(&buffer)
}

func main() {
	// Задание 1
	expression := strings.NewReader("Hello, World!")
	var buffer bytes.Buffer

	copyData(expression, &buffer)
	fmt.Println(&buffer)

	// Задание 2
	data := strings.NewReader("abcdefghij")
	var buffer2 bytes.Buffer

	err := copyWithLimit(data, &buffer2, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(&buffer2)

	// Задание 3
	teeDemo()
}
