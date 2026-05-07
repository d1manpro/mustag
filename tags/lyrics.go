package tags

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bogem/id3v2/v2"
)

var ManyLyricsErr = errors.New("file have many lyrics flags")

func GetLyrics(tag *id3v2.Tag) (string, error) {
	if frames := tag.GetFrames("USLT"); len(frames) > 0 {
		if len(frames) == 1 {
			if lyr, ok := frames[0].(id3v2.UnsynchronisedLyricsFrame); ok {
				return lyr.Lyrics, nil
			} else {
				return "", fmt.Errorf("lyrics parse error")
			}
		} else {
			var lyrics []string
			for _, f := range frames {
				if lyr, ok := f.(id3v2.UnsynchronisedLyricsFrame); ok {
					lyrics = append(lyrics, fmt.Sprintf("=== [Description: %s, Lang: %s, Size: %d] ===\n%s", lyr.ContentDescriptor, lyr.Language, lyr.Size(), lyr.Lyrics))
				} else {
					return "", fmt.Errorf("lyrics parse error")
				}
			}
			return strings.Join(lyrics, "\n\n\n"), ManyLyricsErr
		}
	}
	return "", nil
}

func UpdateLyrics(tag *id3v2.Tag, lyrics string) error {
	tag.DeleteFrames("USLT")
	tag.AddUnsynchronisedLyricsFrame(id3v2.UnsynchronisedLyricsFrame{
		Encoding:          tag.DefaultEncoding(),
		Language:          "eng",
		ContentDescriptor: "",
		Lyrics:            lyrics,
	})
	return nil
}

func DeleteLyrics(tag *id3v2.Tag) {
	tag.DeleteFrames("USLT")
}
