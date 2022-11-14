package actions_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/osis/cwv/actions"
	"github.com/urfave/cli/v2"
)

var _ = Describe("Version", func() {
	It("prints the current version", func() {
		r, w, _ := os.Pipe()
		err := actions.Version(w)(&cli.Context{})
		w.Close()
		Expect(err).To(BeNil())

		outBytes, _ := ioutil.ReadAll(r)
		Expect(string(outBytes)).To(Equal("0.0.1\n"))
	})
})
