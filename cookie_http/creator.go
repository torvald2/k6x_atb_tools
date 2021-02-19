package cookie_http

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
)

var once sync.Once

type HttpClientCreator struct{}

func (creator *HttpClientCreator) New(stringCookies map[string]string, stringUrl string) httpClient {
	once.Do(func() {
		var cookies []*http.Cookie

		siteUrl, err := url.Parse(stringUrl)
		if err != nil {
			panic("BAD URL FORMAT")
		}
		jar, err := cookiejar.New(nil)
		if err != nil {
			panic("COOKIEJAR CREATION ERROR")
		}
		for name, value := range stringCookies {
			cookie := http.Cookie{Name: name, Value: value}
			cookies = append(cookies, &cookie)
		}
		jar.SetCookies(siteUrl, cookies)

		thisClient = httpClient{&http.Client{Jar: jar}}
	})
	return thisClient
}
