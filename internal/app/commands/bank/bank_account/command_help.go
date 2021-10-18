package bank_account

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *BankBankAccountCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		`Bank / Bank account commands:
/help__bank__bank_account — print list of commands
/get__bank__bank_account $ACCOUNT_ID — get a entity
/list__bank__bank_account — get a list of your entity
/delete__bank__bank_account $ACCOUNT_ID — delete an existing entity
/new__bank__bank_account — create a new entity
/edit__bank__bank_account $ACCOUNT_ID — edit a entity
`,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBankAccountCommander.Help: error sending reply message to chat - %v", err)
	}
}
