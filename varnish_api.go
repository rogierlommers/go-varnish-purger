package varnish

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

const defaultHTTPTimeout = 10

// Invalidator is the API to varnish
type Invalidator struct {
	hostname  string
	port      int64
	protocol  string
	debug     bool
	keepAlive bool

	httpClient        *http.Client
	beforeRequestFunc func(*http.Request)
}

func NewInvalidator(varnishAddress string, port int64, keepAlive bool) (*Invalidator, error) {

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

// SetRetryConfig can be used to inject a retry configuration to the http client
// The retry policy is exponential backoff
func (i *Invalidator) SetRetryConfig(retryWaitMin time.Duration, retryMax int) {
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Timeout = defaultHTTPTimeout * time.Second

	retryClient.RetryWaitMin = retryWaitMin
	retryClient.RetryMax = retryMax

	i.httpClient = retryClient.StandardClient()
}

// BeforeRequest can be used to inject behavior before sending a request to the client
// for example: setting specific headers
func (i *Invalidator) BeforeRequest(f func(*http.Request)) {
	i.beforeRequestFunc = f
}

// GetHostname can be used to retrieve the varnish cache hostname
func (i *Invalidator) GetHostname() string {
	return i.hostname
}
