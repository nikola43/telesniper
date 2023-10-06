package main

import (
	"fmt"

	"github.com/mymmrac/telego"
)

func HandleAction(chatID int64, action string, bot *telego.Bot) error {
	_, err := SendMessage(chatID, INPUT_CAPTIONS[action], nil, bot)
	if err != nil {
		return err
	}
	state[chatID]["InputMode"] = action
	return nil
}

func ActionStart(chatID int64, bot *telego.Bot) error {
	fmt.Println("start")

	_, err := SendMessage(chatID, "welcome", nil, bot)
	if err != nil {
		return err
	}

	txt := "Welcome to the wallet bot. Please choose an option below."

	_, err = SendMessage(chatID, txt, WALLET_KEYBOARD, bot)
	if err != nil {
		return err
	}

	return nil
}

func ActionGenWallet(chatID int64, bot *telego.Bot) error {
	account := GenerateWallet()
	txt := "⚠️ Save your wallet ⚠️\n\n" + "Address: " + account.Address + "\n" + "Private Key: " + account.PrivateKey
	state[chatID]["Account"] = account

	HandleConfirm(Start, ShowMenu, chatID, txt, bot)
	return nil
}

func HandleActionWithKeyboard(chatID int64, action string, bot *telego.Bot) error {
	keyboard := KEYBOARDS[action]
	_, err := SendMessage(chatID, "Welcome", keyboard, bot)
	if err != nil {
		return err
	}
	return nil
}


func ActionBack(chatID int64, bot *telego.Bot) error {
	fmt.Println("start")

	_, err := SendMessage(chatID, "welcome", nil, bot)
	if err != nil {
		return err
	}

	txt := "Welcome to the wallet bot. Please choose an option below."

	_, err = SendMessage(chatID, txt, WALLET_KEYBOARD, bot)
	if err != nil {
		return err
	}
	return nil
}

