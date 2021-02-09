package evaluator

type PraceConfig struct {
	Title     *pattern `json:"title"`
	Body      *pattern `json:"body"`
	Branch    *pattern `json:"branch"`
	Reviewers *struct {
		Minimum int      `json:"minimum"`
		Users   []string `json:"users"`
		Teams   []string `json:"teams"`
	} `json:"reviewers"`
	Additions int      `json:"additions"`
	Labels    []string `json:"labels"`
}

type pattern struct {
	Patterns []string `json:"patterns"`
	Error    string   `json:"error"`
}
