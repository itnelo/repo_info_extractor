package languages_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"

	"github.com/codersrank-org/repo_info_extractor/librarydetection/languages"
)

var _ = Describe("TypeScriptLibraryDetection", func() {
	fixture, err := ioutil.ReadFile("./fixtures/typescript.fixture")
	if err != nil {
		panic(err)
	}

	expectedLibraries := []string{
		"lib1",
		"lib2",
		"lib3",
		"lib4",
	}

	analyzer := languages.NewTypeScriptAnalyzer()

	Describe("Extract TypeScript Libraries", func() {
		It("Should be able to extract libraries", func() {
			libs := analyzer.ExtractLibraries(string(fixture))
			assertSameUnordered(libs, expectedLibraries)
		})
	})
})
