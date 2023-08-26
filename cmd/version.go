package cmd

import (
	"archiver/version"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

const (
	AppName = "Archiver"
	Creator = "Arlet (github.com/Arlet2)"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of archiver",
	Long:  "Print the version of archiver",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		revision := version.Revision
		if revision == "generated" {
			revision = time.Now().Format("02012006150405")
		}

		fmt.Printf(AppName + " " + version.Name + "-" + revision + " by " + Creator)
	},
}
