package cmd

import (
	"github.com/spf13/cobra"
)

func ImagesCmd() *cobra.Command {
	command := &cobra.Command{
		Use:          "images",
		Short:        "plugin that helps in fetching images from the selected kubernetes kind.",
		SilenceUsage: true,
	}
	command.AddCommand(getImagesCmd())
	return command
}

func getImagesCmd() *cobra.Command {
	command := &cobra.Command{
		Use:          "get",
		Short:        "Gets images from the selected kubernetes kind.",
		SilenceUsage: true,
		RunE:         c.GetImages,
	}
	registerFlags(command)
	return command	
}