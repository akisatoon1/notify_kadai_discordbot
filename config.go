package main

import (
	"flag"
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
	f := loadFlags()

	if err := env(f.TestMode); err != nil {
		return nil, err
	}

	config := &Config{
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		ChannelID:    os.Getenv("CHANNEL_ID"),
		StudentID:    os.Getenv("STUDENT_ID"),
		Password:     os.Getenv("PASSWORD"),
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func env(testMode bool) error {
	if testMode {
		log.Println("テストモードで実行中です。.envファイルから環境変数を読み込みます。")
		err := godotenv.Load()
		return err
	}
	return nil
}

func (c *Config) validate() error {
	if c.DiscordToken == "" {
		return fmt.Errorf("DISCORD_TOKENが設定されていません")
	}
	if c.ChannelID == "" {
		return fmt.Errorf("CHANNEL_IDが設定されていません")
	}
	if c.StudentID == "" {
		return fmt.Errorf("STUDENT_IDが設定されていません")
	}
	if c.Password == "" {
		return fmt.Errorf("PASSWORDが設定されていません")
	}
	return nil
}

type flags struct {
	TestMode bool
}

func loadFlags() *flags {
	t := flag.Bool("t", false, "テストモードで実行")
	flag.Parse()

	f := &flags{}
	f.TestMode = *t

	return f
}
