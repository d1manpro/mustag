package tags

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/bogem/id3v2/v2"
)

func SetStringFrame(tag *id3v2.Tag, k, v string) error {
	switch k {
	case "title":
		tag.SetTitle(v)
	case "artist":
		tag.SetArtist(v)
	case "album":
		tag.SetAlbum(v)
	case "album-artist":
		tag.AddTextFrame("TPE2", tag.DefaultEncoding(), v)
	case "genre":
		tag.SetGenre(v)
	default:
		return fmt.Errorf("unknown string tag: %s", k)
	}
	return nil
}

func SetIntFrame(tag *id3v2.Tag, k string, v int) error {
	s := strconv.Itoa(v)
	switch k {
	case "year":
		tag.AddTextFrame("TYER", tag.DefaultEncoding(), s)
	case "number":
		tag.AddTextFrame("TRCK", tag.DefaultEncoding(), s)
	case "disk":
		tag.AddTextFrame("TPOS", tag.DefaultEncoding(), s)
	default:
		return fmt.Errorf("unknown int tag: %s", k)
	}

	return nil
}

func SetFrameByID(tag *id3v2.Tag, id, v string) error {
	if id == "" {
		return fmt.Errorf("empty frame id")
	}
	if v == "" {
		return fmt.Errorf("empty frame value")
	}

	tag.AddTextFrame(id, tag.DefaultEncoding(), v)
	return nil
}

func DeleteFrame(tag *id3v2.Tag, k string) error {
	switch k {
	case "title":
		tag.DeleteFrames("TIT2")
	case "artist":
		tag.DeleteFrames("TIT2")
	case "album":
		tag.DeleteFrames("TIT2")
	case "album-artist":
		tag.DeleteFrames("TIT2")
	case "genre":
		tag.DeleteFrames("TIT2")
	case "year":
		tag.DeleteFrames("TYER")
	case "number":
		tag.DeleteFrames("TRCK")
	case "disk":
		tag.DeleteFrames("TPOS")
	default:
		return fmt.Errorf("unknown string tag: %s", k)
	}
	return nil
}

func SetLyrics(tag *id3v2.Tag, lyricsPath string) error {
	data, err := os.ReadFile(lyricsPath)
	if err != nil {
		return fmt.Errorf("read cover: %w", err)
	}

	tag.AddUnsynchronisedLyricsFrame(id3v2.UnsynchronisedLyricsFrame{
		Encoding:          tag.DefaultEncoding(),
		Language:          "eng",
		ContentDescriptor: "",
		Lyrics:            string(data),
	})
	return nil
}

func SetCover(tag *id3v2.Tag, coverPath string) error {
	data, err := os.ReadFile(coverPath)
	if err != nil {
		return fmt.Errorf("read cover: %w", err)
	}

	mime := http.DetectContentType(data)

	tag.AddAttachedPicture(id3v2.PictureFrame{
		Encoding:    tag.DefaultEncoding(),
		MimeType:    mime,
		PictureType: id3v2.PTFrontCover,
		Description: "Cover",
		Picture:     data,
	})

	return nil
}
