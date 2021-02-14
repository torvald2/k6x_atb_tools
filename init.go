package xk6_atb_tools

import (
	"github.com/loadimpact/k6/js/modules"
	"github.com/torvald2/k6x_atb_tools/increments"
)

func init() {
	modules.Register("k6/atb/dateincrement", new(increments.DateIncrementCreator))
	modules.Register("k6/atb/numincrement", new(increments.NumIncrementCreator))
}
