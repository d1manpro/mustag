package cmd

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/d1manpro/mustag/tags"
	"github.com/d1manpro/mustag/ui"
	"github.com/spf13/cobra"
)

var lyricsCmd = &cobra.Command{
	Use:     "lyrics <file>",
	Aliases: []string{"l"},
	Short:   "Edit embedded lyrics in audio file",
	Long: `Open embedded lyrics from audio file in external editor, allow modification or deletion, and write changes back to tag.

If editor returns unchanged content, no update is performed.
If content is empty after editing, lyrics tag is removed.`,
	Example: `  mustag lyrics song.mp3
  mustag lyrics -e nvim song.mp3
  EDITOR=vim mustag lyrics song.mp3`,
	Args: cobra.ExactArgs(1),
	RunE: runLyricsCmd,
}

func init() {
	lyricsCmd.Flags().StringP("editor", "e", "", "Text editor")

	rootCmd.AddCommand(lyricsCmd)
}

func runLyricsCmd(cmd *cobra.Command, args []string) error {
	tag, err := tags.Open(args[0])
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer tag.Close()

	lyrics, err := tags.GetLyrics(tag)
	if err != nil {
		return err
	}

	var editor string
	if cmd.Flags().Changed("editor") {
		editor, _ = cmd.Flags().GetString("editor")
	} else {
		editor = os.Getenv("EDITOR")
	}
	if editor == "" {
		editor = "vi"
	}

	newLyrics, changed, err := editInEditor(lyrics, editor)
	if err != nil {
		return fmt.Errorf("editor error: %w", err)
	}
	if !changed {
		ui.Info("Lyrics not changed")
		return nil
	}
	newLyrics = strings.TrimSpace(newLyrics)
	var status string
	if newLyrics == "" {
		tags.DeleteLyrics(tag)
		status = ("New lyrics is empty. Tag was deleted")
	} else {
		normalized := tags.NormalizeLyrics(newLyrics)
		if normalized != newLyrics {
			ui.Info("Invalid characters were auto-replaced")
			newLyrics = normalized
		}
		badRune, line, col, err := tags.UpdateLyrics(tag, newLyrics)
		if err != nil {
			return fmt.Errorf("update lyrics error: %w", err)
		}
		if badRune != 0 {
			ui.Warn(fmt.Sprintf("character %q (U+%04X) at line %d, col %d is not supported by the tag encoding, saving as UTF-8", badRune, badRune, line, col))
		}
		status = "Lyrics successfully updated"
	}

	if err = tags.SaveTag(tag); err != nil {
		return fmt.Errorf("save changes error: %w", err)
	}

	ui.Info(status)

	return nil
}

func editInEditor(initial string, editor string) (string, bool, error) {
	tmpFile, err := os.CreateTemp("", ".lyrics-*.txt")
	if err != nil {
		return "", false, err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(initial); err != nil {
		return "", false, err
	}
	tmpFile.Close()

	before := checksum(strings.TrimSpace(initial))

	cmdExec := exec.Command(editor, tmpFile.Name())
	cmdExec.Stdin = os.Stdin
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr

	if err := cmdExec.Run(); err != nil {
		return "", false, err
	}

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return "", false, err
	}

	after := checksum(strings.TrimSpace(string(data)))

	if before == after {
		return "", false, nil
	}

	if len(data) == 0 {
		return "", true, nil
	}

	return string(data), true, nil
}

func checksum(s string) string {
	h := sha1.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}
