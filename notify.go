package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type Notifier interface {
	Notify(ks []Kadai) error
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

func (n *notifier) Notify(ks []Kadai) error {
	session, err := discordgo.New("Bot " + n.token)
	if err != nil {
		return err
	}
	defer session.Close()

	_, err = session.ChannelMessageSend(n.channelID, createMsg(ks))
	if err != nil {
		return err
	}

	return nil
}

func createMsg(ks []Kadai) string {
	msg := time.Now().Format(time.DateTime) + "\n"
	if len(ks) == 0 {
		msg += "現在、課題はありません。\n"
	} else {
		msg += "課題一覧:\n"
		for _, k := range ks {
			msg += "-------------------------------\n"
			msg += "タイトル:	" + k.title + "\n"
			msg += "科目:			 " + k.course + "\n"
			msg += "締切:			 " + k.deadline.Format(time.DateTime) + "\n"
		}
	}
	msg += "^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n"
	return msg
}
