package cookie_http

import (
	"io"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

var thisClient httpClient

type httpClient struct {
	Client *http.Client
}

func (client *httpClient) getRequest(method string, url string, data io.Reader, stat *httpstat.Result) (*http.Request, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return req, err
	}
	context := httpstat.WithHTTPStat(req.Context(), stat)
	req = req.WithContext(context)
	return req, nil

}

//get env
func (client *httpClient) Get(url string) measuresResponse {
	var results httpstat.Result
	req, err := client.getRequest("GET", url, nil, &results)
	if err != nil {
		panic(err.Error())
	}
	res, err := client.Client.Do(req)
	if err != nil {
		panic(err.Error())

	}
	response_time := time.Now()
	response, err := createResponse(res, results, response_time)
	if err != nil {
		panic(err.Error())
	}
	return response

}
