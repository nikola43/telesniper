package main

import (
	"fmt"

	"github.com/mymmrac/telego"
)

func ActionAddToken(chatID int64, bot *telego.Bot) error {

	txt := INPUT_CAPTIONS[AddToken]
	_, err := SendMessage(chatID, txt, nil, bot)
	if err != nil {
		return err
	}

	state["InputMode"] = AddToken
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

func ActionImportWallet(chatID int64, bot *telego.Bot) error {

	txt := INPUT_CAPTIONS[ImportWallet]
	_, err := SendMessage(chatID, txt, nil, bot)
	if err != nil {
		return err
	}

	state["InputMode"] = ImportWallet
	return nil
}

func ActionGenWallet(chatID int64, bot *telego.Bot) error {

	account := GenerateWallet()
	txt := "⚠️ Save your wallet ⚠️\n\n" + "Address: " + account.Address + "\n" + "Private Key: " + account.PrivateKey
	state["Account"] = account

	HandleConfirm(Start, ShowMenu, chatID, txt, bot)
	return nil
}

func ActionShowMainMenu(chatID int64, bot *telego.Bot) error {

	_, err := SendMessage(chatID, "Welcome", MAIN_MENU_KEYBOARD, bot)
	if err != nil {
		return err
	}
	return nil
}

func ActionDisconnect(chatID int64, bot *telego.Bot) error {

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

func ActionBack(chatID int64, bot *telego.Bot) {
	DeleteMessage(chatID, state["BotMessageID"].(int), bot)
}
