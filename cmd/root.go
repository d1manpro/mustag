package cmd

import (
	"github.com/d1manpro/mustag/system"
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

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Show mustag version")
}

func runRootCmd(cmd *cobra.Command, args []string) error {
	if cmd.Flags().Changed("version") {
		ui.Info("mustag version " + system.Version)
		return nil
	}
	if len(args) == 0 {
		return cmd.Help()
	}
	return runGetCmd(cmd, args)
}
