package tags

import (
	"fmt"

	"github.com/bogem/id3v2/v2"
)

func DeleteAllFrames(tag *id3v2.Tag) {
	tag.DeleteAllFrames()
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

func DeleteLyrics(tag *id3v2.Tag) {
	tag.DeleteFrames("USLT")
}

func DeleteImages(tag *id3v2.Tag) {
	tag.DeleteFrames("APIC")
}
