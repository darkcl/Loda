package services_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/darkcl/loda/application/services"
	"github.com/stretchr/testify/mock"
)

type MockPathService struct {
	services.PathService
	mock.Mock
}

func (m *MockPathService) WorkspaceDir() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockPathService) YoutubeDLPath() string {
	args := m.Called()
	return args.String(0)
}

func TestMatchURL(t *testing.T) {
	mockPathService := new(MockPathService)
	mockPathService.On("YoutubeDLPath").Return("../")
	sut := services.NewMatcherService(mockPathService)
	result, err := sut.Match("http://example.com/text.txt", "./")
	assert.Equal(t, result.Identifier(), "url")
	assert.Nil(t, err)
}

func TestMatchYoutube(t *testing.T) {
	mockPathService := new(MockPathService)

	if runtime.GOOS == "windows" {
		mockPathService.On("YoutubeDLPath").Return("../../bin/youtube-dl.exe")
	} else {
		mockPathService.On("YoutubeDLPath").Return("../../bin/youtube-dl")
	}

	sut := services.NewMatcherService(mockPathService)
	result, err := sut.Match("https://www.youtube.com/watch?v=cOeLz87i0XE", "./")
	assert.Equal(t, result.Identifier(), "youtube-dl")
	assert.Nil(t, err)
}

func TestMagnetLink(t *testing.T) {
	mockPathService := new(MockPathService)
	mockPathService.On("YoutubeDLPath").Return("../")

	sut := services.NewMatcherService(mockPathService)
	result, err := sut.Match(`magnet:?xt=urn:btih:546cf15f724d19c4319cc17b179d7e035f89c1f4&dn=ubuntu-14.04.2-desktop-amd64.iso&tr=http%%3A%%2F%%2Ftorrent.ubuntu.com%%3A6969%%2Fannounce&tr=http%%3A%%2F%%2Fipv6.torrent.ubuntu.com%%3A6969%%2Fannounce`, "./")
	id := result.Identifier()
	assert.Equal(t, id, "magnet")
	assert.Nil(t, err)
}

func TestTorrent(t *testing.T) {
	mockPathService := new(MockPathService)
	mockPathService.On("YoutubeDLPath").Return("../")

	sut := services.NewMatcherService(mockPathService)
	result, err := sut.Match("./testdata/test.torrent", "./")
	assert.Equal(t, result.Identifier(), "torrent")
	assert.Nil(t, err)
}
