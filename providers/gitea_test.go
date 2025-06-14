package providers

import (
	"sort"
	"testing"

	"github.com/h2non/gock"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/converter"

	"github.com/google/go-cmp/cmp"
)

func TestGetGiteaFilesChangedCommit(t *testing.T) {
	defer gock.Off()

	gock.New("https://gitea.example.com").
		Get("/api/v1/repos/testuser/drone-yml-test/git/commits/537575f44a09c57dfc472e26fe067754fd2f9374").
		Reply(200).
		Type("application/json").
		JSON(map[string]interface{}{
			"files": []map[string]interface{}{
				{"filename": ".drone.yml"},
			},
		})

	req := &converter.Request{
		Build: drone.Build{
			Before: "",
			After:  "537575f44a09c57dfc472e26fe067754fd2f9374",
		},
		Repo: drone.Repo{
			Namespace: "testuser",
			Name:      "drone-yml-test",
			Slug:      "testuser/drone-yml-test",
			Config:    ".drone.yml",
		},
	}

	got, err := GetGiteaFilesChanged(req.Repo, req.Build, "validtoken", "https://gitea.example.com")
	if err != nil {
		t.Error(err)
		return
	}

	want := []string{".drone.yml"}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}
}

func TestGetGiteaFilesChangedCompare(t *testing.T) {
	defer gock.Off()

	gock.New("https://gitea.example.com").
		Get("/api/v1/repos/testuser/drone-yml-test/compare/e3c0ff4d5cef439ea11b30866fb1ed79b420801d...537575f44a09c57dfc472e26fe067754fd2f9374").
		Reply(200).
		Type("application/json").
		JSON(map[string]interface{}{
			"commits": []map[string]interface{}{
				{
					"files": []map[string]interface{}{
						{"filename": ".drone.yml"},
						{"filename": "README.md"},
					},
				},
			},
			"total_commits": 1,
		})

	req := &converter.Request{
		Build: drone.Build{
			Before: "e3c0ff4d5cef439ea11b30866fb1ed79b420801d",
			After:  "537575f44a09c57dfc472e26fe067754fd2f9374",
		},
		Repo: drone.Repo{
			Namespace: "testuser",
			Name:      "drone-yml-test",
			Slug:      "testuser/drone-yml-test",
			Config:    ".drone.yml",
		},
	}

	got, err := GetGiteaFilesChanged(req.Repo, req.Build, "validtoken", "https://gitea.example.com")
	if err != nil {
		t.Error(err)
		return
	}

	want := []string{".drone.yml", "README.md"}

	// Sort both slices to ensure consistent comparison
	sort.Strings(got)
	sort.Strings(want)

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}
}

func TestGetGiteaFilesChangedServerFormat(t *testing.T) {
	defer gock.Off()

	// Test that server URL formatting works correctly
	gock.New("https://gitea.example.com").
		Get("/api/v1/repos/testuser/drone-yml-test/git/commits/537575f44a09c57dfc472e26fe067754fd2f9374").
		Reply(200).
		Type("application/json").
		JSON(map[string]interface{}{
			"files": []map[string]interface{}{
				{"filename": "test.txt"},
			},
		})

	req := &converter.Request{
		Build: drone.Build{
			Before: "",
			After:  "537575f44a09c57dfc472e26fe067754fd2f9374",
		},
		Repo: drone.Repo{
			Namespace: "testuser",
			Name:      "drone-yml-test",
			Slug:      "testuser/drone-yml-test",
			Config:    ".drone.yml",
		},
	}

	// Test with server URL without protocol - should add https://
	got, err := GetGiteaFilesChanged(req.Repo, req.Build, "validtoken", "gitea.example.com")
	if err != nil {
		t.Error(err)
		return
	}

	want := []string{"test.txt"}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}
}
