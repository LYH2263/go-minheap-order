
package pq

import (
	"container/heap"
	"testing"
)

func TestPopAscending(t *testing.T) {
	h := &IntHeap{5, 1, 3}
	heap.Init(h)
	var out []int
	for h.Len() > 0 {
		out = append(out, heap.Pop(h).(int))
	}
	want := []int{1, 3, 5}
	for i := range want {
		if out[i] != want[i] {
			t.Fatalf("got %v want %v", out, want)
		}
	}
}
