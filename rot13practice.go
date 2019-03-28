package main

import (
	"io"
	"os"
	"strings"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {

	for {
		x, err := rot13.r.Read(b)
		fmt.Printf("x = %v err = %v \n", x, err)
		if err == io.EOF {
			break
		}
	}	

	return 0, nil
}

func main() {

	// s is the string read out as a string of bytes
	s := strings.NewReader("Lbh penpxrq gur pbqr!")

	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
