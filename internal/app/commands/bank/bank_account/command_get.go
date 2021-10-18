package bank_account

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BankBankAccountCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			`/get__bank__bank_account $ACCOUNT_ID
expect ACCOUNT_ID integer number`,
		)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("BankBankAccountCommander.Get: error sending reply message to chat - %v", err)
		}

		log.Println("wrong args", args)
		return
	}

	account, err := c.bankAccountService.Describe(idx)
	if err != nil {
		errText := fmt.Sprintf("fail to get product with idx %d: %v", idx, err)
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
