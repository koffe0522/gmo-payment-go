package httpc

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// ResponseBody is
type ResponseBody struct {
	StatusCode int
	Body       []byte
}

// Client is
type Client struct {
	BaseURL string
}

// IClinet is
type IClinet interface {
	Do(path string, paramsJSON []byte) (*ResponseBody, error)
}

// Do is
func (c *Client) Do(path string, paramsJSON []byte) (*ResponseBody, error) {
	req, err := http.NewRequest("POST", c.BaseURL+path, bytes.NewBuffer(paramsJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("User-Agent", "GMO PG Client: Unofficial")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &ResponseBody{
		Body:       body,
		StatusCode: res.StatusCode,
	}, err
}
