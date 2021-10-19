package bank_account

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const getCommandHelpString = `/get__bank__bank_account $ACCOUNT_ID — get a entity
	- ACCOUNT_ID - int, id of account`

func (c *BankBankAccountCommander) Get(inputMessage *tgbotapi.Message) {

	args := newArgsScanner(inputMessage.CommandArguments())
	idx, err := args.nextUInt64()
	if err != nil {
		log.Println(err)

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			`/get__bank__bank_account $ACCOUNT_ID
expect ACCOUNT_ID integer number`,
		)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("BankBankAccountCommander.Get: error sending reply message to chat - %v", err)
		}

		return
	}

	account, err := c.bankAccountService.Describe(idx)
	if err != nil {
		errText := fmt.Sprintf("fail to get account with idx %d: %v", idx, err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, errText)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("BankBankAccountCommander.Get: error sending reply message to chat - %v", err)
		}
		log.Print(errText)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		account.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBankAccountCommander.Get: error sending reply message to chat - %v", err)
	}
}
