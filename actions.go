package main

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ActionStart(chatID int64, bot *telego.Bot) error {
	fmt.Println("start")
	btnLabels := []string{"Import wallet", "Generate wallet"}
	callbacks := []string{ImportWallet, GenWallet}

	btns := make([][]telego.InlineKeyboardButton, 0)
	for i := 0; i < len(btnLabels); i++ {
		btns = append(btns, []telego.InlineKeyboardButton{
			tu.InlineKeyboardButton(btnLabels[i]).WithCallbackData(callbacks[i]),
		})
	}

	k := tu.InlineKeyboard(
		btns...,
	)

	msg := "Welcome to the wallet bot. Please choose an option below."
	SendMessage(chatID, msg, k, bot)

	return nil
}

func ActionImportWallet(chatID int64, bot *telego.Bot) error {
	fmt.Println("import wallet")
	txt := "Please send me your wallet file."
	msg, error := SendMessage(chatID, txt, nil, bot)
	if error != nil {
		return error
	}

	botMessageID = msg.MessageID
	inputMode = ImportWallet

	return nil
}

func ActionGenWallet(chatID int64, bot *telego.Bot) error {
	fmt.Println("import wallet")
	msg := "Please send me your wallet file."
	SendMessage(chatID, msg, nil, bot)

	return nil
}
