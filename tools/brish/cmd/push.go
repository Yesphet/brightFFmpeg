package cmd

import (
	"github.com/spf13/cobra"
)

func genPushCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "push",
		Short: "Compress the example asserts and push to the cloud",
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {

	}

	return cmd
}
