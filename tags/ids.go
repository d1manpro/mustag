package tags

var frameNameMap = map[string]string{
	// basic metadata
	"TIT2": "title",
	"TPE1": "artist",
	"TALB": "album",
	"TPE2": "album_artist",
	"TPE3": "conductor",
	"TPE4": "remixed_by",
	"TCOM": "composer",
	"TCON": "genre",
	"TRCK": "track",
	"TPOS": "disc",

	// dates
	"TDRC": "recording_date",
	"TDOR": "original_release_date",
	"TYER": "year",

	// misc
	"TBPM": "bpm",
	"TLAN": "language",
	"TSRC": "isrc",
	"TPUB": "publisher",
	"TCOP": "copyright",

	// text containers
	"TXXX": "user_text",
	"COMM": "comment",
	"USLT": "lyrics",

	// media
	"APIC": "cover",

	// technical
	"POPM": "rating",
	"UFID": "file_id",
}

func frameLabel(id string) string {
	if v, ok := frameNameMap[id]; ok {
		return v
	}
	return "id:" + id
}
