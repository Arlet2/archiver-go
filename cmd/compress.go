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
	Use:   "compress [filename] [output filename]",
	Short: "Compress file in .eva archive file",
	Long:  "Compress file in .eva archive file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(args[0])

		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}

		lines, err := readers.ReadFile(file)
		file.Close()

		freqDict := encoding.CreateFreqDictFromLines(lines)

		codingDict := encoding.CreateCodingDict(encoding.CreateCodingTree(freqDict))

		lines, err = compressor.Compress(lines, codingDict)

		outputFile, err := os.Create(args[1] + ".eva")

		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
		defer outputFile.Close()

		err = writers.WriteFile(outputFile, lines)

		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
	},
}
