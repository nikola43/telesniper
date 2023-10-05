package main

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func DeleteMessage(params *telego.DeleteMessageParams, bot *telego.Bot) {
	// Deleting message
	err := bot.DeleteMessage(params)
	if err != nil {
		fmt.Println(err)
	}
}

func HandleUpdates(updates <-chan telego.Update, bot *telego.Bot) {
	for update := range updates {
		// handle messages
		if update.Message != nil {
			err := HandleMessage(update, bot)
			if err != nil {
				fmt.Println(err)
			}
		}

		// handle callback
		if update.CallbackQuery != nil {
			err := HandleCallback(update.CallbackQuery, bot)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
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
	fmt.Println(res)

}
