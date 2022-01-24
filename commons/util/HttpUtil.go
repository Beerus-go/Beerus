package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var client = new(http.Client)

// RequestBody
// Initiating a request with body parameters
func RequestBody(reqUrl string, method string, header map[string]string, reqBody io.Reader) (*http.Response, error) {
	if strings.ToUpper(method) == "GET" {
		return nil, errors.New("")
	}

	req, err := http.NewRequest(method, reqUrl, reqBody)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for key, val := range header {
			req.Header.Set(key, val)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get
// Sending get requests
func Get(reqUrl string, param interface{}) (*http.Response, error) {

	params := url.Values{}
	URL, err := url.Parse(reqUrl)
	if err != nil {
		return nil, err
	}

	if param != nil {
		jsonByte, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}

		data := make(map[string]interface{})

		err = json.Unmarshal(jsonByte, data)
		if err != nil {
			return nil, err
		}

		if len(data) > 0 {
			for key, val := range data {
				params.Set(key, ToString(val))
			}
		}
	}

	URL.RawQuery = params.Encode()
	resp, err := http.Get(URL.String())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
