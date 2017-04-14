package gittp

// RequestInfo is used for checking if we should serve or not
type RequestInfo struct {
	RepoOwner string
	RepoName  string

	Username string
	Password string
}
