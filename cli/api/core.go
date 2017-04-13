package api

type status struct {
	responseBase
}

// Status checks the status of the api
func Status() (bool, error) {
	resp := status{}

	err := get("core/status", nil, &resp)
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}
