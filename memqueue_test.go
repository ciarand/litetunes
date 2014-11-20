package litetunes_test

import (
	"reflect"
	"testing"

	"github.com/ciarand/litetunes"
)

func TestMemoryQueueCount(t *testing.T) {
	q := litetunes.NewMemoryQueue()

	if q.Count() != 0 {
		t.Error("Count should be 0 for an empty queue")
		return
	}

	track := &litetunes.Track{}

	iters := 1000

	for i := 0; i < iters; i++ {
		q.Queue(track)

		if q.Count() != i+1 {
			t.Errorf("Count should be %d after queueing %d items", i+1, i+1)
			return
		}
	}

	for i := 0; i < iters; i++ {
		dq, err := q.Dequeue()

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			return
		}

		if !reflect.DeepEqual(dq, track) {
			t.Errorf("Dequeue should return the same pointer passed in (expected %#v, got %#v)", track, dq)
			return
		}

		if q.Count() != iters-i-1 {
			t.Errorf("Expected count of %d after dequeuing %d items, got %d", iters-i-1, i+1, q.Count())
			return
		}
	}
}
