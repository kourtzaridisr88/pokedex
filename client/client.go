package client

import (
	"net/http"
)

const baseEndpoint string = "https://pokeapi.co/api/v2/"

func GetEndpoint(path string) (*http.Response, error) {
	fullUrl := baseEndpoint + path

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
