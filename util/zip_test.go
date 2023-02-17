package util_test

import (
	"archive/zip"
	"log"
	"os"
	"path"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tenkeylabs/cwv/util"
)

var testBasePath = "../tmp"
var versionPath = path.Join(testBasePath, "1.0.0")
var version2Path = path.Join(testBasePath, "2.0.0")

var dirPaths = []string{path.Join(versionPath), version2Path}
var filePaths = []string{
	path.Join(versionPath, "code.js"),
	path.Join(versionPath, "assets", "image.png"),
	path.Join(version2Path, "code2.js"),
	path.Join(version2Path, "assets", "image2.png"),
}

func setupTestStructure() {
	for _, dirPath := range dirPaths {
		err := os.MkdirAll(path.Join(dirPath, "assets"), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, filepath := range filePaths {
		f, err := os.Create(filepath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
}

var _ = AfterSuite(func() {
	os.RemoveAll(testBasePath)
	os.RemoveAll("*.zip")
})

var _ = Describe("ZipDirs", func() {
	expectedZips := make(map[string][]string)

	BeforeEach(func() {
		setupTestStructure()
	})

	AfterEach(func() {
		os.RemoveAll(testBasePath)

		zipPaths, _ := filepath.Glob("*.zip")
		for _, zipPath := range zipPaths {
			os.RemoveAll(zipPath)
		}
	})

	It("zips up a directories at a path", func() {

		expectedZips["coinbase-wallet-chrome-1.0.0.zip"] = []string{
			"assets/image.png",
			"code.js",
		}

		expectedZips["coinbase-wallet-chrome-2.0.0.zip"] = []string{
			"assets/image2.png",
			"code2.js",
		}

		util.ZipDirs(testBasePath)

		for zipPath := range expectedZips {
			_, err := os.Stat(zipPath)
			Expect(err).To(BeNil())
		}

		for zipPath, zipContentPaths := range expectedZips {
			zipFile, err := zip.OpenReader(zipPath)
			Expect(err).To(BeNil())

			for i, zipContentPath := range zipContentPaths {
				Expect(zipFile.File[i].Name).To(Equal(zipContentPath))
			}
		}
	})
})
