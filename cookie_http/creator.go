package cookie_http

import (
	"net/http"
	"sync"
)

var once sync.Once

type HttpClientCreator struct{}

func (creator *HttpClientCreator) New(cookies map[string]string, url string) *http.Client {
	once.Do(func() { setCookies(cookies, url) })
	return thisClient
}
