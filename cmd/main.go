package main

import (
	"flag"
	"log"
	// "github.com/RobertKhachatryan/telegram_bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	// tgClient = telegram.New(mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	//consumer.Start(fetcher,processor)
}

func mustToken() string {
	// bot -tg-bot-token "token"
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("Token is not specified")
	}

	return *token
}
