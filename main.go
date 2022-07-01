package main

import (
	"encoding/hex"
	"fmt"
	"github.com/deroproject/derohe/walletapi"
	"github.com/docopt/docopt-go"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
)

var suffix = ""
var prefix = ""
var contains = ""
var arguments = map[string]interface{}{}
var exit = make(chan bool)

var command_line string = `dero-wallet-gen
Generate Dero Wallet with matching suffix

Usage:
  dero-wallet-gen --suffix=<string>
  dero-wallet-gen --prefix=<string>
  dero-wallet-gen --contains=<string>
  dero-wallet-gen -h | --help

Options:
  -h --help           Show this screen.
  --suffix=<string>   Search for wallet with this string suffix
  --prefix=<string>   Search for wallet with this string prefix
  --contains=<string> Search for wallet with this string anywhere

Example: ./dero-wallet-gen --suffix dead

`
var threadwrite = sync.Mutex{}

func main() {
	arguments, _ = docopt.Parse(command_line, nil, true, "v0.0.1", false)
	if arguments["--suffix"] != nil {
		suffix = arguments["--suffix"].(string)
	} else if arguments["--prefix"] != nil {
		prefix = arguments["--prefix"].(string)
	} else if arguments["--contains"] != nil {
		contains = arguments["--contains"].(string)
	} else {
		return
	}

	threads := runtime.GOMAXPROCS(0)
	go func() {
		var gracefulStop = make(chan os.Signal, 1)
		signal.Notify(gracefulStop, os.Interrupt)
		for {
			sig := <-gracefulStop
			fmt.Fprintf(os.Stderr, "\nReceived signal %s\n", sig)

			if sig.String() == "interrupt" {
				close(exit)
			}
		}
	}()
	fmt.Fprintf(os.Stderr, "Generating wallets....this could take some time.\n")

	for i := 0; i < threads; i++ {
		if suffix != "" {
			go findWalletSuffix(suffix)
		} else if prefix != "" {
			go findWalletPrefix(prefix)
		} else if contains != "" {
			go findWalletContains(contains)
		} else {
			return
		}
	}
	<-exit
	return
}

func findWalletSuffix(suffix string) {
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

func findWalletPrefix(prefix string) {
	for {
		account, _ := walletapi.Generate_Keys_From_Random()
		account.SeedLanguage = "English"
		address := account.GetAddress()
		address.Mainnet = true
		mainaddr := address.String()
		if strings.HasPrefix(mainaddr, ("dero1qy" + prefix)) {
			getSeeds(account)
		}
	}
}

func findWalletContains(contains string) {
	for {
		account, _ := walletapi.Generate_Keys_From_Random()
		account.SeedLanguage = "English"
		address := account.GetAddress()
		address.Mainnet = true
		mainaddr := address.String()
		if strings.Contains(mainaddr, contains) {
			getSeeds(account)
		}
	}
}

func getSeeds(account *walletapi.Account) {
	bytes := account.Keys.Public.EncodeCompressed()
	private := account.Keys.Secret.String()
	w, err := walletapi.Create_Encrypted_Wallet_Memory("yes", account.Keys.Secret)
	if err == nil {
		address := w.GetAddress()
		address.Mainnet = true
		threadwrite.Lock()
		defer threadwrite.Unlock()
		fmt.Println(fmt.Sprintf("%-10s %s", "Public: ", hex.EncodeToString(bytes[:])))
		fmt.Println(fmt.Sprintf("%-10s %s", "Private: ", private))
		fmt.Println(fmt.Sprintf("%-10s %s", "Wallet: ", address))
		seed := w.GetSeed()
		fmt.Println(seed, "\n")
	}
}
