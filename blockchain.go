package blockchain

import (
	"encoding/json"
	"fmt"
	"time"
)

// Output is a transaction output.
type Output struct {
	// Blockchain.info's internal transcation ID.
	TxIndex uint64
	// Output index.
	N       uint32
	Address string
	// Amount in satoshis.
	Value   uint64
	IsSpent bool
	Script  string
}

// Input is a transaction input.
type Input struct {
	PrevOutput *Output
	Script     string
}

// Tx represents a Bitcoin transaction.
type Tx struct {
	// Blockchain.info's internal transcation ID.
	Index       uint64 `json:"tx_index"`
	Hash        string
	BlockHeight uint32    `json:"block_height"`
	Timestamp   Timestamp `json:"time"`
	Inputs      []Input
	Outputs     []Output
}

// Address provides a summary of Bitcoin address.
type Address struct {
	Address       string
	TxCount       uint   `json:"n_tx"`
	TotalReceived uint64 `json:"total_received"`
	TotalSent     uint64 `json:"total_sent"`
	FinalBalance  int64  `json:"final_balance"`
	Txs           []Tx
}

// Timestamp is used to parse Unix time in 1449471605 format.
type Timestamp time.Time

// UnmarshalJSON decodes Unix time given in seconds to Timestamp (which is time.Time)
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var sec int64
	if err := json.Unmarshal(data, &sec); err != nil {
		return fmt.Errorf("time should be an int, got %d", data)
	}

	*t = Timestamp(time.Unix(sec, 0))
	return nil
}

// BlockchainService handles communication with the Blockchain.info Block Explorer API.
type blockchainService struct {
	client *Client
}

// Address requests an address summary.
func (s *blockchainService) Address(address string) (*Address, error) {
	req, err := s.client.NewRequest("address/" + address)
	if err != nil {
		return nil, err
	}

	v := new(Address)
	_, err = s.client.Do(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
