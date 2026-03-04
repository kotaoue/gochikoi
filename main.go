package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kotaoue/gochikoi/pkg/gachikoi"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, arg := range args {
			fmt.Println(gachikoi.Translate(arg))
		}
		return
	}

	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(in) == 0 {
		return
	}

	fmt.Print(gachikoi.Translate(string(in)))
}
