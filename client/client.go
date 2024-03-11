package client

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	Token      string
}

type Response struct {
	OnCallNow []string `json:"on_call_now"`
	Username  string   `json:"username"`
}

func New() *Client {
	return &Client{
		HTTPClient: &http.Client{},
	}
}

func (api *Client) Request(method, url string, body io.Reader, headers map[string]string) *Response {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		panic(err)
	}

	for header, value := range headers {
		request.Header.Set(header, value)
	}

	response, err := api.HTTPClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	var responseJson Response

	if err := json.Unmarshal(responseBody, &responseJson); err != nil {
		log.Warn().Str("host", request.Host).Msg(err.Error())
		log.Warn().Str("status", response.Status).Msg(err.Error())
		log.Warn().Msg("if the status is 2xx, then everything is ok")
	}

	return &responseJson
}
