package cookie_http

import (
	"testing"
)

func TestHttpClientCreator_New(t *testing.T) {
	creator := HttpClientCreator{}
	client := creator.Create(map[string]string{"foo": "bar"}, "https://google.com")
	res := client.Get("https://google.com")
	if res.Status != 200 {
		t.Errorf("Status not 200")
	}
}
