package actions_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/osis/cwv/actions"
	"github.com/osis/cwv/util"
	"github.com/urfave/cli/v2"
)

var fakeZipCallCount = 0
var fakeZipCallArg string
var fakeZipper util.Zipper = func(path string) error {
	fakeZipCallArg = path
	fakeZipCallCount++
	return nil
}

var _ = Describe("Bundle", func() {
	It("calls the zip util with the correct param", func() {
		err := actions.Bundle(fakeZipper)(&cli.Context{})
		Expect(err).To(BeNil())

		Expect(fakeZipCallCount).To(Equal(1))
		Expect(fakeZipCallArg).Should(ContainSubstring("Library/Application Support/Google/Chrome/Default/Extensions/hnfanknocfeofbddgcijnmhnfnkdnaad"))
	})
})
