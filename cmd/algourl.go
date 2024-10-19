package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hashmapsdata2value/algourl/encoder"
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

	kr, err := encoder.MakeQRKeyRegRequest(encodedTxn)
	if err != nil {
		return err
	}

	kr.Print()

	return nil

}
