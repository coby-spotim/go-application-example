package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/coby-spotim/go-application-example/pkg/types"
)

type ApiClient struct {
	client http.Client
	url    string
}

func NewApiClient(url string) *ApiClient {
	return &ApiClient{
		url: url,
		client: http.Client{
			Timeout:   time.Second * 10,
			Transport: http.DefaultTransport,
		},
	}
}

func (c *ApiClient) Ping(headers map[string]string) error {
	u, _ := url.Parse(c.url + "/ping")

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	addHeaders(headers, req)

	r, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	if r.StatusCode >= 300 {
		err := fmt.Errorf("%s returned status %v", u.String(), r.StatusCode)
		return err
	}

	return nil
}

func (c *ApiClient) GetUser(name string, headers map[string]string) (*types.Person, error) {
	u, _ := url.Parse(c.url + fmt.Sprintf("/user/%s", name))

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	addHeaders(headers, req)

	r, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()
	if r.StatusCode >= 300 {
		err := fmt.Errorf("%s returned status %v", u.String(), r.StatusCode)
		return nil, err
	}

	var person types.Person
	if err = json.NewDecoder(r.Body).Decode(&person); err != nil {
		return nil, err
	}

	return &person, nil
}

func (c *ApiClient) SetUser(person types.Person, headers map[string]string) error {
	u, _ := url.Parse(c.url + "/user")

	buff := new(bytes.Buffer)
	err := json.NewEncoder(buff).Encode(person)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), buff)
	if err != nil {
		return err
	}

	addHeaders(headers, req)

	r, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	if r.StatusCode >= 300 {
		err := fmt.Errorf("%s returned status %v", u.String(), r.StatusCode)
		return err
	}

	return nil
}

func addContentTypeHeader(req *http.Request) {
	if req.Header.Get("content-type") == "" {
		req.Header.Set("content-type", "application/json")
	}
}

func addHeaders(headers map[string]string, req *http.Request) {
	defer addContentTypeHeader(req)

	if headers == nil {
		return
	}

	for h, v := range headers {
		req.Header.Set(h, v)
	}
}
