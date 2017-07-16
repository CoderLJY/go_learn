package main

import (
	"io"
	"log"
	"os"
	"strings"	
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")
	
	if _, err := io.CopyN(os.Stdout, r, 5); err != nil {
		log.Fatal(err)
	}
}
