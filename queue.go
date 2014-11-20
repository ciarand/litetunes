package litetunes

import (
	"time"

	"github.com/landaire/go-taglib"
)

// Queue is a queue of tracks to play
type Queue interface {
	Queue(*Track) error
	Dequeue() (*Track, error)
	All() *Track
	Count() int
}

// Track embodies information about an audio track
type Track struct {
	Path   string
	Artist string
	Album  string
	Track  int
	Length time.Duration
}

// NewTrackFromPath creates a new Track struct from the provided filesystem
// path
func NewTrackFromPath(p string) (*Track, error) {
	file, err := taglib.Read(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return &Track{
		Path:   p,
		Artist: file.Artist(),
		Album:  file.Album(),
		Track:  file.Track(),
		Length: file.Length(),
	}, nil
}
