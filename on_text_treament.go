package main

import (
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
)

// BanRequest ban an user from a group
func BanRequest(m *tb.Message) {
	var member tb.ChatMember
	member.User = m.ReplyTo.Sender
	member.RestrictedUntil = 0
	b.Ban(m.Chat, &member)
}

// Config receives the user's info and make an account to him
func Config(m *tb.Message) {
	var UserInfo string

	user := m.Chat
	name := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	id := user.ID
	UserInfo = fmt.Sprintf("%s, %d\n", name, id)
	b.Send(m.Sender, UserInfo)
	for {
		b.Send(m.Sender, "Reply me your welcome words")
		if m.IsReply() == true {
			words := m.Text
			fmt.Println(words)
		} else {
			break
		}
	}
}
