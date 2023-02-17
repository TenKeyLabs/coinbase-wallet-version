package actions

import (
	"os"
	"path/filepath"

	"github.com/tenkeylabs/cwv/util"
	"github.com/urfave/cli/v2"
)

func Bundle(zipper util.Zipper) func(*cli.Context) error {
	return func(*cli.Context) error {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		// Path to all downloaded extensions versions
		extensionRoot := filepath.Join(homeDir, `Library/Application Support/Google/Chrome/Default/Extensions/hnfanknocfeofbddgcijnmhnfnkdnaad`)

		// Bundle extension versions
		zipper(extensionRoot)

		return nil
	}
}
