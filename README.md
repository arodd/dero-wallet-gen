# Dero Wallet Generator

This application will allow you to supply a suffix/prefix and generate wallet keys until a
matching suffix/prefix is located.  Once you've located the keys a seed output will be presented
which can be used to restore your wallet in the official wallet client.  

The process of searching for a wallet suffix is similar to mining where more complex 
suffixes will require more time to search for.  Simple 4 letter words with non-repeating
characters tend to have the best luck.

Speaking of luck, feel free to donate if you enjoy the use of this software :)

```
dero1qyxg6dmw22xh9v0hkp3xm6cns6qn9cwl9zp25mxxwmy7mvdea7akyqgetluck
```

# Usage
```shell
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
```

# Building

Local OS and architecture
```shell
go build .
```
Multiple Targets
```shell
./build.sh <version>
```