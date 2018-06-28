package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	bufferSize = 1024
)

func main() {
	args := os.Args[1:]

	if args[0] != "" {
		file, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		buf := make([]byte, bufferSize)

		for {
			n, err := file.Read(buf)

			if n > 0 {
				s := Tokenize(buf)
				p := Parse(s)
				g := Generate(p)

				fmt.Println(g)
			} else if err == io.EOF {
				break
			} else if err != nil {
				log.Panic(err)
				break
			}
		}
	}
}
