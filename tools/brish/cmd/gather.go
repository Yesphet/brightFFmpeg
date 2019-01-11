package cmd

import "github.com/spf13/cobra"

func genGatherCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gather",
		Short: "Gather all examples' information and write to List.md/Tags.md",
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {

	}

	return cmd
}
