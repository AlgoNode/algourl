package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	err := makeQRFromStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func makeQRFromStdin() error {
	encodedTxn, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	kr, err := makeQRKeyRegRequest(encodedTxn)
	if err != nil {
		return err
	}

	kr.Print()

	return nil

}
