package selfmanage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const url = "https://api.github.com/repos/n7x3n/cli-snoopy/releases/latest"
const CurrentVersion = "v0.2.0"

type GithubRelease struct {
	TagName string        `json:"tag_name"`
	Asset   []githubAsset `json:"assets"`
}
type githubAsset struct {
	Filename string `json:"name"`
	Url      string `json:"browser_download_url"`
}

func VersionCheck() (*GithubRelease, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "cli-snoopy")
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var githubCall GithubRelease
	err = json.Unmarshal(bodyBytes, &githubCall)
	if err != nil {
		return nil, err
	}
	return &githubCall, nil
}
func Updater(githubData *GithubRelease) {
	exepath, err := os.Executable()
	if err != nil {
		println("")
		return
	}
	path, err := filepath.EvalSymlinks(exepath)
	system := runtime.GOOS
	arch := runtime.GOARCH
	for _, asset := range githubData.Asset {
		if strings.Contains(asset.Filename, system) && strings.Contains(asset.Filename, arch) {
			req, err := http.NewRequest(http.MethodGet, asset.Url, nil)
			if err != nil {
				fmt.Printf("Error encountered: %v", err)
				return
			}
			req.Header.Set("User-Agent", "cli-snoopy")
			client := &http.Client{
				Timeout: 30 * time.Second,
			}
			fmt.Println("Downloading newer version")
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Download failed. Please, try again later")
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				fmt.Println("Couldn't establish connection")
				return
			}
			exeName := fmt.Sprintf("%s.new", path)
			app, err := os.Create(exeName)
			if err != nil {
				return
			}
			_, err = io.Copy(app, resp.Body)
			if err != nil {
				fmt.Println("Download interrupted, please try again.")
				app.Close()
				os.Remove(exeName)
				return
			}
			app.Close()
			if system == "linux" || system == "darwin" {
				os.Chmod(exeName, 0755)
				os.Rename(exeName, path)
				fmt.Println("Snoopy has been successfully updated")
				return
			}
			if system == "windows" {
				oldExeName := fmt.Sprintf("%s.old", path)
				os.Rename(path, oldExeName)
				os.Rename(exeName, path)
				fmt.Println("Snoopy has been successfuly updated")
				return
			}
		}
	}
	fmt.Println("Sorry, I couldn't find update for your system")
}

func WinOldPurger() {
	exepath, err := os.Executable()
	if err != nil {
		return
	}
	path, err := filepath.EvalSymlinks(exepath)
	if err != nil {
		return
	}
	oldExePath := fmt.Sprintf("%s.old", path)
	os.Remove(oldExePath)
}
