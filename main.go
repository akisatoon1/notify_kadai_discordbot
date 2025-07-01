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
	if len(os.Args) == 2 && os.Args[1] == "-t" {
		log.Println("テストモードで実行中です。.envファイルから環境変数を読み込みます。")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("環境変数の読み込みに失敗しました。.envファイルが存在するか確認してください。", err)
		}
	}

	TOKEN = os.Getenv("DISCORD_TOKEN")
	CHANNEL_ID = os.Getenv("CHANNEL_ID")
	STUDENT_ID = os.Getenv("STUDENT_ID")
	PASSWORD = os.Getenv("PASSWORD")

	if TOKEN == "" || CHANNEL_ID == "" || STUDENT_ID == "" || PASSWORD == "" {
		log.Fatal("環境変数が設定されていません。DISCORD_TOKEN, CHANNEL_ID, STUDENT_ID, PASSWORDを設定してください。")
	}

	log.Println("環境変数の読み込みに成功しました。")
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
		// 締め切りまで3日以内の課題をフィルター
		if k.deadline != nil && k.deadline.Before(time.Now().Add(72*time.Hour)) {
			notified = append(notified, k)
		}
	}

	return notifier.Notify("締め切りが近い課題です。", notified)
}
