package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"notify_kadai_discordbot/internal/ds"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken string
	ChannelID    string
	StudentID    string
	Password     string

	ExcludedCourses ds.Set
}

func LoadConfig() (*Config, error) {
	f := loadFlags()

	if err := env(f.TestMode); err != nil {
		return nil, err
	}

	excluded := ds.NewSet()
	if f.ExcludePath != "" {
		ex, err := loadExcludedCourses(f.ExcludePath)
		if err != nil {
			return nil, err
		}
		excluded = ex
	}

	config := &Config{
		DiscordToken:    os.Getenv("DISCORD_TOKEN"),
		ChannelID:       os.Getenv("CHANNEL_ID"),
		StudentID:       os.Getenv("STUDENT_ID"),
		Password:        os.Getenv("PASSWORD"),
		ExcludedCourses: excluded,
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
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
	if c.ExcludedCourses == nil {
		return fmt.Errorf("除外コースのセットが初期化されていません")
	}
	return nil
}

func env(testMode bool) error {
	if testMode {
		log.Println("テストモードで実行中です。.envファイルから環境変数を読み込みます。")
		err := godotenv.Load()
		return err
	}
	return nil
}

func loadExcludedCourses(path string) (ds.Set, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("除外ファイルの読み込みに失敗しました: %w", err)
	}
	defer file.Close()

	excludedCourses := ds.NewSet()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		excludedCourses.Add(line)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("除外ファイルの読み込み中にエラーが発生しました: %w", err)
	}

	return excludedCourses, nil
}

type flags struct {
	TestMode    bool
	ExcludePath string
}

func loadFlags() *flags {
	t := flag.Bool("t", false, "テストモードで実行")
	e := flag.String("exclude", "", "通知しないコースを設定するファイルのパスを指定")
	flag.Parse()

	f := &flags{}
	f.TestMode = *t
	f.ExcludePath = *e

	return f
}
