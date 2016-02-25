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
		t.Errorf("Blockchain.Address address %q, want %q", a.Address, address)
	}

	if a.TxCount != 10 {
		t.Errorf("Blockchain.Address tx count %d, want %d", a.TxCount, 10)
	}

	if a.TotalReceived != 335550944460 {
		t.Errorf("Blockchain.Address total received %d, want %d", a.TotalReceived, 335550944460)
	}

	if a.TotalSent != 20090584076 {
		t.Errorf("Blockchain.Address total sent %d, want %d", a.TotalSent, 20090584076)
	}

	if a.FinalBalance != 315460360384 {
		t.Errorf("Blockchain.Address final balance %d, want %d", a.FinalBalance, 315460360384)
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
		t.Errorf("Blockchain.Address transactions count %d, want 1", len(a.Txs))
	}

	tx := a.Txs[0]

	if tx.Index != 114834113 {
		t.Errorf("Blockchain.Address tx index %d, want 114834113", tx.Index)
	}

	hash := "d5e1ffb5e0a235731f84a0e616f4ad1264db43bd61e7a00751b1151b9b01b488"
	if tx.Hash != hash {
		t.Errorf("Blockchain.Address tx hash %q, want %q", tx.Hash, hash)
	}

	if tx.BlockHeight != 387122 {
		t.Errorf("Blockchain.Address tx block height %d, want 387122", tx.BlockHeight)
	}

	timestamp := time.Unix(1449471605, 0)
	if tx.Timestamp != Timestamp(timestamp) {
		t.Errorf("Blockchain.Address tx timestamp %v, want %v", tx.Timestamp, timestamp)
	}
}

func TestBlockchainAddressTxsOutput(t *testing.T) {
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

	out := a.Txs[0].Outputs[0]

	if out.TxIndex != 114834113 {
		t.Errorf("Blockchain.Address output tx index %d, want 114834113", out.TxIndex)
	}

	if out.N != 1 {
		t.Errorf("Blockchain.Address output number %d, want 1", out.N)
	}

	outAddr := "3LKxFxbYeQaaRrKE1zRBxrHSzZftuTUDKB"
	if out.Address != outAddr {
		t.Errorf("Blockchain.Address output address %q, want %q", out.Address, outAddr)
	}

	if out.Value != 4599990000 {
		t.Errorf("Blockchain.Address output value %d, want 4599990000", out.Value)
	}

	if !out.IsSpent {
		t.Errorf("Blockchain.Address output spent %t, want true", out.IsSpent)
	}

	script := "a914cc6e98586ab52d57bd6272b89295f943d1544a8687"
	if out.Script != script {
		t.Errorf("Blockchain.Address output script %q, want %q", out.Script, script)
	}
}
