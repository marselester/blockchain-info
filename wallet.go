package blockchain

// WalletAddress represents a Bitcoin wallet address.
type WalletAddress struct {
	Address string
	Balance int64
}

// WalletService handles communication with the Blockchain.info merchant API.
type WalletService struct {
	client *Client
}

// Addresses requests a slice of wallet addresses.
func (s *WalletService) Addresses() ([]WalletAddress, error) {
	req, err := s.client.NewMerchantRequest("list")
	if err != nil {
		return nil, err
	}

	var aux struct {
		Addresses []WalletAddress
	}
	_, err = s.client.Do(req, &aux)
	if err != nil {
		return nil, err
	}

	return aux.Addresses, nil
}
