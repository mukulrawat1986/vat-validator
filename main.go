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

}
