package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/gdic/cli"
)

func execContainerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ex",
		Short: "Exec Docker Container",
		RunE:  execContainer,
	}
	return cmd
}

func execContainer(cmd *cobra.Command, args []string) error {
	execContainerCLI, err := cli.NewExecContainerCLI()
	if err != nil {
		return err
	}
	return cli.Run(execContainerCLI)
}
