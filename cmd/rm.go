package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/gdic/cli"
)

func rmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove Docker Container",
		RunE:  rm,
	}
	return cmd
}

func rm(cmd *cobra.Command, args []string) error {
	rmCLI, err := cli.NewRmCLI()
	if err != nil {
		return err
	}
	return cli.Run(rmCLI)
}
