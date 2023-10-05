package main

import tu "github.com/mymmrac/telego/telegoutil"

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
