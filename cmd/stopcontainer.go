package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/gdic/cli"
)

func stopContainerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop Docker Container",
		RunE:  stopContainer,
	}
	return cmd
}

func stopContainer(cmd *cobra.Command, args []string) error {
	stopContainerCLI, err := cli.NewStopContainerCLI()
	if err != nil {
		return err
	}
	return cli.Run(stopContainerCLI)
}
