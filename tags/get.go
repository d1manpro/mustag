package tags

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/d1manpro/mustag/ui"
)

type Field struct {
	Key   string
	Value string
}

func GetField(tag *id3v2.Tag, field string) (string, bool) {
	switch field {
	case "title":
		return tag.Title(), true
	case "artist":
		return tag.Artist(), true
	case "album":
		return tag.Album(), true
	case "year":
		return tag.Year(), true
	case "genre":
		return tag.Genre(), true
	case "lyrics":
		lyr, err := GetLyrics(tag)
		if err != nil && !errors.Is(err, ManyLyricsErr) {
			ui.Error(err.Error())
			return "", false
		}
		return lyr, true
	default:
		return "", false
	}
}

func GetAll(tag *id3v2.Tag) []Field {
	var out []Field

	add := func(k, v string) {
		if v != "" {
			out = append(out, Field{k, v})
		}
	}

	add("title", tag.Title())
	add("artist", tag.Artist())
	add("album", tag.Album())
	add("year", tag.Year())

	if tf := tag.GetTextFrame("TRCK"); tf.Text != "" {
		add("track", tf.Text)
	}

	if frames := tag.GetFrames("APIC"); len(frames) > 0 {
		var picData []string
		for _, f := range frames {
			if pic, ok := f.(id3v2.PictureFrame); ok {
				picData = append(picData, fmt.Sprintf("image (%s, %d bytes)", pic.MimeType, len(pic.Picture)))
			}
		}
		add("cover", fmt.Sprintf("%d image(s): ", len(frames))+strings.Join(picData, "; "))
	}

	if frames := tag.GetFrames("USLT"); len(frames) > 0 {
		totalLines := 0

		for _, f := range frames {
			if lyr, ok := f.(id3v2.UnsynchronisedLyricsFrame); ok {
				totalLines += strings.Count(lyr.Lyrics, "\n") + 1
			}
		}

		add("lyrics", fmt.Sprintf("%d item(s), %d lines", len(frames), totalLines))
	}

	return out
}
