package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	tgutil "github.com/nikola43/telesniper/utils"
)

// actions
const (
	Back       = "/back"
	Proceed    = "/proceed"
	Start      = "/start"
	LoadWallet = "/load_wallet"
	GenWallet  = "/gen_wallet"
)

const (
	LoadWalletMsg = "Are you sure you want to load wallet?"
	GenWalletMsg  = "Are you sure you want to generate wallet?"
)

func main() {
	//botToken := os.Getenv("TOKEN")
	// init bot
	botToken := "5931596960:AAEpkYhCtdUj6PQhZeDjMW-QOOYJMeEVShA"
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Getting updates
	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer bot.StopLongPolling()

	// Receiving updates
	HandleUpdates(updates, bot)
}

func HandleMessage(update telego.Update, bot *telego.Bot) error {
	chatID := update.Message.Chat.ID
	msgText := update.Message.Text
	fmt.Println("Received message:", update.Message.Text)

	switch msgText {
	case Start:
		fmt.Println("start")
		btnLabels := []string{"Load wallet", "Generate wallet"}
		callbacks := []string{LoadWallet, GenWallet}

		message, err := BuildMessage(chatID, "Hello", btnLabels, callbacks, bot)
		if err != nil {
			return err
		}

		fmt.Println("mss", message)
	case LoadWallet:
		fmt.Println("load wallet")
	case GenWallet:
		fmt.Println("gen wallet")
	}

	return nil
}

func HandleCallback(callback *telego.CallbackQuery, bot *telego.Bot) error {
	callbackQuery := callback
	fmt.Println("Received callback with data:", callbackQuery.Data)
	fmt.Println("Received callback with message:", callbackQuery.Message.Text)
	fmt.Println("Received callback with chat id:", callbackQuery.Message.Chat.ID)

	// get chat id
	chatID := callbackQuery.Message.Chat.ID

	switch callbackQuery.Data {
	case LoadWallet:
		HandleConfirm("a", "b", chatID, callbackQuery.Message.Text, bot)
	}

	// send message with callback data
	//SendMessage(chatID, callbackQuery.Message.Text, nil, bot)
	return nil
}

func SendMessage(chatId int64, msg string, replyMarkup telego.ReplyMarkup, bot *telego.Bot) (*telego.Message, error) {
	var message *telego.SendMessageParams

	if replyMarkup == nil {
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

func BuildMessage(chatId int64, msg string, btnLabels []string, callbacks []string, bot *telego.Bot) (*telego.Message, error) {
	btns := make([][]telego.InlineKeyboardButton, 0)
	for i := 0; i < len(btnLabels); i++ {
		btns = append(btns, []telego.InlineKeyboardButton{
			tu.InlineKeyboardButton(btnLabels[i]).WithCallbackData(callbacks[i]),
		})
	}

	inlineKeyboard := tu.InlineKeyboard(
		btns...,
	)

	message := tu.Message(
		tu.ID(chatId),
		msg,
	).WithReplyMarkup(inlineKeyboard)

	// Sending message
	res, err := bot.SendMessage(message)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func handleStartCommand() {

}
