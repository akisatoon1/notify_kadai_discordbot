package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken string
	ChannelID    string
	StudentID    string
	Password     string
}

func LoadConfig() (*Config, error) {
	if err := env(); err != nil {
		return nil, err
	}

	config := &Config{
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		ChannelID:    os.Getenv("CHANNEL_ID"),
		StudentID:    os.Getenv("STUDENT_ID"),
		Password:     os.Getenv("PASSWORD"),
	}

	if config.DiscordToken == "" || config.ChannelID == "" || config.StudentID == "" || config.Password == "" {
		return nil, fmt.Errorf("環境変数が設定されていません。DISCORD_TOKEN, CHANNEL_ID, STUDENT_ID, PASSWORDを設定してください。")
	}

	return config, nil
}

func env() error {
	if len(os.Args) == 2 && os.Args[1] == "-t" {
		log.Println("テストモードで実行中です。.envファイルから環境変数を読み込みます。")
		err := godotenv.Load()
		return err
	}
	return nil
}
