package main

type Account struct {
	Address    string
	PrivateKey string
}

type TokenConfig struct {
	Address string
	Name    string
	Symbol  string
	BuyAmount float64
}
