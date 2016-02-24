package blockchain

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux
	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
	// client is the Blockchain.info client being tested.
	client *Client
)

// setup sets up a test HTTP server along with a Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provides mock responses for the API method being tested.
func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil, "w1731", "R@GK")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url.String()
	client.MerchantURL = url.String() + "/merchant"
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}
