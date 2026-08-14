package notify

import (
	"testing"
	"time"
)

func TestCloneReadAtPointerIndependence(t *testing.T) {
	s := New()
	now := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	s.Create(CreateInput{ID: "T1", Recipient: "u", Content: "c"}, now)
	s.MarkSent("T1", now.Add(time.Hour))

	readTime := now.Add(2 * time.Hour)
	returned, err := s.MarkRead("T1", readTime)
	if err != nil {
		t.Fatal(err)
	}

	// 修改返回值的 ReadAt 指针指向的值
	*returned.ReadAt = time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)

	// store 内部不应被影响
	got, _ := s.Get("T1")
	if !got.ReadAt.Equal(readTime) {
		t.Errorf("clone shares ReadAt pointer with store: got %v, want %v", got.ReadAt, readTime)
	}
}
