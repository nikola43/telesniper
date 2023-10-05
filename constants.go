package main

// actions
const (
	Back         = "back"
	Proceed      = "proceed"
	Start        = "/start"
	ImportWallet = "import_wallet"
	GenWallet    = "gen_wallet"
)

var INPUT_CAPTIONS = map[string]string{
	ImportWallet: "Are you sure you want to load wallet?",
	GenWallet:    "Are you sure you want to load wallet?",
}