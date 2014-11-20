package litetunes

import (
	"errors"
	"fmt"
)

// MemoryQueue is an in-memory implementation of the Queue interface
type MemoryQueue struct {
	tracks []*Track
}

// NewMemoryQueue constructs a new MemoryQueue to use
func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{tracks: []*Track{}}
}

// Queue adds a new track to the queue
func (m *MemoryQueue) Queue(t *Track) error {
	if curlen := len(m.tracks); cap(m.tracks) == curlen {
		bigger := make([]*Track, curlen, 2*curlen+1)
		copy(bigger, m.tracks)
		m.tracks = bigger
	}

	m.tracks = append(m.tracks, t)

	return nil
}

// Dequeue removes the track at the head of the queue
func (m *MemoryQueue) Dequeue() (*Track, error) {
	c := m.Count()

	if c < 1 {
		return nil, errors.New("No track to dequeue")
	}

	t := m.tracks[0]
	if t == nil {
		for k, v := range m.tracks {
			fmt.Printf("%d = %#v\n", k, v)
		}
	}

	m.tracks = m.tracks[1:]

	return t, nil
}

// Count returns the number of tracks in the queue
func (m *MemoryQueue) Count() int {
	return len(m.tracks)
}
