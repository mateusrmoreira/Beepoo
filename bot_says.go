package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
)

// system returns the operative system which runs the Bot APP
var system = fmt.Sprintf("I run on a %s computer", runtime.GOOS)

// Bot is a struct that gets from a json the
// Bot's information
type Bot struct {
	Token string `json:"token"`
}

// GetToken returns a token
func (b *Bot) GetToken() string {
	configJSON, _ := ioutil.ReadFile("../beepoocfg/config.json")
	json.Unmarshal([]byte(configJSON), &b)
	return b.Token
}

// Botsays map[string]string that returns all things that bot says
// This map has the follow comands intro, help, rules.
var Botsays = map[string]string{
	"intro":  `Hello I'm Beepoo the Bot... What's your name?`,
	"help":   "I can send you a cat pic just say: Send me the cat",
	"system": system,
}
