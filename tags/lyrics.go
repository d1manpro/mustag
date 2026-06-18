package tags

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bogem/id3v2/v2"
)

var charReplacer = strings.NewReplacer(
	// Apostrophe variants -> standard apostrophe
	"’", "'", // RIGHT SINGLE QUOTATION MARK
	"‘", "'", // LEFT SINGLE QUOTATION MARK
	"‛", "'", // SINGLE HIGH-REVERSED-9 QUOTATION MARK
	"ʼ", "'", // MODIFIER LETTER APOSTROPHE
	"ʻ", "'", // MODIFIER LETTER TURNED COMMA
	"ˈ", "'", // MODIFIER LETTER VERTICAL LINE
	"´", "'", // ACUTE ACCENT
	"`", "'", // GRAVE ACCENT
	"′", "'", // PRIME

	// Double quote variants -> standard double quote
	"“", "\"", // LEFT DOUBLE QUOTATION MARK
	"”", "\"", // RIGHT DOUBLE QUOTATION MARK
	"„", "\"", // DOUBLE LOW-9 QUOTATION MARK
	"‟", "\"", // DOUBLE HIGH-REVERSED-9 QUOTATION MARK
	"«", "\"", // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
	"»", "\"", // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
	"〝", "\"", // REVERSED DOUBLE PRIME QUOTATION MARK
	"〞", "\"", // DOUBLE PRIME QUOTATION MARK
	"〟", "\"", // LOW DOUBLE PRIME QUOTATION MARK
	"＂", "\"", // FULLWIDTH QUOTATION MARK
	"″", "\"", // DOUBLE PRIME
	"‶", "\"", // REVERSED DOUBLE PRIME
	"❝", "\"", // HEAVY DOUBLE TURNED COMMA QUOTATION MARK ORNAMENT
	"❞", "\"", // HEAVY DOUBLE COMMA QUOTATION MARK ORNAMENT

	// Dash variants -> standard dash
	"—", "-", // EM DASH
	"–", "-", // EN DASH
	"―", "-", // HORIZONTAL BAR
	"−", "-", // MINUS SIGN

	// Space variants -> standard space
	" ", " ", // MEDIUM MATHEMATICAL SPACE
	"\u00A0", " ", // NO-BREAK SPACE
	"\u2000", " ",
	"\u2001", " ",
	"\u2002", " ",
	"\u2003", " ",
	"\u2004", " ",
	"\u2005", " ",
	"\u2006", " ",
	"\u2007", " ",
	"\u2008", " ",
	"\u2009", " ",
	"\u200A", " ",
	"\u202F", " ",
	"\u205F", " ",
	"\u3000", " ",

	// 3 dots -> 3 separated dots
	"…", "...",
)

func NormalizeLyrics(s string) string {
	return charReplacer.Replace(s)
}

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

// EncodingFor returns the encoding to use for s.
// If preferred cannot represent s (Latin-1 with non-ASCII runes),
// it falls back to UTF-8 and returns the first unsupported rune with its line/col (1-based).
func EncodingFor(preferred id3v2.Encoding, s string) (enc id3v2.Encoding, badRune rune, line, col int) {
	if preferred.Key != 0 {
		return preferred, 0, 0, 0
	}
	for ln, lineStr := range strings.Split(s, "\n") {
		for cl, r := range lineStr {
			if r > 0xFF {
				return id3v2.EncodingUTF8, r, ln + 1, cl + 1
			}
		}
	}
	return preferred, 0, 0, 0
}

func UpdateLyrics(tag *id3v2.Tag, lyrics string) (rune, int, int, error) {
	enc, badRune, line, col := EncodingFor(tag.DefaultEncoding(), lyrics)
	tag.DeleteFrames("USLT")
	tag.AddUnsynchronisedLyricsFrame(id3v2.UnsynchronisedLyricsFrame{
		Encoding:          enc,
		Language:          "eng",
		ContentDescriptor: "",
		Lyrics:            lyrics,
	})
	return badRune, line, col, nil
}
