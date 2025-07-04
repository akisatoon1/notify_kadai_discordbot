package main

import (
	"notify_kadai_discordbot/internal/ds"
	"time"
)

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
		if filter(k, c.ExcludedCourses) {
			notified = append(notified, k)
		}
	}

	return notifier.Notify("締め切りが近い課題です。", notified)
}

func filter(k Kadai, exc ds.Set) bool {
	// 締め切りまで3日以内の課題
	deadlineCond := k.deadline != nil && k.deadline.Before(time.Now().Add(72*time.Hour))

	// 除外コースに含まれない課題
	excludedCond := !exc.Contains(k.course)

	return deadlineCond && excludedCond
}
