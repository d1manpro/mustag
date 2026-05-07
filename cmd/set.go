package cmd

import (
	"fmt"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/d1manpro/mustag/tags"
	"github.com/d1manpro/mustag/ui"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set <file>",
	Short: "Update ID3v2 metadata tags in audio file",
	Long: `Set or update ID3v2 metadata tags in an audio file.

Supports standard tags like title, artist, album, genre, year, track number, disk number,
as well as lyrics, cover art, and custom frames by ID.`,
	Example: `  mustag set song.mp3 -t "Title" -a "Artist"
  mustag set song.mp3 -y 2024 -n 1
  mustag set song.mp3 --lyrics lyrics.txt --cover cover.jpg
  mustag set song.mp3 --custom TXXX:MyTag:value`,
	Args: cobra.ExactArgs(1),
	RunE: runSetCmd,
}

var custom []string

func init() {
	setCmd.Flags().StringP("title", "t", "", "Set Title tag (TIT2)")
	setCmd.Flags().StringP("artist", "a", "", "Set Artist tag (TPE1)")
	setCmd.Flags().StringP("album", "A", "", "Set Album tag (TALB)")
	setCmd.Flags().String("album-artist", "", "Set Album-Artist tag (TPE2)")
	setCmd.Flags().StringP("genre", "g", "", "Set Genre tag (TCON)")
	setCmd.Flags().StringP("lyrics", "l", "", "Load Lyrics from file (USLT)")
	setCmd.Flags().StringP("cover", "c", "", "Load Cover image (APIC)")

	setCmd.Flags().IntP("year", "y", 0, "Set Year tag (TYER)")
	setCmd.Flags().IntP("number", "n", 0, "Set Track Number tag (TRCK)")
	setCmd.Flags().IntP("disk", "d", 0, "Set Disk tag (TPOS)")

	setCmd.Flags().StringArrayVar(&custom, "custom", nil, "Set Custom tag by ID. Format: id:value")

	rootCmd.AddCommand(setCmd)
}

func runSetCmd(cmd *cobra.Command, args []string) error {
	if cmd.Flags().NFlag() == 0 {
		return fmt.Errorf("no specified tags to update")
	}

	tag, err := tags.Open(args[0])
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer tag.Close()

	if err = applyFlags(tag, cmd); err != nil {
		return fmt.Errorf("tags apply error: %w", err)
	}

	if err = tags.SaveTag(tag); err != nil {
		return fmt.Errorf("save changes error: %w", err)
	}

	ui.Info("Tags successfully updated")

	return nil
}

func applyFlags(tag *id3v2.Tag, cmd *cobra.Command) error {
	if cmd.Flags().Changed("title") {
		v, _ := cmd.Flags().GetString("title")
		if e := tags.SetStringFrame(tag, "title", v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("title") {
		v, _ := cmd.Flags().GetString("title")
		if e := tags.SetStringFrame(tag, "title", v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("artist") {
		v, _ := cmd.Flags().GetString("artist")
		if e := tags.SetStringFrame(tag, "artist", v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("album") {
		v, _ := cmd.Flags().GetString("album")
		if e := tags.SetStringFrame(tag, "album", v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("album-artist") {
		v, _ := cmd.Flags().GetString("album-artist")
		if e := tags.SetStringFrame(tag, "album-artist", v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("genre") {
		v, _ := cmd.Flags().GetString("genre")
		if e := tags.SetStringFrame(tag, "genre", v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("year") {
		v, _ := cmd.Flags().GetInt("year")
		if e := tags.SetIntFrame(tag, "year", v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("number") {
		v, _ := cmd.Flags().GetInt("number")
		if e := tags.SetIntFrame(tag, "number", v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("disk") {
		v, _ := cmd.Flags().GetInt("disk")
		if e := tags.SetIntFrame(tag, "disk", v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("lyrics") {
		v, _ := cmd.Flags().GetString("lyrics")
		if e := tags.SetLyrics(tag, v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("cover") {
		v, _ := cmd.Flags().GetString("cover")
		if e := tags.SetCover(tag, v); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("custom") {
		arr, _ := cmd.Flags().GetStringArray("custom")

		for _, item := range arr {
			parts := strings.SplitN(item, ":", 2)
			if len(parts) != 2 {
				return fmt.Errorf("invalid custom format: %s", item)
			}

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			if key == "" {
				return fmt.Errorf("empty custom key: %s", item)
			}

			tag.AddTextFrame(key, tag.DefaultEncoding(), value)
		}
	}

	return nil
}
