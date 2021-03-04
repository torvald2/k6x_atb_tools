package oidc_login

import (
	"sync"
	"time"
)

var once sync.Once

type CookiesCreator struct{}

func (creator *CookiesCreator) New(server_url, target_url, login, password, login_selector, password_selector, submit_selector, main_page_selector string) *cookies {
	once.Do(func() {
		page, err := createPage(server_url, target_url)
		if err != nil {
			panic(err.Error())
		}
		defer page.Close()
		if err := page.WaitForWithTimeout(login_selector, 2*time.Second); err != nil {
			panic("2s wait for login page error")
		}
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
		if err := page.WaitForWithTimeout(main_page_selector, 2*time.Second); err != nil {
			panic("2s wait for main page error")
		}

		data, err := page.GetCookie()
		if err != nil {
			panic(err.Error())
		}
		thisCookies = cookies{data}

	})
	return &thisCookies
}
