package cmd

import "github.com/spf13/cobra"

func genPullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull the example package from the cloud and extract",
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {

	}

	return cmd
}
