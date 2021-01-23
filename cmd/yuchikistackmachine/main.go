package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"io/ioutil"
	"log"

	"github.com/yuchiki/YuchikiStackMachine/pkg/runner"
)

func main() {
	codeFile := flag.String("code", "", "filename of binary")
	codeBytes, err := ioutil.ReadFile(*codeFile)
	if err != nil {
		log.Fatal(err)
	}

	codeBinary := toUint64Array(codeBytes)

	r := runner.NewRunner(codeBinary)

	r.Run()
}

func toUint64Array(bs []byte) []uint64 {
	buf := bytes.NewBuffer(bs)

	var uint64Arr []uint64
	for {
		ui, err := binary.ReadUvarint(buf)
		if err != nil {
			break
		}
		uint64Arr = append(uint64Arr, ui)
	}
	return uint64Arr
}
