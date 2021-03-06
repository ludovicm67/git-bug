package gitlab

import (
	"net/http"
	"time"

	"github.com/xanzy/go-gitlab"

	"github.com/MichaelMure/git-bug/bridge/core"
	"github.com/MichaelMure/git-bug/bridge/core/auth"
)

const (
	target = "gitlab"

	metaKeyGitlabId      = "gitlab-id"
	metaKeyGitlabUrl     = "gitlab-url"
	metaKeyGitlabLogin   = "gitlab-login"
	metaKeyGitlabProject = "gitlab-project-id"
	metaKeyGitlabBaseUrl = "gitlab-base-url"

	keyProjectID     = "project-id"
	keyGitlabBaseUrl = "base-url"

	defaultBaseURL = "https://gitlab.com/"
	defaultTimeout = 60 * time.Second
)

type Gitlab struct{}

func (*Gitlab) Target() string {
	return target
}

func (*Gitlab) NewImporter() core.Importer {
	return &gitlabImporter{}
}

func (*Gitlab) NewExporter() core.Exporter {
	return &gitlabExporter{}
}

func buildClient(baseURL string, token *auth.Token) (*gitlab.Client, error) {
	httpClient := &http.Client{
		Timeout: defaultTimeout,
	}

	gitlabClient := gitlab.NewClient(httpClient, token.Value)
	err := gitlabClient.SetBaseURL(baseURL)
	if err != nil {
		return nil, err
	}

	return gitlabClient, nil
}
