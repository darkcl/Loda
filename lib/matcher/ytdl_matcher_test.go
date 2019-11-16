package matcher_test

import (
	"errors"
	"testing"

	"github.com/darkcl/loda/lib/inspector"

	"github.com/darkcl/loda/lib/matcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockYoutubeDLInspector struct {
	inspector.Inspector
	mock.Mock
}

// Process is a mock process of MockYoutubeDLInspector
func (m *MockYoutubeDLInspector) Process(input string) (interface{}, error) {
	args := m.Called(input)
	return args.Get(0), args.Error(1)
}

func TestYoutubeDLSupportedSite(t *testing.T) {
	mockInspector := new(MockYoutubeDLInspector)
	mockInspector.On("Process", mock.Anything).Return("meta data", nil)

	sut := matcher.YoutubeDLMatcher{
		YtdlInspector: mockInspector,
	}

	result := sut.Process("youtube url")
	assert.True(t, result, "should match youtube video")
}

func TestYoutubeDLNotSupportedSite(t *testing.T) {
	mockInspector := new(MockYoutubeDLInspector)
	mockInspector.On("Process", mock.Anything).Return(nil, errors.New("Some Error"))

	sut := matcher.YoutubeDLMatcher{
		YtdlInspector: mockInspector,
	}

	result := sut.Process("non youtube url")
	assert.False(t, result, "should not match youtube video")
}
