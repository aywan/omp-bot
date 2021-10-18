package bank_account

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *BankBankAccountCommander) Delete(inputMessage *tgbotapi.Message) {

	idx, err := parseAccountId(inputMessage)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			`/delete__bank__bank_account $ACCOUNT_ID
expect ACCOUNT_ID integer number`,
		)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("BankBankAccountCommander.Delete: error sending reply message to chat - %v", err)
		}

		log.Println(err)
		return
	}

	_, err = c.bankAccountService.Remove(idx)
	if err != nil {
		errText := fmt.Sprintf("fail to delete account with idx %d: %v", idx, err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, errText)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("BankBankAccountCommander.Delete: error sending reply message to chat - %v", err)
		}
		log.Print(errText)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Succesfuly delete account with idx %d", idx),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBankAccountCommander.Delete: error sending reply message to chat - %v", err)
	}
}
