package main

// actions
const (
	Back         = "back"
	Proceed      = "proceed"
	Start        = "/start"
	ImportWallet = "import_wallet"
	GenWallet    = "gen_wallet"
	ShowMenu     = "show_menu"
	Disconnect   = "disconnect"
)

var INPUT_CAPTIONS = map[string]string{
	ImportWallet: "Are you sure you want to load wallet?",
	GenWallet:    "Are you sure you want to load wallet?",
}
