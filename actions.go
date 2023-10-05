package main

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ActionStart(chatID int64, bot *telego.Bot) error {
	fmt.Println("start")

	SendMessage(chatID, "welcome", nil, bot)

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
	DeleteMessage(chatID, state["BotMessageID"].(int), bot)
	fmt.Println("import wallet")
	txt := "Please send me your wallet file."
	msg, error := SendMessage(chatID, txt, nil, bot)
	if error != nil {
		return error
	}

	state["BotMessageID"] = msg.MessageID
	state["InputMode"] = ImportWallet

	return nil
}

func ActionGenWallet(chatID int64, bot *telego.Bot) error {
	DeleteMessage(chatID, state["BotMessageID"].(int), bot)
	account := GenerateWallet()
	txt := "⚠️ Save your wallet ⚠️\n\n" + "Address: " + account.Address + "\n" + "Private Key: " + account.PrivateKey
	state["Account"] = account

	HandleConfirm(Start, ShowMenu, chatID, txt, bot)
	return nil
}

func ActionShowMainMenu(chatID int64, bot *telego.Bot) {
	DeleteMessage(chatID, state["BotMessageID"].(int), bot)
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow( // Row 2
			tu.InlineKeyboardButton("Continue").WithCallbackData("callback_1"),
			tu.InlineKeyboardButton("Continue").WithCallbackData("callback_1"),
			tu.InlineKeyboardButton("Continue").WithCallbackData("callback_1"),
			tu.InlineKeyboardButton("Continue").WithCallbackData("callback_1"),
		),
		tu.InlineKeyboardRow( // Row 2
			tu.InlineKeyboardButton("❌ Disconnect").WithCallbackData(Disconnect),
		),
	)

	SendMessage(chatID, "Welcome", inlineKeyboard, bot)
}

func ActionDisconnect(chatID int64, bot *telego.Bot) {
	DeleteMessage(chatID, state["BotMessageID"].(int), bot)

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
}

func ActionBack(chatID int64, bot *telego.Bot) {
	DeleteMessage(chatID, state["BotMessageID"].(int), bot)

	/*
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
	*/
}
