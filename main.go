package main

import (
	"log"
	"time"
)

var TOKEN string
var CHANNEL_ID string
var STUDENT_ID string
var PASSWORD string

func main() {
	if err := process(); err != nil {
		log.Print(err)
	}
}

func process() error {
	c, err := LoadConfig()
	if err != nil {
		return err
	}

	getter := NewKadaiGetter(c.StudentID, c.Password)

	notifier := NewNotifier(c.DiscordToken, c.ChannelID)

	kadais, err := getter.GetAll()
	if err != nil {
		return err
	}

	var notified []Kadai
	for _, k := range kadais {
		// 締め切りまで3日以内の課題をフィルター
		if k.deadline != nil && k.deadline.Before(time.Now().Add(72*time.Hour)) {
			notified = append(notified, k)
		}
	}

	return notifier.Notify("締め切りが近い課題です。", notified)
}
