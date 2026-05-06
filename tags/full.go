package tags

import (
	"fmt"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/d1manpro/mustag/ui"
)

const maxLyricsLen = 80

func PrintFull(tag *id3v2.Tag) {
	frames := tag.AllFrames()

	for id, list := range frames {
		for _, f := range list {
			printFrame(id, f)
		}
	}
}

func printFrame(id string, f id3v2.Framer) {
	id = frameLabel(id)
	switch v := f.(type) {
	case id3v2.TextFrame:
		ui.KeyValue(id, v.Text)
	case id3v2.CommentFrame:
		ui.KeyValue(id, v.Text)
	case id3v2.UnsynchronisedLyricsFrame:
		ui.KeyValue(id, trim(v.Lyrics))
	case id3v2.PictureFrame:
		ui.KeyValue(id, fmt.Sprintf("image (%s, %d bytes)", v.MimeType, len(v.Picture)))
	case id3v2.UserDefinedTextFrame:
		ui.KeyValue(id, fmt.Sprintf("%s=%s", v.Description, v.Value))
	default:
		ui.KeyValue(id, fmt.Sprintf("<%T>", f))
	}
}

func trim(s string) string {
	if len(s) > maxLyricsLen {
		return s[:maxLyricsLen] + "\n... " + fmt.Sprintf("%d lines more", 1+strings.Count(s[maxLyricsLen:], "\n"))
	}
	return s
}
