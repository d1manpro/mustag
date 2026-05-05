package cmd

import (
	"fmt"
	"os"

	"github.com/d1manpro/mustag/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mustag",
	Short: "ID3v2 tags editor",
	Run:   runRootCmd,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.Err("Error: "), err)
	}
}

func init() {
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
}

func runRootCmd(cmd *cobra.Command, args []string) {}
