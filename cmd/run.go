package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/gdic/cli"
)

func runCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run Docker Container",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	runCLI, err := cli.NewRunCLI()
	if err != nil {
		return err
	}
	return cli.Run(runCLI)
}
