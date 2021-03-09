package cookie_http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

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

func (client *httpClient) postData(url string, data io.Reader) measuresResponse {
	var results httpstat.Result
	req, err := client.getRequest("POST", url, data, &results)
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

func (client *httpClient) PostFormData(url string, data string) measuresResponse {
	var parsedData map[string]string
	buf := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(buf)
	defer multipartWriter.Close()

	byte_data := []byte(data)
	if err := json.Unmarshal(byte_data, &parsedData); err != nil {
		panic(err.Error())
	}

	for k, v := range parsedData {
		field, err := multipartWriter.CreateFormField(k)
		if err != nil {
			panic(err.Error())
		}
		_, err = field.Write([]byte(v))
		if err != nil {
			panic(err.Error())
		}
	}
	response := client.postData(url, buf)
	return response

}
