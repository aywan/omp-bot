package bank_account

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/bank"
	"log"
)

const newCommandHelpString = `/new__bank__bank_account $DATA — create a new entity
	- DATA - json object
		example: {"userId": 123, "isLegal": true, "number": "202020", "currency": "ISK"}
`

type NewBankAccountRequest struct {
	UserId   uint64 `json:"userId"`
	IsLegal  bool   `json:"isLegal"`
	Number   string `json:"number"`
	Currency string `json:"currency"`
}

func (c *BankBankAccountCommander) New(inputMessage *tgbotapi.Message) {

	args := newArgsScanner(inputMessage.CommandArguments())

	data, err := args.bytesToEnd()
	if err != nil {
		log.Println(err)
		sendErrorMessageToUserInNew(inputMessage, c)
		return
	}
	var request NewBankAccountRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		log.Println(err)
		sendErrorMessageToUserInNew(inputMessage, c)
		return
	}

	newBankAccount := bank.NewBankAccount(request.UserId, request.IsLegal, request.Number, request.Currency)

	newId, err := c.bankAccountService.Create(newBankAccount)
	if err != nil {
		log.Println(err)
		sendErrorMessageToUserInNew(inputMessage, c)
		return
	}
	account, err := c.bankAccountService.Describe(newId)
	if err != nil {
		log.Println(err)
		sendErrorMessageToUserInNew(inputMessage, c)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Create new model as \n\n"+account.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBank/new__bank__bank_account {\"userId\": 123, \"isLegal\": true, \"number\": \"202020\", \"currency\": \"ISK\"}AccountCommander.New: error sending reply message to chat - %v", err)
	}
}

func sendErrorMessageToUserInNew(inputMessage *tgbotapi.Message, c *BankBankAccountCommander) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "error parse arguments\n"+newCommandHelpString)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBankAccountCommander.NEW: error sending reply message to chat - %v", err)
	}
}
