package cmd

import (
	"github.com/d1manpro/mustag/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mustag",
	Short: "ID3v2 tags editor",
	Args:  cobra.ArbitraryArgs,
	Run:   runRootCmd,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		ui.Error(err.Error())
	}
}

func runRootCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Help()
		return
	}

	runGetCmd(cmd, args)
}
