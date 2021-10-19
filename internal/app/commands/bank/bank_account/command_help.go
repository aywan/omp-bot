package bank_account

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const helpCommandHelpString = `/help__bank__bank_account — print list of commands`

func (c *BankBankAccountCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Bank / Bank account commands:\n"+
			helpCommandHelpString+"\n"+
			getCommandHelpString+"\n"+
			listCommandHelpString+"\n"+
			deleteCommandDeleteString+"\n"+
			newCommandHelpString+"\n"+
			editCommandHelpString,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBankAccountCommander.Help: error sending reply message to chat - %v", err)
	}
}
