package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	tgbot "github.com/wawan93/bot-framework"
)

func main() {
	token := os.Getenv("TOKEN")
	log.Printf("token=%v", token)

	var chatID int64 = 178132

	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	api.Debug = false

	log.Printf("logged in as %v", api.Self.UserName)

	bot := tgbot.NewBotFramework(api)

	var voices []string

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(pwd+"/voices/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		msg := tgbotapi.NewVoiceUpload(chatID, path)
		m, err := bot.Send(msg)
		if err != nil {
			err = errors.New("can't send: " + err.Error())
			log.Println(err)
			return err
		}
		voices = append(voices, m.Voice.FileID)
		return nil
	})

	f, err := os.OpenFile("voices.json", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(f).Encode(voices); err != nil {
		log.Fatal(err)
	}

}
