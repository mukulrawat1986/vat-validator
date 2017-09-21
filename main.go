package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	flag.Parse()
	in := flag.Args()[0]
	Run(in, os.Stdout)
}

// Run is the function where all the action happens
func Run(in string, out io.Writer) {
	// countryCode, vatNumber := SplitVatNumber(in)
}

// SplitVatNumber splits the vat number into country code
// and vat number.
func SplitVatNumber(in string) (string, string) {
	return in[:2], in[2:]
}
