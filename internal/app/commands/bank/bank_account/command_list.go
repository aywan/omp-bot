package bank_account

import (
	"encoding/json"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const defaultLimit = 10

func (c *BankBankAccountCommander) List(inputMessage *tgbotapi.Message, cursor uint64) {
	outputMsgText := "Here all the accounts: \n\n"

	accounts, err := c.bankAccountService.List(cursor, defaultLimit+1)

	if err != nil {
		log.Printf("BankBankAccountCommander.List: error sending reply message to chat - %v", err)

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Error while get accounts. Please, try again")
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("BankBankAccountCommander.List: error sending reply message to chat - %v", err)
		}

		return
	}

	hasNext := len(accounts) > defaultLimit
	if hasNext {
		accounts = accounts[:len(accounts)-1]
	}
	showPagerKeyboard := true

	if len(accounts) > 0 {
		for _, p := range accounts {
			outputMsgText += p.String()
			outputMsgText += "\n"
		}
	} else {
		outputMsgText += "No accounts found"
		showPagerKeyboard = false
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if showPagerKeyboard {
		if hasNext {
			serializedData, _ := json.Marshal(CallbackListData{Cursor: accounts[len(accounts)-1].GetId()})

			callbackPath := path.CallbackPath{
				Domain:       "bank",
				Subdomain:    "bank_account",
				CallbackName: "list",
				CallbackData: string(serializedData),
			}

			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
				),
			)
		} else {
			serializedData, _ := json.Marshal(CallbackListData{Cursor: 0})

			callbackPath := path.CallbackPath{
				Domain:       "bank",
				Subdomain:    "bank_account",
				CallbackName: "list",
				CallbackData: string(serializedData),
			}

			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Return to start", callbackPath.String()),
				),
			)
		}
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBankAccountCommander.List: error sending reply message to chat - %v", err)
	}
}
