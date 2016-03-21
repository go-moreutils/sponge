package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	if len(os.Args) != 2 || os.Args[1] == "-h" {
		fmt.Fprintln(os.Stderr, "sponge <file>: soak up all input from stdin and write it to <file>")
		os.Exit(0)
	}

	b, _ := ioutil.ReadAll(os.Stdin)
	f, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening output file: %s", err)
		os.Exit(1)
	}
	defer f.Close()
	_, err = f.Write(b)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error writing output file: %s", err)
		os.Exit(1)
	}
}
