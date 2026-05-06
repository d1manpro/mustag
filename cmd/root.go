package cmd

import (
	"github.com/d1manpro/mustag/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "mustag <file>",
	Short:         "ID3v2 tags editor",
	Args:          cobra.ArbitraryArgs,
	RunE:          runRootCmd,
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		ui.Error(err.Error())
	}
}

func runRootCmd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Help()
	}
	return runGetCmd(cmd, args)
}
