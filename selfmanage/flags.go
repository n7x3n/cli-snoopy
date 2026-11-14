package selfmanage

import (
	"flag"
	"fmt"
)

func HandleFlags() bool {
	installFlag := flag.Bool("install", false, "install snoopy to system path")
	removeFlag := flag.Bool("remove", false, "remove snoopy from system path")
	updateFlag := flag.Bool("update", false, "update snoopy to latest release")
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
		latestTag, err := VersionCheck()
		if err == nil {
			if latestTag != CurrentVersion {
				fmt.Println("Pulling new version")
			} else {
				println("Already up to date")
			}
		}
		return true
	}
	return false
}
