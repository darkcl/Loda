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
