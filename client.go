package blockchain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	baseURL     = "https://blockchain.info"
	merchantURL = "https://blockchain.info/merchant"
)

// Client manages communication with the Blockchain.info API.
type Client struct {
	HTTPClient *http.Client

	BaseURL      string
	MerchantURL  string
	WalletID     string
	MainPassword string
	APICode      string // API Code to bypass the request limiter.

	Wallet     *walletService
	Blockchain *blockchainService
}

// NewClient returns a new Blockchain.info API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client, wallet, pass, code string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		HTTPClient:   httpClient,
		BaseURL:      baseURL,
		MerchantURL:  merchantURL,
		WalletID:     wallet,
		MainPassword: pass,
		APICode:      code,
	}
	c.Wallet = &walletService{client: c}
	c.Blockchain = &blockchainService{client: c}
	return c
}

// NewRequest creates a request to the public API.
func (c *Client) NewRequest(path string) (*http.Request, error) {
	urlStr := fmt.Sprintf("%s/%s", c.BaseURL, path)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("format", "json")
	if c.APICode != "" {
		q.Set("api_code", c.APICode)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	return req, err
}

// NewMerchantRequest creates a request to the Merchant API.
func (c *Client) NewMerchantRequest(path string) (*http.Request, error) {
	urlStr := fmt.Sprintf("%s/%s/%s", c.MerchantURL, c.WalletID, path)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("password", c.MainPassword)
	if c.APICode != "" {
		q.Set("api_code", c.APICode)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	return req, err
}

// Do uses the Blockchain.info API client's HTTP client to execute the request
// and unmarshals the response into v.
// It also handles unmarshaling errors returned by the API.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("blockchain request failed: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

var (
	// ReqLimit defines how many requests are allowed to be sent in ReqWindow.
	// Default value is 6 requests in 5 minutes.
	ReqLimit = 6
	// ReqWindow is how long we should wait before trying to send another requests.
	ReqWindow = time.Duration(5 * time.Minute)
)
var throttle = struct {
	sync.Mutex
	// reqCount is how many requests were sent in a ReqWindow.
	reqCount int
	// firstReqSentAt is when we started a ReqWindow.
	firstReqSentAt time.Time
}{}

// IsReqThrottled helps to keep track of API requests to avoid hitting API limits.
// When limit is reached, it returns estimated time when it's ok to try again.
func IsReqThrottled() (bool, time.Duration) {
	throttle.Lock()
	defer throttle.Unlock()

	if throttle.reqCount >= ReqLimit {
		// Let's check if we can send more requests.
		elapsed := time.Since(throttle.firstReqSentAt)
		if elapsed >= ReqWindow {
			throttle.reqCount = 0
		} else {
			// Estimated time for next request attempt.
			return true, ReqWindow - elapsed
		}
	}

	throttle.reqCount++
	if throttle.reqCount == 1 {
		throttle.firstReqSentAt = time.Now()
	}
	return false, time.Duration(0)
}
