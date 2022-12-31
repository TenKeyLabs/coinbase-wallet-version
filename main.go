package main

import (
	"log"
	"os"

	"github.com/osis/cwv/actions"
	"github.com/osis/cwv/util"

	"github.com/urfave/cli/v2"
)

// This redirects
var DOWNLOAD_URL = "https://clients2.google.com/service/update2/crx?response=redirect&os=mac&arch=arm64&os_arch=arm64&nacl_arch=arm&prod=chromiumcrx&prodchannel=&prodversion=109.0.5361.0&lang=en-US&acceptformat=crx3&x=id%3Dhnfanknocfeofbddgcijnmhnfnkdnaad%26installsource%3Dondemand%26uc"
var WEBSTORE_URL = "https://chrome.google.com/webstore/detail/coinbase-wallet-extension/hnfanknocfeofbddgcijnmhnfnkdnaad"

func main() {
	app := &cli.App{
		Name:  "cwv",
		Usage: "Coinbase Wallet Version",
		Commands: []*cli.Command{
			{
				Name:   "bundle",
				Usage:  "export locally installed versions of the chrome extension",
				Action: actions.Bundle(util.ZipDirs),
			},
			{
				Name:   "check",
				Usage:  "print the version of the latest chrome extension in the store",
				Action: actions.Check(os.Stdout, WEBSTORE_URL),
			},
			{
				Name:   "download",
				Usage:  "download the latest chrome extension from the store",
				Action: actions.Download(DOWNLOAD_URL),
			},
			{
				Name:   "version",
				Usage:  "print cwv version",
				Action: actions.Version(os.Stdout),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
