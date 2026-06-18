package cmd

import (
	"fmt"

	"github.com/bogem/id3v2/v2"
	"github.com/d1manpro/mustag/tags"
	"github.com/d1manpro/mustag/ui"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove <file>",
	Aliases: []string{"rm"},
	Short:   "Remove ID3v2 metadata tags in audio file",
	Long: `Remove ID3v2 metadata tags in an audio file.

Supports standard tags like title, artist, album, genre, year, track number, disk number,
as well as lyrics, cover art, and custom frames by ID.`,
	Example: `  mustag remove song.mp3 --title -g
  mustag rm song.mp3 --album-artist --lyrics`,
	Args: cobra.ExactArgs(1),
	RunE: runRemoveCmd,
}

func init() {
	removeCmd.Flags().BoolP("title", "t", false, "Remove Title tag (TIT2)")
	removeCmd.Flags().BoolP("artist", "a", false, "Remove Artist tag (TPE1)")
	removeCmd.Flags().BoolP("album", "A", false, "Remove Album tag (TALB)")
	removeCmd.Flags().Bool("album-artist", false, "Remove Album-Artist tag (TPE2)")
	removeCmd.Flags().BoolP("genre", "g", false, "Remove Genre tag (TCON)")
	removeCmd.Flags().BoolP("lyrics", "l", false, "Remove Lyrics (USLT)")
	removeCmd.Flags().BoolP("cover", "c", false, "Remove Cover image (APIC)")

	removeCmd.Flags().BoolP("year", "y", false, "Remove Year tag (TYER)")
	removeCmd.Flags().BoolP("number", "n", false, "Remove Track Number tag (TRCK)")
	removeCmd.Flags().BoolP("disk", "d", false, "Remove Disk tag (TPOS)")

	removeCmd.Flags().StringArrayVar(&custom, "custom", nil, "Remove Custom tag by ID")

	removeCmd.Flags().Bool("all", false, "Remove All tags")

	rootCmd.AddCommand(removeCmd)
}

func runRemoveCmd(cmd *cobra.Command, args []string) error {
	if cmd.Flags().NFlag() == 0 {
		return fmt.Errorf("no specified tags to delete")
	}

	tag, err := tags.Open(args[0])
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer tag.Close()

	if err = deleteFlags(tag, cmd); err != nil {
		return fmt.Errorf("tags apply error: %w", err)
	}

	if err = tags.SaveTag(tag); err != nil {
		return fmt.Errorf("save changes error: %w", err)
	}

	ui.Info("Tags successfully deleted")

	return nil
}

func deleteFlags(tag *id3v2.Tag, cmd *cobra.Command) error {
	if cmd.Flags().Changed("all") {
		tags.DeleteAllFrames(tag)
	}

	if cmd.Flags().Changed("title") {
		if e := tags.DeleteFrame(tag, "title"); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("artist") {
		if e := tags.DeleteFrame(tag, "artist"); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("album") {
		if e := tags.DeleteFrame(tag, "album"); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("album-artist") {
		if e := tags.DeleteFrame(tag, "album-artist"); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("genre") {
		if e := tags.DeleteFrame(tag, "genre"); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("year") {
		if e := tags.DeleteFrame(tag, "year"); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("number") {
		if e := tags.DeleteFrame(tag, "number"); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("disk") {
		if e := tags.DeleteFrame(tag, "disk"); e != nil {
			return e
		}
	}

	if cmd.Flags().Changed("lyrics") {
		tags.DeleteLyrics(tag)
	}

	if cmd.Flags().Changed("cover") {
		tags.DeleteImages(tag)
	}

	if cmd.Flags().Changed("custom") {
		arr, _ := cmd.Flags().GetStringArray("custom")

		for _, item := range arr {
			tag.DeleteFrames(item)
		}
	}

	return nil
}
