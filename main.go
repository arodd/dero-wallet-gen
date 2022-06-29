package main

import (
	"fmt"
	"github.com/deroproject/derohe/walletapi"
	"github.com/docopt/docopt-go"
	"os"
	"os/signal"
	"runtime"
	"strings"
)

var arguments = map[string]interface{}{}

var command_line string = `dero-wallet-gen
Generate Dero Wallet with matching suffix

Usage:
  dero-wallet-gen --suffix=<suffix>
  dero-wallet-gen -h | --help

Options:
  -h --help           Show this screen.
  --suffix=<suffix>   Search for wallet with this string suffix

Example: ./dero-wallet-gen --suffix dead
`
var suffix string
var exit = make(chan bool)

func main() {
	arguments, _ = docopt.Parse(command_line, nil, true, "v0.0.1", false)
	if arguments["--suffix"] != nil {
		suffix = arguments["--suffix"].(string)
	} else {
		return
	}
	threads := runtime.GOMAXPROCS(0)
	go func() {
		var gracefulStop = make(chan os.Signal, 1)
		signal.Notify(gracefulStop, os.Interrupt)
		for {
			sig := <-gracefulStop
			fmt.Printf("\nReceived signal %s\n", sig)

			if sig.String() == "interrupt" {
				close(exit)
			}
		}
	}()

	for i := 0; i < threads; i++ {
		go findWallet(suffix)
	}
	<-exit
	return
}

func findWallet(suffix string) {
	for {
		account, _ := walletapi.Generate_Keys_From_Random()
		account.SeedLanguage = "English"
		address := account.GetAddress()
		address.Mainnet = true
		mainaddr := address.String()
		if strings.HasSuffix(mainaddr, suffix) {
			getSeeds(account)
		}
	}
}

func getSeeds(account *walletapi.Account) {
	w, err := walletapi.Create_Encrypted_Wallet_Memory("yes", account.Keys.Secret)
	if err == nil {
		address := w.GetAddress()
		address.Mainnet = true
		fmt.Println(address)
		seed := w.GetSeed()
		fmt.Println(seed, "\n")
	}
}
