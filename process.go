package main

import "time"

func Process() error {
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
		if filter(k) {
			notified = append(notified, k)
		}
	}

	return notifier.Notify("締め切りが近い課題です。", notified)
}

// 締め切りまで3日以内の課題をフィルター
func filter(k Kadai) bool {
	return k.deadline != nil && k.deadline.Before(time.Now().Add(72*time.Hour))
}
