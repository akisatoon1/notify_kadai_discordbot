package main

import (
	"testing"
	"time"
)

func TestNotify(t *testing.T) {
	n := NewNotifier(TOKEN, CHANNEL_ID)

	kadais := []Kadai{
		NewKadai("レポート", "第8回課題", "数理基礎1", time.Date(2025, 6, 10, 13, 20, 0, 0, time.UTC), time.Date(2025, 6, 24, 13, 20, 0, 0, time.UTC)),
		NewKadai("レポート", "第9回課題", "数理基礎1", time.Date(2025, 6, 11, 13, 20, 0, 0, time.UTC), time.Date(2025, 6, 25, 13, 20, 0, 0, time.UTC)),
	}

	err := n.Notify(kadais)
	if err != nil {
		t.Errorf("Notify failed: %v", err)
	}

	err = n.Notify([]Kadai{})
	if err != nil {
		t.Errorf("Notify with empty list failed: %v", err)
	}
}
