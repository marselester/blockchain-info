package blockchain

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestBlockchainAddress(t *testing.T) {
	setup()
	defer teardown()

	js, err := ioutil.ReadFile("json/blockchain_address.json")
	if err != nil {
		t.Error(err)
	}

	address := "13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY"
	mux.HandleFunc("/address/13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY", func(w http.ResponseWriter, r *http.Request) {
		w.Write(js)
	})

	a, err := client.Blockchain.Address(address)
	if err != nil {
		t.Error(err)
	}

	if a.Address != address {
		t.Errorf("Blockchain.Address wrong address %q, want %q", a.Address, address)
	}

	if a.TxCount != 10 {
		t.Errorf("Blockchain.Address wrong tx count %d, want %d", a.TxCount, 10)
	}

	if a.TotalReceived != 335550944460 {
		t.Errorf("Blockchain.Address wrong total received %d, want %d", a.TotalReceived, 335550944460)
	}

	if a.TotalSent != 20090584076 {
		t.Errorf("Blockchain.Address wrong total sent %d, want %d", a.TotalSent, 20090584076)
	}

	if a.FinalBalance != 315460360384 {
		t.Errorf("Blockchain.Address wrong final balance %d, want %d", a.FinalBalance, 315460360384)
	}
}

func TestBlockchainAddressTxs(t *testing.T) {
	setup()
	defer teardown()

	js, err := ioutil.ReadFile("json/blockchain_address.json")
	if err != nil {
		t.Error(err)
	}

	address := "13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY"
	mux.HandleFunc("/address/13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY", func(w http.ResponseWriter, r *http.Request) {
		w.Write(js)
	})

	a, err := client.Blockchain.Address(address)
	if err != nil {
		t.Error(err)
	}

	if len(a.Txs) != 1 {
		t.Errorf("Blockchain.Address wrong transactions count %d, want 1", len(a.Txs))
	}

	tx := a.Txs[0]

	if tx.Index != 114834113 {
		t.Errorf("Blockchain.Address wrong tx index %d, want 114834113", tx.Index)
	}

	hash := "d5e1ffb5e0a235731f84a0e616f4ad1264db43bd61e7a00751b1151b9b01b488"
	if tx.Hash != hash {
		t.Errorf("Blockchain.Address wrong tx hash %q, want %q", tx.Hash, hash)
	}

	if tx.BlockHeight != 387122 {
		t.Errorf("Blockchain.Address wrong tx block height %d, want 387122", tx.BlockHeight)
	}

	timestamp := time.Unix(1449471605, 0)
	if tx.Timestamp != Timestamp(timestamp) {
		t.Errorf("Blockchain.Address wrong tx timestamp %v, want %v", tx.Timestamp, timestamp)
	}
}
