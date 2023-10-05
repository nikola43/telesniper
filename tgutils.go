package main

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleMessage(update telego.Update, bot *telego.Bot) error {
	chatID := update.Message.Chat.ID
	msgText := update.Message.Text
	fmt.Println("Received message:", msgText)

	switch msgText {
	case Start:
		ActionStart(chatID, bot)
	}

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
	return res, nil
}

func HandleCallback(callback *telego.CallbackQuery, bot *telego.Bot) error {
	callbackQuery := callback
	fmt.Println("Received callback with data:", callbackQuery.Data)
	fmt.Println("Received callback with message:", callbackQuery.Message.Text)
	fmt.Println("Received callback with chat id:", callbackQuery.Message.Chat.ID)

	// get chat id
	chatID := callbackQuery.Message.Chat.ID

	switch callbackQuery.Data {
	case ImportWallet:
		ActionImportWallet(chatID, bot)
		//HandleConfirm(Back, callbackQuery.Data, chatID, callbackQuery.Message.Text, bot)
	case GenWallet:
		ActionGenWallet(chatID, bot)
	}

	// send message with callback data
	//SendMessage(chatID, callbackQuery.Message.Text, nil, bot)
	return nil
}

func DeleteMessage(params *telego.DeleteMessageParams, bot *telego.Bot) {
	// Deleting message
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
	//btnLabels := []string{"🔙 Back", "✅ Proceed"}
	//callbacks := []string{Back, Proceed}

	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow( // Row 1
			tu.InlineKeyboardButton("🔙 Back"). // Column 1
								WithCallbackData(Back+"/"+backRoute),
			tu.InlineKeyboardButton("✅ Proceed"). // Column 2
								WithCallbackData(Proceed+"/"+proceedRoute),
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
	//fmt.Println(res)
	_ = res

}
