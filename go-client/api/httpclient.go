package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpResponse struct {
	data []byte
	err  error
}

func (hr *HttpResponse) HasError() bool {
	return hr.err != nil
}
func (hr *HttpResponse) ToData(holder interface{}) error {
	if hr.err != nil {
		return hr.err
	}
	buff := bytes.NewBuffer(hr.data)
	json.NewDecoder(buff).Decode(holder)
	return nil
}

type HttpClient interface {
	Get(uri string) HttpResponse
}

type _httpClient struct {
	server string
	client *http.Client
}

func (hc *_httpClient) Get(uri string) HttpResponse {
	res, err := hc.client.Get(fmt.Sprintf("%s/%s", hc.server, uri))
	if err != nil {
		return HttpResponse{
			err: err,
		}
	} else {
		defer res.Body.Close()
		b, berr := ioutil.ReadAll(res.Body)
		if res.StatusCode >= 400 {
			berr = errors.New(fmt.Sprintf("api call returned %d", res.StatusCode))
		}
		return HttpResponse{
			data: b,
			err:  berr,
		}
	}
}

func newHttpClient(url string) HttpClient {
	return &_httpClient{
		server: url,
		client: &http.Client{
			Timeout: 2000 * time.Millisecond,
		},
	}
}
