package varnish

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const defaultHTTPTimeout = 10

// Invalidator is the API to varnish
type Invalidator struct {
	hostname  string
	port      int
	protocol  string
	debug     bool
	keepAlive bool

	httpClient        *http.Client
	beforeRequestFunc func(*http.Request)
}

func NewInvalidator(varnishAddress string, port int, keepAlive bool) (*Invalidator, error) {

	addr, err := url.Parse(varnishAddress)
	if err != nil {
		return &Invalidator{}, fmt.Errorf("error parsing varnish address")
	}

	if !addr.IsAbs() {
		return &Invalidator{}, fmt.Errorf("no valid protocol supplied (varnish address should start with http or https)")
	}

	invalidator := Invalidator{
		httpClient: &http.Client{Timeout: time.Second * defaultHTTPTimeout},
		hostname:   addr.Hostname(),
		port:       port,
		protocol:   addr.Scheme,
		keepAlive:  keepAlive,
	}

	return &invalidator, nil
}

// BeforeRequest can be used to inject behavior before sending a request to the client
// for example: setting specific headers
func (i *Invalidator) BeforeRequest(f func(*http.Request)) {
	i.beforeRequestFunc = f
}
