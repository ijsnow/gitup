package api

import "gitup.io/isaac/gitup/types"

// RepoResp is the response type for a repo endpoint
type RepoResp struct {
	responseBase
	types.Repo
}

// CreateRemoteRepo attempts to create a user
func CreateRemoteRepo(r *types.Repo) (bool, RepoResp) {
	resp := RepoResp{}

	p := params{
		"name": r.Name,
	}

	err := authPost("repo/", p, &resp)
	if err != nil {
		return false, resp
	}

	return resp.Success, resp
}

// DeleteRepo attempts to create a user
func DeleteRepo(r *types.Repo) (bool, RepoResp) {
	resp := RepoResp{}

	p := params{
		"name": r.Name,
	}

	err := authDelete("repo/", p, &resp)
	if err != nil {
		return false, resp
	}

	return resp.Success, resp
}
