package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	bot    = Bot{}
	b, err = tb.NewBot(tb.Settings{
		Token:  bot.GetToken(),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second}})
)

func main() {

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, OnTextTreatment)
	b.Handle(tb.OnAddedToGroup, OnTextTreatment)

	b.Handle("/configure", Config)
	b.Handle("/ban", BanRequest)
	
	b.Start()
}

func terminalprint(id int64, text string) {
	fmt.Printf("%d said: %s\n", id, text)
}

// OnTextTreatment returns every action from the bot
func OnTextTreatment(m *tb.Message) {
	if m.IsReply() {
		fmt.Printf("Reply From %s to %s", m.Sender.FirstName, m.ReplyTo.Sender.FirstName)
	}
	switch {
	case strings.Contains(m.Text, "/start"):
		b.Send(m.Sender, Botsays["intro"])
	case strings.Contains(m.Text, "/help"):
		b.Send(m.Sender, Botsays["help"])
	case strings.Contains(m.Text, "which system do you run?"):
		b.Send(m.Sender, Botsays["system"])
	}

	terminalprint(m.Chat.ID, m.Text)
}
