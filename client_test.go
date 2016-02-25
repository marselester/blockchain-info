package blockchain

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	walletID := "w1731"
	mainPass := "R@GK"
	c := NewClient(nil, walletID, mainPass)

	if c.HTTPClient != http.DefaultClient {
		t.Errorf("NewClient should have default HTTP client, got %v", c.HTTPClient)
	}

	baseURL := "https://blockchain.info"
	if c.BaseURL != baseURL {
		t.Errorf("NewClient should have base URL %q, got %v", baseURL, c.BaseURL)
	}

	merchantURL := "https://blockchain.info/merchant"
	if c.MerchantURL != merchantURL {
		t.Errorf("NewClient should have merchant URL %q, got %v", merchantURL, c.MerchantURL)
	}

	if c.WalletID != walletID {
		t.Errorf("NewClient wrong wallet ID %q, want %q", c.WalletID, walletID)
	}

	if c.MainPassword != mainPass {
		t.Errorf("NewClient wrong main password %q, want %q", c.MainPassword, mainPass)
	}
}

func TestClientNewRequest(t *testing.T) {
	c := NewClient(nil, "", "")
	req, err := c.NewRequest("address/15zyMv6T4SGkZ9ka3dj1BvSftvYuVVB66S")
	if err != nil {
		t.Error(err)
	}

	if req.Method != "GET" {
		t.Errorf("NewRequest wrong method %q, want GET", req.Method)
	}

	uri := "https://blockchain.info/address/15zyMv6T4SGkZ9ka3dj1BvSftvYuVVB66S?format=json"
	if req.URL.String() != uri {
		t.Errorf("NewRequest wrong URL %q, want %q", req.URL, uri)
	}
}

func TestClientNewMerchantRequest(t *testing.T) {
	c := NewClient(nil, "w1731", "R@GK")
	req, err := c.NewMerchantRequest("list")
	if err != nil {
		t.Error(err)
	}

	if req.Method != "GET" {
		t.Errorf("NewMerchantRequest wrong method %q, want GET", req.Method)
	}

	uri := "https://blockchain.info/merchant/w1731/list?password=R%40GK"
	if req.URL.String() != uri {
		t.Errorf("NewMerchantRequest wrong URL %q, want %q", req.URL, uri)
	}
}
