package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type Notifier interface {
	Notify(prefix string, ks []Kadai) error
}

type notifier struct {
	token     string
	channelID string
}

func NewNotifier(token, channelID string) Notifier {
	return &notifier{
		token:     token,
		channelID: channelID,
	}
}

func (n *notifier) Notify(prefix string, ks []Kadai) error {
	session, err := discordgo.New("Bot " + n.token)
	if err != nil {
		return err
	}
	defer session.Close()

	_, err = session.ChannelMessageSend(n.channelID, createMsg(prefix, ks))
	if err != nil {
		return err
	}

	return nil
}

func createMsg(prefix string, ks []Kadai) string {
	msg := prefix + "\n"

	msg += time.Now().Format(time.DateTime) + "\n"
	if len(ks) == 0 {
		msg += "現在、課題はありません。\n"
	} else {
		msg += "課題一覧:\n"
		for _, k := range ks {
			msg += "-------------------------------\n"
			msg += "タイトル:	" + k.title + "\n"
			msg += "科目:			 " + k.course + "\n"
			msg += "締切:			 " + format(k.deadline) + "\n"
		}
	}
	msg += "^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n"
	return msg
}

func format(t *time.Time) string {
	if t == nil {
		return "なし"
	}
	return t.Format(time.DateTime)
}
