package main

import (
	"log"
	"time"
)

func main() {
	if err := process(); err != nil {
		log.Print(err)
	}
}

func process() error {
	getter := NewKadaiGetter()

	notifier := NewNotifier()

	kadais, err := getter.GetAll()
	if err != nil {
		return err
	}

	var notified []Kadai
	for _, k := range kadais {
		// 締め切りまで2日以内の課題をフィルター
		if k.deadline.Before(time.Now().Add(48 * time.Hour)) {
			notified = append(notified, k)
		}
	}

	return notifier.Notify(notified)
}
