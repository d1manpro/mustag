package cmd

import (
	"fmt"

	"github.com/d1manpro/mustag/tags"
	"github.com/d1manpro/mustag/ui"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <file> [field...]",
	Short: "Show metadata of an audio file",
	Long: `Show ID3v2 metadata of the specified file.

If no fields are provided, all available tags are displayed.

You can optionally specify one or more fields to filter output.
Common fields include: artist, title, album, year, genre, track.

Output behavior:
- Single field → raw value (useful for scripts)
- Multiple fields → formatted output
- No fields → full formatted tag list`,
	Example: `  mustag get song.mp3
  mustag get song.mp3 title
  mustag get song.mp3 title artist album`,
	Run: runGetCmd,
}

func init() {
	getCmd.Flags().Bool("full", false, "Show all raw ID3 frames")
	rootCmd.AddCommand(getCmd)
}

func runGetCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		ui.Error("file is required")
		return
	}

	file := args[0]
	fields := args[1:]

	tag, err := tags.Open(file)
	if err != nil {
		ui.Error(err.Error())
		return
	}
	defer tag.Close()

	full, _ := cmd.Flags().GetBool("full")
	if full {
		tags.PrintFull(tag)
		return
	}

	if len(fields) == 0 {
		basic := tags.GetAll(tag)

		for _, v := range basic {
			ui.KeyValue(v.Key, v.Value)
		}
		return
	}

	if len(fields) == 1 {
		val, ok := tags.GetField(tag, fields[0])
		if !ok {
			ui.Error("unknown field: " + fields[0])
			return
		}

		fmt.Println(val)
		return
	}

	for _, f := range fields {
		val, ok := tags.GetField(tag, f)
		if !ok {
			ui.Warn("unknown field: " + f)
			continue
		}

		ui.KeyValue(f, val)
	}
}
