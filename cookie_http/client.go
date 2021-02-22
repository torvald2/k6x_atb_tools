package cookie_http

import (
	"io"
	"log"
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
	response_time := time.Now()

	if err != nil {
		log.Print(err.Error())
		return createErrorResponse(results, response_time)
	}
	return createResponse(res, results, response_time)

}
