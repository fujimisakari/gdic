package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/gdic/cli"
)

func rmiCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rmi",
		Short: "Remove Docker Image",
		RunE:  rmi,
	}
	return cmd
}

func rmi(cmd *cobra.Command, args []string) error {
	rmiCLI, err := cli.NewRmiCLI()
	if err != nil {
		return err
	}
	return cli.Run(rmiCLI)
}
