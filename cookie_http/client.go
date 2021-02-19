package cookie_http

import (
	"net/http"
)

var thisClient httpClient

type httpClient struct {
	Client *http.Client
}

//get env
func (client *httpClient) Get(url string) int {
	res, _ := client.Client.Get(url)
	return res.StatusCode

}
