package selfmanage

import (
	"flag"
	"fmt"
)

func HandleFlags() bool {
	updateFlag := flag.Bool("update", false, "update snoopy to latest release")
	versionFlag := flag.Bool("version", false, "print version of snoopy")
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
		fmt.Printf("cli-snoopy version %s", CurrentVersion)
		return true
	}
	return false
}
