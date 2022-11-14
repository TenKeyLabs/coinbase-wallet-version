package actions

import (
	"fmt"
	"io"
	"net/http"
	"os"
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

		// Read the incoming filename and adjust format
		fmt.Println(resp.Request.URL.String())
		rawFilename := path.Base(resp.Request.URL.String())
		filename := fmt.Sprintf("coinbase-wallet-chrome-%s.zip", strings.Join(strings.Split(rawFilename, "_")[1:4], "."))

		// Create the file
		out, err := os.Create(filename)
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

		return nil
	}
}
