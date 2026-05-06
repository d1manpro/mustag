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
	Args: cobra.MinimumNArgs(1),
	RunE: runGetCmd,
}

func init() {
	getCmd.Flags().BoolP("full", "f", false, "Show all raw ID3 frames")
	rootCmd.AddCommand(getCmd)
}

func runGetCmd(cmd *cobra.Command, args []string) error {
	file := args[0]
	fields := args[1:]

	tag, err := tags.Open(file)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer tag.Close()

	full, _ := cmd.Flags().GetBool("full")
	if full {
		ui.Header(args[0])
		tags.PrintFull(tag)
		return nil
	}

	if len(fields) == 0 {
		basic := tags.GetAll(tag)
		ui.Header(args[0])
		for _, v := range basic {
			ui.KeyValue(v.Key, v.Value)
		}
		return nil
	}

	if len(fields) == 1 {
		val, ok := tags.GetField(tag, fields[0])
		if !ok {
			return fmt.Errorf("unknown field: %s", fields[0])
		}

		fmt.Println(val)
		return nil
	}

	for _, f := range fields {
		val, ok := tags.GetField(tag, f)
		if !ok {
			ui.Warn("unknown field: " + f)
			continue
		}

		ui.KeyValue(f, val)
	}
	return nil
}
