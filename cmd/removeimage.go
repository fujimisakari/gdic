package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/docker-increment-cli/cli"
)

func removeImageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rmi",
		Short: "Remove Docker Image",
		Run:   removeImage,
	}
	return cmd
}

func removeImage(cmd *cobra.Command, args []string) {
	rmImageCLI := cli.NewRemoveImageCLI()
	cli.Run(rmImageCLI)
}
