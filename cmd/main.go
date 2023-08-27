package main

import (
	"flag"
	"log"
)

func main() {

	// t := mustToken()
	// token = flags.Get(token)

	//tgClient = telegram.New(token)

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