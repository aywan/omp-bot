package bank_account

import (
	"encoding/json"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const defaultPerPage = 5

func (c *BankBankAccountCommander) List(inputMessage *tgbotapi.Message, page uint64) {
	outputMsgText := fmt.Sprintf("Here the accounts on page #%d: \n\n", page+1)

	accounts, err := c.bankAccountService.List(page*defaultPerPage, defaultPerPage+1)

	if err != nil {
		log.Printf("BankBankAccountCommander.List: error sending reply message to chat - %v", err)

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Error while get accounts. Please, try again")
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("BankBankAccountCommander.List: error sending reply message to chat - %v", err)
		}

		return
	}

	hasNext := len(accounts) > defaultPerPage
	if hasNext {
		accounts = accounts[:len(accounts)-1]
	}

	if len(accounts) > 0 {
		for _, p := range accounts {
			outputMsgText += p.String()
			outputMsgText += "\n"
		}
	} else {
		outputMsgText += "No accounts found"
		page = 0
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	// Previous page
	if page > 0 {
		// to first page
		buttons = append(buttons, createButtonToPage("⏮", 0))

		prevPage := page - 1
		buttons = append(buttons, createButtonToPage(
			fmt.Sprintf("⬅ #%d", prevPage+1),
			prevPage,
		))
	}

	// Refresh current page
	buttons = append(buttons, createButtonToPage("🔄 refresh", page))

	if hasNext {
		nextPage := page + 1
		buttons = append(buttons, createButtonToPage(
			fmt.Sprintf("#%d ️➡", nextPage+1),
			nextPage,
		))
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBankAccountCommander.List: error sending reply message to chat - %v", err)
	}
}

func createButtonToPage(caption string, page uint64) tgbotapi.InlineKeyboardButton {
	serializedData, _ := json.Marshal(CallbackListData{Page: page})

	callbackPath := path.CallbackPath{
		Domain:       "bank",
		Subdomain:    "bank_account",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return tgbotapi.NewInlineKeyboardButtonData(caption, callbackPath.String())
}
