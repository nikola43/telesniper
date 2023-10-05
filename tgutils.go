package main

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleMessage(update telego.Update, bot *telego.Bot) error {
	botMessageID := state["BotMessageID"].(int)
	inputMode := state["InputMode"].(string)

	chatID := update.Message.Chat.ID
	msgText := update.Message.Text

	switch msgText {
	case Start:
		ActionStart(chatID, bot)
	}

	HandleInput(inputMode, update, bot)
	DeleteMessage(chatID, botMessageID, bot)

	return nil
}

func HandleInput(inputMode string, update telego.Update, bot *telego.Bot) error {
	chatID := update.Message.Chat.ID
	msgText := update.Message.Text

	switch inputMode {
	case ImportWallet:
		if msgText[0:2] == "0x" {
			msgText = msgText[2:]
		}
		if len(msgText) != 64 {
			txt := "Invalid private key, please try again."
			SendMessage(chatID, txt, nil, bot)
			return nil
		}

		account := ImportAccount(msgText)
		txt := "Address: " + account.Address + "\n" + "Private Key: " + account.PrivateKey
		state["Account"] = account
		HandleConfirm(Start, ShowMenu, chatID, txt, bot)

	case GenWallet:
		account := GenerateWallet()
		txt := "Address: " + account.Address + "\n" + "Private Key: " + account.PrivateKey
		msg, err := SendMessage(chatID, txt, nil, bot)
		if err != nil {
			return err
		}
		_ = msg
		state["Account"] = account
	}

	DeleteMessage(chatID, update.Message.MessageID, bot)
	return nil

}

/*
func BuildMessage(chatId int64, msg string, btnLabels []string, callbacks []string, bot *telego.Bot) (*telego.Message, error) {
	btns := make([][]telego.InlineKeyboardButton, 0)
	for i := 0; i < len(btnLabels); i++ {
		btns = append(btns, []telego.InlineKeyboardButton{
			tu.InlineKeyboardButton(btnLabels[i]).WithCallbackData(callbacks[i]),
		})
	}

	message := tu.Message(
		tu.ID(chatId),
		msg,
	).WithReplyMarkup(tu.InlineKeyboard(
		btns...,
	))

	// Sending message
	res, err := bot.SendMessage(message)
	if err != nil {
		return nil, err
	}
	return res, nil
}
*/

func SendMessage(chatId int64, msg string, replyMarkup telego.ReplyMarkup, bot *telego.Bot) (*telego.Message, error) {
	var message *telego.SendMessageParams

	if replyMarkup != nil {
		message = tu.Message(
			tu.ID(chatId),
			msg,
		).WithReplyMarkup(replyMarkup)
	} else {
		message = tu.Message(
			tu.ID(chatId),
			msg,
		)
	}
	// Sending message
	res, err := bot.SendMessage(message)
	if err != nil {
		return nil, err
	}

	state["BotMessageID"] = res.MessageID
	return res, nil
}

func HandleCallback(callback *telego.CallbackQuery, bot *telego.Bot) error {
	fmt.Println("Received callback with data:", callback.Data)
	fmt.Println("Received callback with message:", callback.Message.Text)
	fmt.Println("Received callback with chat id:", callback.Message.Chat.ID)

	// get chat id
	chatID := callback.Message.Chat.ID

	// replace message with callback data

	switch callback.Data {
	case ImportWallet:
		ActionImportWallet(chatID, bot)
		//HandleConfirm(Back, callbackQuery.Data, chatID, callbackQuery.Message.Text, bot)
	case GenWallet:
		ActionGenWallet(chatID, bot)

	case ShowMenu:
		ActionShowMainMenu(chatID, bot)

	case Disconnect:
		ActionDisconnect(chatID, bot)
	case Back:
		ActionBack(chatID, bot)
	}

	// send message with callback data
	//SendMessage(chatID, callbackQuery.Message.Text, nil, bot)
	return nil
}

func DeleteMessage(chatID int64, msgId int, bot *telego.Bot) {
	params := &telego.DeleteMessageParams{
		ChatID:    telego.ChatID{ID: chatID},
		MessageID: msgId,
	}
	err := bot.DeleteMessage(params)
	if err != nil {
		fmt.Println(err)
	}
}

func HandleUpdates(updates <-chan telego.Update, bot *telego.Bot) error {
	for update := range updates {
		// handle messages
		if update.Message != nil {
			err := HandleMessage(update, bot)
			if err != nil {
				return err
			}
		}

		// handle callback
		if update.CallbackQuery != nil {
			err := HandleCallback(update.CallbackQuery, bot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func HandleConfirm(backRoute, proceedRoute string, chatID int64, msg string, bot *telego.Bot) {
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow( // Row 1
			tu.InlineKeyboardButton("🔙 Back"). // Column 1
								WithCallbackData(Back+"/"+backRoute),
			tu.InlineKeyboardButton("✅ Proceed"). // Column 2
								WithCallbackData(proceedRoute),
		),
	)

	message := tu.Message(
		tu.ID(chatID),
		msg,
	).WithReplyMarkup(inlineKeyboard)

	// Sending message
	res, err := bot.SendMessage(message)
	if err != nil {
		fmt.Println(err)
	}

	state["BotMessageID"] = res.MessageID
	_ = res

}
