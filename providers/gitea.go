package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/drone/drone-go/drone"
)

// GiteaCommit represents a Gitea commit response
type GiteaCommit struct {
	Files []GiteaFile `json:"files"`
}

// GiteaFile represents a file in a Gitea commit
type GiteaFile struct {
	Filename string `json:"filename"`
}

// GiteaCompare represents a Gitea compare response
type GiteaCompare struct {
	Commits      []GiteaCommit `json:"commits"`
	TotalCommits int           `json:"total_commits"`
}

func GetGiteaFilesChanged(repo drone.Repo, build drone.Build, token string, giteaServer string) ([]string, error) {
	ctx := context.Background()

	// Ensure giteaServer has proper format
	if !strings.HasPrefix(giteaServer, "http://") && !strings.HasPrefix(giteaServer, "https://") {
		giteaServer = "https://" + giteaServer
	}
	giteaServer = strings.TrimSuffix(giteaServer, "/")

	var files []string
	var err error

	if build.Before == "" || build.Before == "0000000000000000000000000000000000000000" {
		// For new commits, get the commit details
		files, err = getGiteaCommitFiles(ctx, giteaServer, repo.Slug, build.After, token)
	} else {
		// For commit ranges, use compare API
		files, err = getGiteaCompareFiles(ctx, giteaServer, repo.Slug, build.Before, build.After, token)
	}

	return files, err
}

func getGiteaCommitFiles(ctx context.Context, server, repo, sha, token string) ([]string, error) {
	url := fmt.Sprintf("%s/api/v1/repos/%s/git/commits/%s", server, repo, sha)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gitea API returned status %d", resp.StatusCode)
	}

	var commit GiteaCommit
	if err := json.NewDecoder(resp.Body).Decode(&commit); err != nil {
		return nil, err
	}

	var files []string
	for _, file := range commit.Files {
		files = append(files, file.Filename)
	}

	return files, nil
}

func getGiteaCompareFiles(ctx context.Context, server, repo, base, head, token string) ([]string, error) {
	url := fmt.Sprintf("%s/api/v1/repos/%s/compare/%s...%s", server, repo, base, head)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gitea API returned status %d", resp.StatusCode)
	}

	var compare GiteaCompare
	if err := json.NewDecoder(resp.Body).Decode(&compare); err != nil {
		return nil, err
	}

	// Collect files from all commits in the comparison
	fileSet := make(map[string]bool)
	for _, commit := range compare.Commits {
		for _, file := range commit.Files {
			fileSet[file.Filename] = true
		}
	}

	// Convert set to slice
	var files []string
	for filename := range fileSet {
		files = append(files, filename)
	}

	return files, nil
}
