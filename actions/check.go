package actions

import (
	"fmt"
	"io"

	"github.com/antchfx/htmlquery"
	"github.com/urfave/cli/v2"
)

func Check(w io.Writer, url string) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		// Load the raw version of the store front
		doc, err := htmlquery.LoadURL(url)
		if err != nil {
			return err
		}

		// Extract version from meta element
		metaNode, err := htmlquery.Query(doc, `//div[text()="Version"]/following-sibling::*`)
		if err != nil {
			return err
		}

		fmt.Fprintln(w, htmlquery.InnerText(metaNode))

		return nil
	}
}
