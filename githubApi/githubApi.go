package githubApi

import (
	"context"
	"errors"
	"github.com/Bullrich/prace.go/evaluator"
	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

type GitApi interface {
	getConfig(branch string) *evaluator.PraceConfig
	reportFailed(message string)
	getReviewers() []Reviewer
	setResult()
	log(message string)
}

type actionClient struct {
	Client *github.Client
	Owner  string
	Repo   string
	WorkflowId string
}

func BuildClient(ctx context.Context, githubToken string) (*actionClient, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repositoryData := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
	if len(repositoryData) < 2 {
		return nil, errors.New("'GITHUB_REPOSITORY' env var not set correctly")
	}
	owner, repo := repositoryData[0], repositoryData[1]

	workflowId:= os.Getenv("GITHUB_WORKFLOW");

	return &actionClient{Client: client, Owner: owner, Repo: repo, WorkflowId: workflowId}, nil
}

func (c *actionClient) getConfig(branch string) *evaluator.PraceConfig{
	return nil
}