package xk6_atb_tools

import (
	"github.com/loadimpact/k6/js/modules"
	"github.com/torvald2/k6x_atb_tools/cookie_http"
	"github.com/torvald2/k6x_atb_tools/increments"
	"github.com/torvald2/k6x_atb_tools/oidc_login"
)

func init() {
	modules.Register("k6/x/atb/httpclient", new(cookie_http.HttpClientCreator))
	modules.Register("k6/x/atb/dateincrement", new(increments.DateIncrementCreator))
	modules.Register("k6/x/atb/numericincrement", new(increments.NumIncrementCreator))
	modules.Register("k6/x/atb/cookies", new(oidc_login.CookiesCreator))

}
