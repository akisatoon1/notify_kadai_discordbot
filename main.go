package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
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

func init() {
	_ = godotenv.Load()
	TOKEN = os.Getenv("DISCORD_TOKEN")
	CHANNEL_ID = os.Getenv("CHANNEL_ID")
	STUDENT_ID = os.Getenv("STUDENT_ID")
	PASSWORD = os.Getenv("PASSWORD")
}

func process() error {
	getter := NewKadaiGetter(STUDENT_ID, PASSWORD)

	notifier := NewNotifier(TOKEN, CHANNEL_ID)

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
