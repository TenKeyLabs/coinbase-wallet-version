package actions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/urfave/cli/v2"
)

func Download(url string) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Read the incoming .crx's filename and adjust format
		rawFilename := path.Base(resp.Request.URL.String())

		// Create the file
		out, err := os.Create(rawFilename)
		if err != nil {
			return err
		}
		defer out.Close()

		// Check server response
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("bad status: %s", resp.Status)
		}

		// Writer the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}

		// Rename to zip and fix headers
		filename := fmt.Sprintf("coinbase-wallet-chrome-%s.zip", strings.Join(strings.Split(rawFilename, "_")[1:4], "."))
		cmd := exec.Command("zip", "-FF", rawFilename, "--out", filename)
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	}
}
