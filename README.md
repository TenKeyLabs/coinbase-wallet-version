# Coinbase Wallet Version CLI

`cwv` is a tool to download or bundle versions of the [Coinbase Wallet Chrome extension](https://chrome.google.com/webstore/detail/coinbase-wallet-extension/hnfanknocfeofbddgcijnmhnfnkdnaad?hl=en).

## Usage

```
NAME:
   cwv - Coinbase Wallet Version

USAGE:
   cwv [global options] command [command options] [arguments...]

COMMANDS:
   bundle    export locally installed versions of the chrome extension
   check     print the version of the latest chrome extension in the store
   download  download the latest chrome extension from the store
   version   print cwv version
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Installation

```bash
  git clone git@github.com:TenKeyLabs/coinbase-wallet-version.git
  go build main.go -o cwv
  cwv help
```

## Why

Mostly made this for fun and I hope a most official distribution channel for this extension is exposed to the community. Until then, this is will be used to power the [coinbase wallet versions](https://github.com/TenKeyLabs/coinbase-wallet-versions) archive.
