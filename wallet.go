package blockchain

// WalletAddress represents a Bitcoin wallet address.
type WalletAddress struct {
	Address string
	Balance int64
}

// walletService handles communication with the Blockchain.info Wallet API.
type walletService struct {
	client *Client
}

// Addresses requests a slice of wallet addresses.
func (s *walletService) Addresses() ([]WalletAddress, error) {
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
