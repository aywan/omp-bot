package bank_account

import (
	"github.com/ozonmp/omp-bot/internal/service/bank/bank_account"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type BankBankAccountCommander struct {
	bot                *tgbotapi.BotAPI
	bankAccountService bank_account.ServiceInterface
}

func NewBankBankAccountCommander(
	bot *tgbotapi.BotAPI,
) *BankBankAccountCommander {
	bankAccountService := bank_account.NewDummyBankAccountService()

	return &BankBankAccountCommander{
		bot:                bot,
		bankAccountService: bankAccountService,
	}
}

func (c *BankBankAccountCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("BankBankAccountCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *BankBankAccountCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg, 0)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
