package blockchain

// Address represents a Bitcoin address.
type Address struct {
	Address string
	Balance int64
}

// AddressService handles communication with the Blockchain.info API resources.
type AddressService struct {
	client *Client
}

// List requests a slice of wallet addresses.
func (s *AddressService) List() ([]Address, error) {
	req, err := s.client.NewMerchantRequest("list")
	if err != nil {
		return nil, err
	}

	var aux struct {
		Addresses []Address
	}
	_, err = s.client.Do(req, &aux)
	if err != nil {
		return nil, err
	}

	return aux.Addresses, nil
}
