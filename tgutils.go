package main

import (
	"fmt"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleInput(inputMode string, update telego.Update, bot *telego.Bot) error {
	chatID := update.Message.Chat.ID
	msgText := update.Message.Text

	switch inputMode {
	// --------------------- Wallet ---------------------
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
		state[chatID]["Account"] = account
		HandleConfirm(Start, ShowMenu, chatID, txt, bot)

	// --------------------- GenWallet ---------------------
	case GenWallet:
		account := GenerateWallet()
		txt := "Address: " + account.Address + "\n" + "Private Key: " + account.PrivateKey
		msg, err := SendMessage(chatID, txt, nil, bot)
		if err != nil {
			return err
		}
		_ = msg
		state[chatID]["Account"] = account

	// --------------------- AddToken ---------------------
	case AddToken:
		if len(msgText) != 42 {
			txt := "Invalid token address, please try again."
			SendMessage(chatID, txt, nil, bot)
			return nil
		}

		tokenConfig := TokenConfig{
			Address: msgText,
		}

		state[chatID]["TokenConfig"] = tokenConfig
	}

	err := DeleteMessage(chatID, update.Message.MessageID, bot)
	if err != nil {
		return err
	}
	return nil
}

func HandleButtonCallback(callback *telego.CallbackQuery, bot *telego.Bot) error {
	fmt.Println("Received callback with data:", callback.Data)
	fmt.Println("Received callback with message:", callback.Message.Text)
	fmt.Println("Received callback with chat id:", callback.Message.Chat.ID)

	// get chat id
	chatID := callback.Message.Chat.ID

	if state[chatID]["BotMessageID"] != nil {
		err := DeleteMessage(chatID, state[chatID]["BotMessageID"].(int), bot)
		if err != nil {
			return err
		}
	}

	// check fi callback data contais "back" word
	isBack := strings.Contains(callback.Data, Back)
	if isBack {
		// check if callback data is back/..., if so, handle back
		back := strings.Split(callback.Data, "/")[0]
		path := strings.Split(callback.Data, "/")[1]
		fmt.Println("back:", back)
		fmt.Println("path:", path)
		HandleBack(chatID, path, bot)
	}

	switch callback.Data {
	case ImportWallet:
		HandleAction(chatID, ImportWallet, bot)
	case GenWallet:
		ActionGenWallet(chatID, bot)
	case ShowMenu:
		HandleActionWithKeyboard(chatID, ShowMenu, bot)
	case Disconnect:
		HandleActionWithKeyboard(chatID, Start, bot)
	case AddToken:
		HandleAction(chatID, ImportWallet, bot)
	}
	return nil
}

func BuildKeyboard(btnLabels []string, callbacks []string) telego.ReplyMarkup {
	btns := make([][]telego.InlineKeyboardButton, 0)
	for i := 0; i < len(btnLabels); i++ {
		btns = append(btns, []telego.InlineKeyboardButton{
			tu.InlineKeyboardButton(btnLabels[i]).WithCallbackData(callbacks[i]),
		})
	}

	keyboard := tu.InlineKeyboard(
		btns...,
	)
	return keyboard
}

func HandleMessage(update telego.Update, bot *telego.Bot) error {

	chatID := update.Message.Chat.ID
	msgText := update.Message.Text

	if state[chatID] == nil {
		state[chatID] = make(map[string]interface{})
		state[chatID]["BotMessageID"] = 0
		state[chatID]["InputMode"] = ""
		state[chatID]["Account"] = nil
	}
	fmt.Println(state[chatID])

	botMessageID := state[chatID]["BotMessageID"].(int)
	inputMode := state[chatID]["InputMode"].(string)

	switch msgText {
	case "/start":
		_, err := SendMessage(chatID, "welcome", nil, bot)
		if err != nil {
			return err
		}
		ActionStart(chatID, bot)
	}

	if inputMode != "" {
		HandleInput(inputMode, update, bot)
		DeleteMessage(chatID, botMessageID, bot)
	}

	return nil
}

/*
/*
func HandleMessage(update telego.Update, bot *telego.Bot) error {

	chatID := update.Message.Chat.ID
	msgText := update.Message.Text

	if state[chatID] == nil {
		state[chatID] = make(map[string]interface{})
	}

	inputMode, ok := state[chatID]["InputMode"].(string)
	botMessageID, ok = state[chatID]["BotMessageID"].(int)


	switch msgText {
	case Start:
		_, err := SendMessage(chatID, "welcome", nil, bot)
		if err != nil {
			return err
		}
		ActionStart(chatID, bot)
	}

	HandleInput(inputMode, update, bot)
	//DeleteMessage(chatID, botMessageID, bot)

	return nil
}

*/

func SendMessage(chatID int64, msg string, replyMarkup telego.ReplyMarkup, bot *telego.Bot) (*telego.Message, error) {
	var message *telego.SendMessageParams

	if replyMarkup != nil {
		message = tu.Message(
			tu.ID(chatID),
			msg,
		).WithReplyMarkup(replyMarkup)
	} else {
		message = tu.Message(
			tu.ID(chatID),
			msg,
		)
	}
	// Sending message
	res, err := bot.SendMessage(message)
	if err != nil {
		return nil, err
	}

	state[chatID]["BotMessageID"] = res.MessageID
	return res, nil
}

func DeleteMessage(chatID int64, msgId int, bot *telego.Bot) error {
	params := &telego.DeleteMessageParams{
		ChatID:    telego.ChatID{ID: chatID},
		MessageID: msgId,
	}
	err := bot.DeleteMessage(params)
	if err != nil {
		return err
	}
	return nil
}

func HandleUpdates(bot *telego.Bot) error {
	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		return err
	}
	defer bot.StopLongPolling()

	for update := range updates {
		// handle messages
		if update.Message != nil {
			err := HandleMessage(update, bot)
			if err != nil {
				return err
			}
		}

		// handle button callback
		if update.CallbackQuery != nil {
			err := HandleButtonCallback(update.CallbackQuery, bot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func HandleBack(chatID int64, backRoute string, bot *telego.Bot) {
	switch backRoute {
	case Start:
		ActionStart(chatID, bot)
	}
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

	state[chatID]["BotMessageID"] = res.MessageID
	_ = res

}
