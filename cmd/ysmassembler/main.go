package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"io"
	"log"
	"os"

	"github.com/yuchiki/YuchikiStackMachine/pkg/instruction"
)

func main() {
	outputFileName := flag.String("o", "", "output file")
	flag.Parse()
	/*	inputFile := flag.Arg(0)

		inputLines, err = readLines(inputFile)
	*/
	// trim
	// convert
	sampleCode := []uint64{
		uint64(instruction.OpPushI)<<32 + 5, // pushi 5
		uint64(instruction.OpPushI)<<32 + 7, // pushi 7
		uint64(instruction.OpAdd) << 32,     // add
		uint64(instruction.OpRet) << 32,     // ret

	}

	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	writeInstructions(outputFile, sampleCode)

}

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var lines []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines, nil
}

func writeInstructions(output io.Writer, instructions []uint64) {
	for _, inst := range instructions {
		binary.Write(output, binary.BigEndian, inst)
	}
}
