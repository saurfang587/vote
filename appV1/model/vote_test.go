package model

import "testing"

func TestDoVote(t *testing.T) {
	New()
	DoVote(1, 1, []int64{1, 2, 3})
	Close()
}
