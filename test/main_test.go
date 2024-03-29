package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type header struct {
	Key   string
	Value string
}

var (
	localhost = "http://localhost:8001"
	testHost  = ""
)

func PerformRequest(method, path string, req, res interface{}, headers ...header) (*http.Response, error) {
	body, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Hel")
		return nil, err
	}

	client := &http.Client{}

	request, err := http.NewRequestWithContext(context.Background(), method, fmt.Sprintf("%s%s", localhost, path), bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Word")
		return nil, err
	}

	for _, h := range headers {
		request.Header.Add(h.Key, h.Value)
	}

	request.Header.Add("Accept", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("err")
		return resp, err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(resp_body, &res)
	if err != nil {
		fmt.Println("Errror:", err)
		return nil, err
	}

	return resp, nil
}
