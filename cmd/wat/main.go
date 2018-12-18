package main

import (
	"bytes"
	"flag"
	"github.com/go-interpreter/wagon/wasm"
	"github.com/go-interpreter/wagon/wast"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetPrefix("wast: ")
	log.SetFlags(0)

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	fname := flag.Arg(0)

	process(fname)
}

func process(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open %q: %v", fname, err)
	}
	defer f.Close()

	m, err := wasm.ReadModule(f, nil)
	if err != nil {
		log.Fatalf("could not read module: %v", err)
	}
	buf := new(bytes.Buffer)
	err = wast.WriteTo(buf, m)
	if err != nil {
		log.Fatalf("could not write wast to buffer: %v", err)
	}
	tname := strings.TrimSuffix(fname, ".wasm") + ".wat"
	err = ioutil.WriteFile(tname, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("could not write .wat file: %v", err)
	}
}
