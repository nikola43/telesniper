package main

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// actions
const (
	Back         = "back"
	Proceed      = "proceed"
	Start        = "start"
	ImportWallet = "import_wallet"
	GenWallet    = "gen_wallet"
	ShowMenu     = "show_menu"
	Disconnect   = "disconnect"
	AddToken     = "add_token"
)

var INPUT_CAPTIONS = map[string]string{
	"Welcome":    "Welcome to app",
	ImportWallet: "Paste your private key:",
	GenWallet:    "Will generate a new wallet. Are you sure?",
	AddToken:     "Paste token address:",
}

var MAIN_MENU_KEYBOARD = tu.InlineKeyboard(
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("Add Token").WithCallbackData(AddToken),
		tu.InlineKeyboardButton("Buy").WithCallbackData("callback_1"),
		tu.InlineKeyboardButton("Continue").WithCallbackData("callback_1"),
		tu.InlineKeyboardButton("Continue").WithCallbackData("callback_1"),
	),
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("❌ Disconnect").WithCallbackData(Disconnect),
	),
)

var WALLET_KEYBOARD = tu.InlineKeyboard(
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("Import wallet").WithCallbackData(ImportWallet),
		tu.InlineKeyboardButton("Generate wallet").WithCallbackData(GenWallet),
	),
)

var KEYBOARDS = map[string]*telego.InlineKeyboardMarkup{
	ShowMenu: MAIN_MENU_KEYBOARD,
	Start:    WALLET_KEYBOARD,
}
