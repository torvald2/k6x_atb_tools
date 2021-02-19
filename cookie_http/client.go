package cookie_http

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

var thisClient *http.Client

func setCookies(stringCookies map[string]string, stringUrl string) {
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

	thisClient = &http.Client{
		Jar: jar,
	}
}
