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
	AddToken     = "add_token"
)

var INPUT_CAPTIONS = map[string]string{
	ImportWallet: "Paste your private key:",
	GenWallet:    "Will generate a new wallet. Are you sure?",
	AddToken:     "Paste token address:",
}
