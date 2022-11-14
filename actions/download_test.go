package actions_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/osis/cwv/actions"
	"github.com/urfave/cli/v2"
)

var _ = Describe("Download", func() {
	var server *httptest.Server
	var fileUrl string

	BeforeEach(func() {
		downloadName := "extension_1_0_1_0.crx"
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			htmlBytes, err := os.ReadFile("../fake.html")
			Expect(err).To(BeNil())

			w.Write(htmlBytes)
		}))

		var err error
		fileUrl, err = url.JoinPath(server.URL, downloadName)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		server.Close()
		zipPaths, _ := filepath.Glob("*.zip")
		for _, zipPath := range zipPaths {
			os.RemoveAll(zipPath)
		}
	})

	It("downloads a chrome extension from the supplied url", func() {
		err := actions.Download(fileUrl)(&cli.Context{})
		Expect(err).To(BeNil())

		bytes, err := os.ReadFile("coinbase-wallet-chrome-1.0.1.zip")
		Expect(err).To(BeNil())
		Expect(string(bytes)).Should(ContainSubstring("Coinbase Wallet extension - Chrome Web Store"))
	})
})
