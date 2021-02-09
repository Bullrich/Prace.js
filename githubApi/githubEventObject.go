package githubApi

type GitHubEvent struct {
	Action      string          `json:"action"`
	Number      int             `json:"number"`
	PullRequest PullRequestData `json:"pull_request"`
}

type PullRequestData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Head  struct {
		Ref string `json:"ref"`
	} `json:"head"`
	Labels []struct {
		Id          int64  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"labels"`
	RequestedReviewers []user `json:"requested_reviewers"`
	RequestedTeams     []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"requested_teams"`
	Additions    int    `json:"additions"`
	Deletions    int    `json:"deletions"`
	State        string `json:"state"`
	ChangedFiles int    `json:"changed_files"`
	User         user   `json:"user"`
}

type user struct {
	Login string `json:"login"`
	Id    int64  `json:"id"`
}

type Reviewer struct {
	State string `json:"state"`
	User  *struct {
		Login string `json:"login"`
		Id    int    `json:"id"`
	} `json:"user"`
}
