package services

import (
	"errors"
	"time"

	"github.com/darkcl/loda/lib/downloader"
	"github.com/darkcl/loda/lib/matcher"
	"github.com/segmentio/ksuid"
)

// MatcherService takes a input (string / file path) and determain which download we use
type MatcherService interface {
	Match(input string, destination string) (downloader.Downloader, error)
}

func createMatcherTree(pathService PathService) matcherTree {
	urlMatcher := matcher.URLMatcher{}
	ytdlMatcher := matcher.NewYoutubeDLMatcher(pathService.YoutubeDLPath())
	magnetMatcher := matcher.MagnetMatcher{}
	torrentMatcher := matcher.TorrentMatcher{}

	rootNode := &matcherNode{
		Matcher: &rootMatcher{},
		Next:    []*matcherNode{},
	}

	urlNode := rootNode.Add(urlMatcher)
	urlNode.Add(ytdlMatcher)

	rootNode.Add(magnetMatcher)
	rootNode.Add(torrentMatcher)

	return matcherTree{
		Root: rootNode,
	}
}

// NewMatcherService creates MatcherService
func NewMatcherService(pathService PathService) MatcherService {
	return &matcherService{
		Tree:        createMatcherTree(pathService),
		pathService: pathService,
	}
}

type rootMatcher struct {
	matcher.Matcher
}

func (r rootMatcher) Process(input string) bool {
	return true
}

func (r rootMatcher) Identifier() string {
	return ""
}

type matcherTree struct {
	Root *matcherNode
}

type matcherNode struct {
	Matcher matcher.Matcher
	Next    []*matcherNode
	Parent  *matcherNode
}

func (t *matcherNode) Add(matcher matcher.Matcher) *matcherNode {
	nextNode := &matcherNode{
		Matcher: matcher,
		Next:    []*matcherNode{},
		Parent:  t,
	}
	t.Next = append(t.Next, nextNode)
	return nextNode
}

func (t *matcherNode) Match(input string) string {
	if t.Matcher.Process(input) {
		if len(t.Next) != 0 {
			for _, node := range t.Next {
				return node.Match(input)
			}
		} else {
			return t.Matcher.Identifier()
		}
	} else {
		if t.Parent != nil {
			return t.Parent.Matcher.Identifier()
		}
		return ""
	}
	return ""
}

type matcherService struct {
	MatcherService
	Tree        matcherTree
	pathService PathService
}

func (m matcherService) MatchWithTree(input string) string {
	return m.Tree.Root.Match(input)
}

func (m matcherService) Match(input string, destination string) (downloader.Downloader, error) {
	downloaderID := m.MatchWithTree(input)
	switch downloaderID {
	case "url":
		interval := 1000 * time.Millisecond
		label := ksuid.New().String()

		loader := downloader.NewURLDownloader(downloader.URLDownloaderParams{
			URL:              input,
			Label:            label,
			Destination:      destination,
			NumOfConnections: 1,
			IsResumable:      true,
			ReportInterval:   interval,
		})
		return loader, nil
	case "youtube-dl":
		interval := 1000 * time.Millisecond

		loader := downloader.NewYoutubeDLDownloader(downloader.YoutubeDLDownloaderParams{
			URL:            input,
			Destination:    destination,
			ReportInterval: interval,
			BinaryPath:     m.pathService.YoutubeDLPath(),
		})
		return loader, nil
	default:
		return nil, errors.New("Downloader not found")
	}
}
