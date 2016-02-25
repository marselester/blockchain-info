package blockchain

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestBlockchainAddressHasNoTxs(t *testing.T) {
	setup()
	defer teardown()

	js := `
{
	"hash160": "1a816ef92b552e5fb73dccfa8c98739b2da16035",
	"address": "13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY",
	"n_tx": 10,
	"total_received": 335550944460,
	"total_sent": 20090584076,
	"final_balance": 315460360384,
	"txs": []
}
`
	address := "13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY"
	mux.HandleFunc("/address/13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, js)
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

	if len(a.Txs) != 0 {
		t.Errorf("Blockchain.Address wrong transactions count %d, want 0", len(a.Txs))
	}
}

func TestBlockchainAddressWithTxs(t *testing.T) {
	setup()
	defer teardown()

	js := `
{
	"hash160": "1a816ef92b552e5fb73dccfa8c98739b2da16035",
	"address": "13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY",
	"n_tx": 10,
	"total_received": 335550944460,
	"total_sent": 20090584076,
	"final_balance": 315460360384,
	"txs": [{
		"ver":1,
		"inputs": [
			{
				"sequence": 4294967295,
				"prev_out": {
				"spent": true,
				"tx_index": 114831414,
				"type": 0,
				"addr": "1Bbq8wAAk3jFT7sdtArhsJrCisosHMxhKy",
				"value": 4600000000,
				"n": 1,
				"script": "76a9147447954676fac24a2c72a5b92407ead8157411e888ac"
			},
			"script":"483045022100ecaa92d4133e5aa77a0b7e9f9faf7f5562d5ff43d1b0b3d9af41578086a4d711022047eb75d22d696c28730ff66d16dd9307cc1a62ba6babd608e2c84468af10e4640121031f6e9b8aaaf76d05afff3fe8536eaa72387e0c0cf040e75d6bc85ce314c7dba5"
		}],
		"block_height": 387122,
		"relayed_by": "127.0.0.1",
		"out": [{
			"spent": true,
			"tx_index": 114834113,
			"type": 0,
			"addr": "3LKxFxbYeQaaRrKE1zRBxrHSzZftuTUDKB",
			"value": 4599990000,
			"n": 0,
			"script": "a914cc6e98586ab52d57bd6272b89295f943d1544a8687"
		}],
		"lock_time": 0,
		"result": 0,
		"size": 190,
		"time": 1449471605,
		"tx_index": 114834113,
		"vin_sz": 1,
		"hash": "d5e1ffb5e0a235731f84a0e616f4ad1264db43bd61e7a00751b1151b9b01b488",
		"vout_sz": 1
	}]
}
`
	address := "13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY"
	mux.HandleFunc("/address/13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, js)
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
