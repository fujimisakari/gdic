package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/gdic/cli"
)

func removeImageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rmi",
		Short: "Remove Docker Image",
		RunE:  removeImage,
	}
	return cmd
}

func removeImage(cmd *cobra.Command, args []string) error {
	rmImageCLI, err := cli.NewRemoveImageCLI()
	if err != nil {
		return err
	}
	return cli.Run(rmImageCLI)
}
