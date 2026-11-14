package selfmanage

import (
	"flag"
	"fmt"
)

func HandleFlags() bool {
	installFlag := flag.Bool("install", false, "install snoopy to system path")
	removeFlag := flag.Bool("remove", false, "remove snoopy from system path")
	updateFlag := flag.Bool("update", false, "update snoopy to latest release")
	versionFlag := flag.Bool("version", false, "print version of snoopy")
	flag.Parse()
	if *installFlag {
		println("Installing Snoopy to system path")
		return true
	}
	if *removeFlag {
		println("Removing snoopy =(")
		return true
	}
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
