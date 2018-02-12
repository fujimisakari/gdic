package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/gdic/cli"
)

func removeContainerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rmc",
		Short: "Remove Docker Container",
		Run:   removeContainer,
	}
	return cmd
}

func removeContainer(cmd *cobra.Command, args []string) {
	rmContainerCLI := cli.NewRemoveContainerCLI()
	cli.Run(rmContainerCLI)
}
