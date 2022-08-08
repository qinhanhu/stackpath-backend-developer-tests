package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Read STDIN into a new buffered reader
	ErrorDetect(os.Stdout)
}

func ErrorDetect(writer io.Writer) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := read(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf(" a real error happended here: %v\n", err)
		}

		if bytes.Contains(line, []byte("error")) {
			if _, err := fmt.Fprintln(writer, string(line)); err != nil {
				log.Fatalf(" a real error happended here: %v\n", err)
			}
		}
	}
}

func read(r *bufio.Reader) ([]byte, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return ln, err
}

// Reference:
// https://devmarkpro.com/working-big-files-golang
// https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it
