package inspector

import (
	"github.com/anacrolix/torrent/metainfo"
)

// MagnetInspector inspect magnet links
type MagnetInspector struct {
	Inspector
}

// Process will process input file path / url and return meta data
func (m MagnetInspector) Process(input string) (interface{}, error) {
	result, err := metainfo.ParseMagnetURI(input)

	if err != nil {
		return nil, err
	}

	return result, nil
}
