package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/gdic/cli"
)

func stopCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop Docker Container",
		RunE:  stop,
	}
	return cmd
}

func stop(cmd *cobra.Command, args []string) error {
	stopCLI, err := cli.NewStopCLI()
	if err != nil {
		return err
	}
	return cli.Run(stopCLI)
}
