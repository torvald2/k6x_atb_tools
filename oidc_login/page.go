package oidc_login

import (
	"time"

	conditions "github.com/serge1peshcoff/selenium-go-conditions"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Page struct {
	driver selenium.WebDriver
}

func (page *Page) FillInput(selector, text string) error {
	if elem, err := page.driver.FindElement(selenium.ByCSSSelector, selector); err == nil {
		err = elem.SendKeys(text)
		return err
	} else {
		return err
	}
}

func (page *Page) ClickElement(selector string) error {
	if elem, err := page.driver.FindElement(selenium.ByCSSSelector, selector); err == nil {
		err = elem.Click()
		return err
	} else {
		return err
	}
}
func (page *Page) GetCookie() (cookies map[string]string, err error) {
	tmp_cookie, err := page.driver.GetCookies()
	if err != nil {
		return
	}
	for _, cookie := range tmp_cookie {
		cookies[cookie.Name] = cookie.Value
	}
	return
}

func (page *Page) WaitForWithTimeout(selector string, timeForWait time.Duration) error {
	return page.driver.WaitWithTimeout(conditions.ElementIsLocated(selenium.ByCSSSelector, selector), timeForWait)
}

func (page *Page) Close() error {
	return page.driver.Close()
}

func createPage(server_url, target_url string) (page Page, err error) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--headless",
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
		},
	}
	caps.AddChrome(chromeCaps)
	driver, err := selenium.NewRemote(caps, server_url)
	if err != nil {
		return
	}
	err = driver.Get(target_url)
	return Page{driver}, err

}
