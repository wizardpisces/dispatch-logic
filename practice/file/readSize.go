// go run readSize.go 1000 /bin/ls | wc
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)

	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return buffer[0:n]
}
func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("<buffer size> <filename>")
		return
	}

	bufferSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// file := os.Args[2]
	f, err := os.Open("./testfile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for {
		readData := readSize(f, bufferSize)
		if readData != nil {
			fmt.Println(string(readData))
		} else {
			break
		}
	}
}