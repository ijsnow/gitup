package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ijsnow/gitup/cli/config"
)

var client = http.Client{Timeout: 10 * time.Second}

func buildURL(endpoint string) string {
	return fmt.Sprintf("%s/api/%s", config.Host, endpoint)
}

func buildURLWithQuery(endpoint string, p params) string {
	return fmt.Sprintf("%s%s", buildURL(endpoint), p.buildQuery())
}

func request(req *http.Request, headers [][]string) (*http.Response, error) {
	for _, h := range headers {
		req.Header.Add(h[0], h[1])
	}

	return client.Do(req)
}

func get(route string, p params, target Response, headers ...[]string) error {
	path := buildURLWithQuery(route, p)

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	resp, err := request(req, headers)

	target.SetStatus(resp.StatusCode)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func authGet(route string, p params, target Response) error {
	return get(route, p, target, []string{"Authorization", config.Token})
}

func post(route string, p params, target Response, headers ...[]string) error {
	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(p)

	req, err := http.NewRequest("POST", buildURL(route), b)
	if err != nil {
		return err
	}

	resp, err := request(req, headers)

	target.SetStatus(resp.StatusCode)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func authPost(route string, p params, target Response) error {
	return post(route, p, target, []string{"Authorization", config.Token})
}

func delete(route string, p params, target Response, headers ...[]string) error {
	path := buildURLWithQuery(route, p)

	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	resp, err := request(req, headers)

	target.SetStatus(resp.StatusCode)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func authDelete(route string, p params, target Response) error {
	return delete(route, p, target, []string{"Authorization", config.Token})
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
