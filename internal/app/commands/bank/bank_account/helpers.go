package bank_account

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func parseAccountId(inputMessage *tgbotapi.Message) (uint64, error) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("wrong args: %s", args)
	}

	return idx, nil
}
