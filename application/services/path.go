package services

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

// PathService get workspace path
type PathService interface {
	WorkspaceDir() string
	YoutubeDLPath() string
}

// NewPathService creates PathService
func NewPathService() PathService {
	return &pathService{}
}

type pathService struct {
}

func (p pathService) WorkspaceDir() string {
	path, err := homedir.Dir()

	if err != nil {
		panic(err)
	}
	defaultWorkspace := filepath.Join(path, ".loda")
	if _, err := os.Stat(defaultWorkspace); os.IsNotExist(err) {
		os.Mkdir(defaultWorkspace, os.ModePerm)
	}

	return defaultWorkspace
}

func (p pathService) YoutubeDLPath() string {
	binName := "youtube-dl"
	if runtime.GOOS == "windows" {
		binName = "youtube-dl.exe"
	}
	return filepath.Join(p.WorkspaceDir(), binName)
}
