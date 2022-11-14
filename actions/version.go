package actions

import (
	"fmt"
	"io"

	"github.com/urfave/cli/v2"
)

var version = "0.0.1"

func Version(w io.Writer) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		fmt.Fprintln(w, version)

		return nil
	}
}
