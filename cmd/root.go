package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "docker-increment-cli",
	Short:         "docker increment cli Tool",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		runCmd(),
		execContainerCmd(),
		stopContainerCmd(),
		removeContainerCmd(),
		removeImageCmd(),
	)
}
