package inspector

// Inspector will inspect on a file (torrent) or a url (youtube, facebook, etc.) for meta data
type Inspector interface {

	// Process will process input file path / url and return meta data
	Process(input string) (interface{}, error)
}
