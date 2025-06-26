package main

import (
	"testing"
	"time"
)

func TestNotify(t *testing.T) {
	n := NewNotifier(TOKEN, CHANNEL_ID)

	tt := time.Date(2025, 6, 10, 13, 20, 0, 0, time.UTC)
	kadais := []Kadai{
		NewKadai("レポート", "テスト課題", "数理基礎1", &tt, &tt),
		NewKadai("レポート", "テストnil課題", "数理基礎1", nil, nil),
	}

	err := n.Notify("this is prefix", kadais)
	if err != nil {
		t.Errorf("Notify failed: %v", err)
	}

	err = n.Notify("this is prefix", []Kadai{})
	if err != nil {
		t.Errorf("Notify with empty list failed: %v", err)
	}
}
