package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/gdic/cli"
)

func execCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec",
		Short: "Exec Docker Container",
		RunE:  exec,
	}
	return cmd
}

func exec(cmd *cobra.Command, args []string) error {
	execCLI, err := cli.NewExecCLI()
	if err != nil {
		return err
	}
	return cli.Run(execCLI)
}
