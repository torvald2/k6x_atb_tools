package cookie_http

import (
	"testing"
)

func TestHttpClientCreator_New(t *testing.T) {
	creator := HttpClientCreator{}
	client := creator.New(map[string]string{"foo": "bar"}, "https://google.com")
	res, err := client.Get("https://google.com")
	if err != nil {
		t.Errorf(err.Error())
	}
	var isFooCookie bool
	cookies := res.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "foo" {
			break
		} else {
			isFooCookie = true
		}
	}
	if !isFooCookie {
		t.Errorf("No foo cookie")
	}
}
