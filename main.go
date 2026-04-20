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
			out, err := gachikoi.Translate(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println(out)
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

	out, err := gachikoi.Translate(string(in))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Print(out)
}
