#!/bin/bash
export GOOS=linux; export GOARCH=amd64; export CGO_ENABLED=0
go build -o dero-wallet-gen-$1-$GOOS-$GOARCH .
export GOOS=linux; export GOARCH=arm64
go build -o dero-wallet-gen-$1-$GOOS-$GOARCH .
export GOOS=linux; export GOARCH=arm; export GOARM=7
go build -o dero-wallet-gen-$1-$GOOS-armv7 .
export GOOS=android; export GOARCH=arm64
go build -o dero-wallet-gen-$1-$GOOS-$GOARCH .
export GOOS=darwin; export GOARCH=amd64
go build -o dero-wallet-gen-$1-$GOOS-$GOARCH .
export GOOS=darwin; export GOARCH=arm64
go build -o dero-wallet-gen-$1-$GOOS-$GOARCH .
export GOOS=windows; export GOARCH=amd64
go build -o dero-wallet-gen-$1-$GOOS-$GOARCH