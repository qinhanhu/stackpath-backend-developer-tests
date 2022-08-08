package main

import (
	"bytes"
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type errorDetectTestCase struct {
	Name     string
	Input    []byte
	Expected string
}

func TestErrorDetect(t *testing.T) {
	testCases := []errorDetectTestCase{
		{
			Name:     "Detected error in multilines",
			Input:    []byte("error1\nhello\nerror2\n"),
			Expected: "error1\nerror2\n",
		},
		{
			Name:     "No error detected",
			Input:    []byte("hello\n"),
			Expected: "",
		},
	}

	for _, tc := range testCases {
		InputProcess(t, tc)
	}

	// test large size input
	t.Run("Detected Random Large Data", func(t *testing.T) {
		randomData := make([]byte, 100000000) // 100MB
		_, _ = rand.Read(randomData)
		tc := errorDetectTestCase{
			Name:     "Detected Random Large Data",
			Input:    randomData,
			Expected: "",
		}
		InputProcess(t, tc)
	})
	t.Run("Detected Large Data with error", func(t *testing.T) {
		randomData := make([]byte, 1000000000) // 1GB
		_, _ = rand.Read(randomData)
		copy(randomData[100:], "\n#here is an error#\n")
		tc := errorDetectTestCase{
			Name:     "Detected Large Data with error",
			Input:    randomData,
			Expected: "#here is an error#\n",
		}
		InputProcess(t, tc)
	})
}

func InputProcess(t *testing.T, tc errorDetectTestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		tmpfile, err := ioutil.TempFile("", "test_example")
		if err != nil {
			log.Fatal(err)
		}

		defer os.Remove(tmpfile.Name())
		if _, err := tmpfile.Write(tc.Input); err != nil {
			log.Fatal(err)
		}
		if _, err := tmpfile.Seek(0, 0); err != nil {
			log.Fatal(err)
		}
		oldStdin := os.Stdin
		defer func() {
			os.Stdin = oldStdin // Restore original Stdin
		}()
		os.Stdin = tmpfile
		buf := &bytes.Buffer{}
		ErrorDetect(buf)
		assert.Equal(t, tc.Expected, buf.String())
	})
}
