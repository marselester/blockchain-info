package blockchain

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAddressList(t *testing.T) {
	setup()
	defer teardown()

	js := `
{"addresses": [
	{
		"address": "15zyMv6T4SGkZ9ka3dj1BvSftvYuVVB66S",
		"balance": 20090584076,
		"label": null,
		"total_received": 335550944460
	}
]}
`
	mux.HandleFunc("/merchant/w1731/list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, js)
	})

	addrs, err := client.Address.List()
	if err != nil {
		t.Error(err)
	}

	want := []Address{
		{Address: "15zyMv6T4SGkZ9ka3dj1BvSftvYuVVB66S", Balance: 20090584076},
	}
	if !reflect.DeepEqual(addrs, want) {
		t.Errorf("Address.List returned %v, want %v", addrs, want)
	}
}
