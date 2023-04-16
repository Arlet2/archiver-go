package main

// #include <stdio.h>
// void print() {
// 	printf("Hello, wowo!\n");
// }
import "C"
import (
	"archiver/internal/compressor"
	"archiver/internal/compressor/encoding"
	"archiver/internal/readers"
	"archiver/internal/writers"
	"fmt"
	"os"
)

func main() {
	C.print()

	if len(os.Args) == 1 {
		fmt.Println("Введите название файл для сжатия в аргументах запуска")
		return
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	lines, err := readers.ReadFile(file)
	_ = file.Close()

	freqDict := encoding.CreateFreqDictFromLines(lines)

	/*
		for key, value := range freqDict {
			fmt.Print("\"", string(key), "\"", ": ", value, "\n")
		}
	*/

	codingDict := encoding.CreateCodingDict(encoding.CreateCodingTree(freqDict))

	lines, err = compressor.Compress(lines, codingDict)

	fmt.Println(lines)

	nfile, err := os.Create("compressed-test")
	defer nfile.Close()

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	err = writers.WriteFile(nfile, lines)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
}
