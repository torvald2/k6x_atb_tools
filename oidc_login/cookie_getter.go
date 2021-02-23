package oidc_login

var thisCookies cookies

type cookies struct {
	data map[string]string
}

func (cookies *cookies) Get() map[string]string {
	return cookies.data
}
