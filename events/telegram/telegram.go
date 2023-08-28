package telegram

import "github.com/RobertKhachatryan/telegram_bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	//storage
}

func New(clent *telegram.Client)
