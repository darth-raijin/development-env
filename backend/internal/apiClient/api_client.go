package apiclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Get(url string, headers map[string]string) (*http.Response, error) {
	return request("GET", url, headers, nil)
}

func Post(url string, headers map[string]string, body interface{}) (*http.Response, error) {
	return request("POST", url, headers, body)
}

func Put(url string, headers map[string]string, body interface{}) (*http.Response, error) {
	return request("PUT", url, headers, body)
}

func request(method, url string, headers map[string]string, body interface{}) (*http.Response, error) {
	client := http.Client{}

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// REmember close body
	return client.Do(req)
}
