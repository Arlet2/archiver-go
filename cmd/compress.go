package cmd

import (
	"archiver/internal/readers"
	"archiver/internal/writers"
	"archiver/pkg/compressor"
	"archiver/pkg/compressor/encoding"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(compressCmd)
}

var compressCmd = &cobra.Command{
	Use:   "compress [filename]",
	Short: "Print the version of archiver",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(args[0])

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
	},
}
