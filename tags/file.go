package tags

import "github.com/bogem/id3v2/v2"

func Open(path string) (*id3v2.Tag, error) {
	return id3v2.Open(path, id3v2.Options{Parse: true})
}

func SaveTag(tag *id3v2.Tag) error {
	return tag.Save()
}
