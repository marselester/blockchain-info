package blockchain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

	Address *AddressService
}

// NewClient returns a new Blockchain.info API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client, wallet, pass string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		HTTPClient:   httpClient,
		BaseURL:      baseURL,
		MerchantURL:  merchantURL,
		WalletID:     wallet,
		MainPassword: pass,
	}
	c.Address = &AddressService{client: c}
	return c
}

// NewRequest creates a request to the public API.
func (c *Client) NewRequest(path string) (*http.Request, error) {
	urlStr := fmt.Sprintf("%s/%s?format=json", c.BaseURL, path)

	req, err := http.NewRequest("GET", urlStr, nil)
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
