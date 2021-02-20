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

func createResponse(response *http.Response, measures httpstat.Result) (res measuresResponse, err error) {
	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	readBodyTime := time.Now()
	res.Body = string(respBody)
	res.Status = response.StatusCode
	res.Headers = response.Header
	res.Timings.DNS_Lookup = int64(measures.DNSLookup / time.Millisecond)
	res.Timings.Duration = int64(measures.Total(readBodyTime) / time.Millisecond)
	res.Timings.TLS_handshaking = int64(measures.TLSHandshake / time.Millisecond)
	res.Timings.Waiting = int64(measures.ServerProcessing / time.Millisecond)
	res.Timings.Sending = int64((measures.StartTransfer + measures.Pretransfer) / time.Millisecond)
	res.Timings.Connecting = int64(measures.Connect / time.Millisecond)

	return
}
