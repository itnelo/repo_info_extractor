package languages_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"

	"github.com/codersrank-org/repo_info_extractor/librarydetection/languages"
)

var _ = Describe("GoLibraryDetection", func() {
	fixture, err := ioutil.ReadFile("./fixtures/go.fixture")
	if err != nil {
		panic(err)
	}

	expectedLibraries := []string{
		"fmt",
		"library1",
		"gitlab.com/username/reponame/library2",
		"gitlab.com/username/library3",
		"gitlab.com/username/reponame/library4",
		"gitlab.com/username/library5",
		"gitlab.com/username/reponame/library6",
		"gitlab.com/username/library7",
		"gitlab.com/username/reponame/library8",
		"gitlab.com/username/library9",
		"library10",
		"gitlab.com/username/reponame/library11",
		"gitlab.com/username/library12",
		"gitlab.com/username/reponame/library13",
		"gitlab.com/username/library14",
		"gitlab.com/username/reponame/library15",
		"gitlab.com/username/library16",
		"gitlab.com/username/reponame/library17",
		"gitlab.com/username/library18",
	}

	analyzer := languages.NewGoAnalyzer()

	Describe("Extract Go sLibraries", func() {
		It("Should be able to extract libraries", func() {
			libs := analyzer.ExtractLibraries(string(fixture))
			Expect(len(libs)).Should(Equal(len(expectedLibraries)))
			for _, v := range libs {
				Expect(v).Should(BeElementOf(expectedLibraries))
			}
		})
	})
})
