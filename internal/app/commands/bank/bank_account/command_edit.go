package bank_account

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/big"
)

const editCommandHelpString = `/edit__bank__bank_account $ACCOUNT_ID $DATA — edit a entity 
	-ACCOUNT_ID - int, id of account
	-DATA - json object
		example: {"totalAmount": '0.0001', "currency": "ETH"}`

type EditBankAccountRequest struct {
	TotalAmount big.Float `json:"totalAmount"`
	Currency    string    `json:"currency"`
}

func (c *BankBankAccountCommander) Edit(inputMessage *tgbotapi.Message) {

	args := newArgsScanner(inputMessage.CommandArguments())

	idx, err := args.nextUInt64()
	if err != nil {
		log.Println(err)
		sendErrorMessageToUserInEdit(inputMessage, c, "error parse arguments")
		return
	}

	model, err := c.bankAccountService.Describe(idx)
	if err != nil {
		log.Println(err)
		sendErrorMessageToUserInEdit(inputMessage, c, fmt.Sprintf("error get model with idx=%d", idx))
		return
	}

	data, err := args.bytesToEnd()
	if err != nil {
		log.Println(err)
		sendErrorMessageToUserInEdit(inputMessage, c, "error parse arguments")
		return
	}

	var request EditBankAccountRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		log.Println(err)
		sendErrorMessageToUserInEdit(inputMessage, c, "error parse arguments")
		return
	}

	model.SetTotalAmount(request.TotalAmount)
	model.Currency = request.Currency

	err = c.bankAccountService.Update(idx, *model)
	if err != nil {
		log.Println(err)
		sendErrorMessageToUserInEdit(inputMessage, c, fmt.Sprintf("error update model idx=%d", idx))
		return
	}

	account, err := c.bankAccountService.Describe(idx)
	if err != nil {
		log.Println(err)
		sendErrorMessageToUserInEdit(inputMessage, c, "error get model")
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Model was updated\n\n"+account.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBankAccountCommander.Edit: error sending reply message to chat - %v", err)
	}
}

func sendErrorMessageToUserInEdit(inputMessage *tgbotapi.Message, c *BankBankAccountCommander, message string) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, message+"\n"+editCommandHelpString)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankBankAccountCommander.Edit: error sending reply message to chat - %v", err)
	}
}
