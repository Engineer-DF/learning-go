package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
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

func copyFile(srcPath, dstPath string) error {
	openedFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer openedFile.Close()

	createdFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer createdFile.Close()

	_, err = io.Copy(createdFile, openedFile)
	if err != nil {
		return err
	}

	return nil
}

func concatReaders(readers ...io.Reader) (string, error) {
	data, err := io.ReadAll(io.MultiReader(readers...))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func writeToMulti(data string, writers ...io.Writer) error {
	_, err := io.MultiWriter(writers...).Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func readAllWithEOF(r io.Reader) ([]byte, error) {
	var result bytes.Buffer
	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf)

		if n > 0 {
			result.Write(buf[:n])
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
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

	// Задание 4
	err = copyFile("input.txt", "output.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Задание 5

	firstPart := strings.NewReader("I love tea! ")
	secondPart := strings.NewReader("And...I love kittens! ")
	thirdPart := strings.NewReader("I think that`s all.")

	concatedMessage, err := concatReaders(firstPart, secondPart, thirdPart)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(concatedMessage)

	// Задание 6
	var firstBuf bytes.Buffer
	var secondBuf bytes.Buffer
	writeToMulti("Hello, MultiWriter!", &firstBuf, &secondBuf)

	fmt.Println(&firstBuf, &secondBuf)

	// Задание 7
	someMessage := strings.NewReader("Meow! <3")
	buffer3, err := readAllWithEOF(someMessage)
	fmt.Println(string(buffer3))
}
