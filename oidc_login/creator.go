package oidc_login

import (
	"sync"
)

var once sync.Once

type CookiesCreator struct{}

func (creator *CookiesCreator) New(server_url, target_url, login, password, login_selector, password_selector, submit_selector string) *cookies {
	once.Do(func() {
		page, err := createPage(server_url, target_url)
		if err != nil {
			panic(err.Error())
		}
		defer page.Close()
		err = page.FillInput(login_selector, login)
		if err != nil {
			panic(err.Error())
		}
		err = page.FillInput(password_selector, password)
		if err != nil {
			panic(err.Error())
		}
		err = page.ClickElement(submit_selector)
		if err != nil {
			panic(err.Error())
		}
		data, err := page.GetCookie()
		if err != nil {
			panic(err.Error())
		}
		thisCookies = cookies{data}

	})
	return &thisCookies
}
