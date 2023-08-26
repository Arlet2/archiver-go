package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "archiver",
	Short: "Archiver is a program for compressing or decompressing a data",
	Long: "Archiver is a simple program for compressing or decompressing files.\n" +
		"All compressed files will save as .eva archive file (Exclusive Valuable Archive).",
	Args: cobra.NoArgs,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
