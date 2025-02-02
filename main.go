package main

import (
	"flag"
	"strings"

	"github.com/codersrank-org/repo_info_extractor/extractor"
)

func main() {

	repoPath := flag.String("repo_path", "", "Path of the repo")
	// Following two flags should be used to disable email prompt
	// Program is going to ask you to choose your emails
	// But if you want, you can provide the emails yourself
	headless := flag.String("headless", "false", "Should run on headless mode?")
	obfuscate := flag.String("obfuscate", "true", "Should obfuscate the result?")
	outputPath := flag.String("output_path", "./repo_data_v2", "Where to put output file")
	gitPath := flag.String("git_path", "", "Where is git binary?")
	emailString := flag.String("emails", "", "Predefined emails") // Example: "alim.giray@codersrank.io,alimgiray@gmail.com"
	seeds := flag.String("seeds", "", "Seed")                     // Example: "alimgiray, alimgiray@codersrank.io"
	flag.Parse()

	if repoPath == nil || *repoPath == "" {
		panic("Please provide a path to the repo")
	}

	emails := make([]string, 0)
	if emailString != nil && len(*emailString) > 0 {
		emails = strings.Split(*emailString, ",")
	}

	seed := make([]string, 0)
	if seeds != nil && len(*seeds) > 0 {
		seed = strings.Split(*seeds, ",")
	}

	repoExtractor := extractor.RepoExtractor{
		RepoPath:   *repoPath,
		OutputPath: *outputPath,
		GitPath:    *gitPath,
		Headless:   *headless == "true",
		Obfuscate:  *obfuscate == "true",
		UserEmails: emails,
		Seed:       seed,
	}

	err := repoExtractor.Extract()
	if err != nil {
		panic(err)
	}
}
