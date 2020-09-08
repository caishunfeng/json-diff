package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//callApi - make a http request to the end points
//It need not be http request, it can be any external api call (grpc, graphql, etc.)
func callAPI(endpoint string, queryParams map[string]string, headerParams map[string]string) (string, error) {
	client := http.Client{}

	url := endpoint

	//Add query params
	qparams := make([]string, 0)
	for k, v := range queryParams {
		qparams = append(qparams, k+"="+v)
	}
	if len(qparams) > 0 {
		url += "?" + strings.Join(qparams, "&")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request %v", err)
	}

	for k, v := range headerParams {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to request api %v", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	return string(b), nil
}
