package selfmanage

import (
	"flag"
	"fmt"

	"github.com/n7x3n/cli-snoopy/docs"
)

func HandleFlags() bool {
	updateFlag := flag.Bool("update", false, "update snoopy to latest release")
	versionFlag := flag.Bool("version", false, "print version of snoopy")
	docsFlag := flag.Bool("docs", false, "display docs for cli-snoopy")
	flag.Parse()
	if *updateFlag {
		githubData, err := VersionCheck()
		if err == nil {
			if githubData.TagName != CurrentVersion {
				Updater(githubData)
			} else {
				fmt.Println("Already up to date")
			}
		}
		return true
	}
	if *versionFlag {
		fmt.Printf("cli-snoopy version: %s\n", CurrentVersion)
		return true
	}
	if *docsFlag {
		if err := docs.DocsServer(); err != nil {
			fmt.Println("Failed to launch docs web server")
		}
		return true
	}
	return false
}
