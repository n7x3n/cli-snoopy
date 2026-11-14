package selfmanage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const url = "https://api.github.com/repos/n7x3n/cli-snoopy/releases/latest"
const CurrentVersion = "v0.1.0"

type githubRelease struct {
	TagName string `json:"tag_name"`
}

func VersionCheck() (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Set("User-Agent", "cli-snoopy")
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Couldn't establish connection")
		return "", err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var release githubRelease
	err = json.Unmarshal(bodyBytes, &release)
	if err != nil {
		return "", err
	}
	return release.TagName, nil
}
