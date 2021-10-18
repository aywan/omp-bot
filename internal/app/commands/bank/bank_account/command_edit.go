package bank_account

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *BankBankAccountCommander) Edit(inputMessage *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"error: not implemented",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBankAccountCommander.Edit: error sending reply message to chat - %v", err)
	}
}
