package cookie_http

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

type measuresResponse struct {
	Body    string      `json:"body"`
	Headers http.Header `json:"headers"`
	Status  int         `json:"status"`
	Timings timings     `json:"timings"`
}

type timings struct {
	DNS_Lookup      int64 `json:"dns_locup"`
	Connecting      int64 `json:"connecting"`
	TLS_handshaking int64 `json:"tls_handshaking"`
	Sending         int64 `json:"sending"`
	Waiting         int64 `json:"waiting"`
	Receiving       int64 `json:"receiving"`
	Duration        int64 `json:"duration"`
}

func createMeasures(measures httpstat.Result, requestTime time.Time) (timings timings) {
	timings.DNS_Lookup = int64(measures.DNSLookup / time.Millisecond)
	timings.Duration = int64(measures.Total(requestTime) / time.Millisecond)
	timings.TLS_handshaking = int64(measures.TLSHandshake / time.Millisecond)
	timings.Waiting = int64(measures.ServerProcessing / time.Millisecond)
	timings.Sending = int64((measures.StartTransfer + measures.Pretransfer) / time.Millisecond)
	timings.Connecting = int64(measures.Connect / time.Millisecond)
	return
}

func createResponse(response *http.Response, measures httpstat.Result, requestTime time.Time) (res measuresResponse) {
	res.Timings = createMeasures(measures, requestTime)
	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		res.Status = 0
		return
	}
	res.Body = string(respBody)
	res.Status = response.StatusCode
	res.Headers = response.Header

	return
}

func createErrorResponse(measures httpstat.Result, requestTime time.Time) (res measuresResponse) {
	res.Status = 0
	res.Timings = createMeasures(measures, requestTime)
	return
}
