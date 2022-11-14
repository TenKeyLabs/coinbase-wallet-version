package actions

import (
	"fmt"
	"io"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/html"
)

func Check(w io.Writer, url string) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		// Load the raw version of the store front
		doc, err := htmlquery.LoadURL(url)
		if err != nil {
			return err
		}

		// Extract embedded doc
		scriptNode, err := htmlquery.Query(doc, `//body/noscript`)
		if err != nil {
			return err
		}

		// Parse embedded doc
		reader := strings.NewReader(htmlquery.InnerText(scriptNode))
		scriptDoc, err := html.Parse(reader)
		if err != nil {
			return err
		}

		// Extract version from meta element
		metaNode, err := htmlquery.Query(scriptDoc, `//div/span/meta[@itemprop='version']`)
		var value string
		for _, att := range metaNode.Attr {
			if att.Key == "content" {
				value = att.Val
			}
		}

		fmt.Fprintln(w, value)

		return nil

	}
}
