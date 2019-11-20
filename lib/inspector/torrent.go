package inspector

import (
	"github.com/anacrolix/torrent/metainfo"
)

// TorrentInspector inspects torrent file
type TorrentInspector struct {
	Inspector
}

// Process will process input file path / url and return meta data
func (t TorrentInspector) Process(input string) (interface{}, error) {
	result, err := metainfo.LoadFromFile(input)

	if err != nil {
		return nil, err
	}

	return result, nil
}
