package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"gitup.io/isaac/gitup/cli/config"
)

var client = http.Client{Timeout: 10 * time.Second}

func buildURL(endpoint string) string {
	return fmt.Sprintf("%s/api/%s", config.Host, endpoint)
}

func buildURLWithQuery(endpoint string, p params) string {
	return fmt.Sprintf("%s%s", buildURL(endpoint), p.buildQuery())
}

func get(route string, p params, target Response) error {
	path := buildURLWithQuery(route, p)

	resp, err := client.Get(path)

	target.SetStatus(resp.StatusCode)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func post(route string, p params, target Response) error {
	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(p)

	resp, err := client.Post(buildURL(route), "application/json; charset=utf-8", b)

	target.SetStatus(resp.StatusCode)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

type params map[string]string

func (p params) buildQuery() string {
	if p == nil {
		return ""
	}

	queryString := ""

	for key, value := range p {
		queryString = fmt.Sprintf("%s&%s=%s", queryString, key, url.QueryEscape(value))
	}

	queryString = queryString[1:]
	queryString = fmt.Sprintf("?%s", queryString)

	return queryString
}
