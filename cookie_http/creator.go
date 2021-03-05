package cookie_http

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type HttpClientCreator struct{}

func (creator *HttpClientCreator) New(stringCookies map[string]string, stringUrl string) *httpClient {
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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//ignore tls

	return &httpClient{&http.Client{Jar: jar, Transport: tr}}
}
