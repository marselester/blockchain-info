package blockchain

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestWalletAddresses(t *testing.T) {
	setup()
	defer teardown()

	js, err := ioutil.ReadFile("json/merchant_list.json")
	if err != nil {
		t.Error(err)
	}

	mux.HandleFunc("/merchant/w1731/list", func(w http.ResponseWriter, r *http.Request) {
		w.Write(js)
	})

	addrs, err := client.Wallet.Addresses()
	if err != nil {
		t.Error(err)
	}

	want := []WalletAddress{
		{Address: "15zyMv6T4SGkZ9ka3dj1BvSftvYuVVB66S", Balance: 20090584076},
	}
	if !reflect.DeepEqual(addrs, want) {
		t.Errorf("Wallet.Addresses returned %v, want %v", addrs, want)
	}
}
