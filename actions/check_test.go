package actions_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tenkeylabs/cwv/actions"
	"github.com/urfave/cli/v2"
)

var _ = Describe("Bundle", func() {
	var server *httptest.Server

	BeforeEach(func() {
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			htmlBytes, err := os.ReadFile("../fake.html")
			Expect(err).To(BeNil())

			w.Write(htmlBytes)
		}))
	})

	AfterEach(func() {
		server.Close()
	})
	It("prints the latest version number from the chrome web store", func() {
		r, w, _ := os.Pipe()
		err := actions.Check(w, server.URL)(&cli.Context{})
		w.Close()
		Expect(err).To(BeNil())

		outBytes, _ := ioutil.ReadAll(r)
		Expect(string(outBytes)).To(Equal("1.0.1\n"))
	})
})
