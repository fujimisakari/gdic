package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "gdic",
	Short:         "go docker incremental cli",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		runCmd(),
		execCmd(),
		stopCmd(),
		rmCmd(),
		rmiCmd(),
	)
}
