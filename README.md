# Go client for [the Blockchain.info API](https://blockchain.info/api/blockchain_wallet_api)

The API client design is inspired by [go-github](https://github.com/google/go-github/)
and [stripe-go](https://github.com/stripe/stripe-go/).

## Wallet API

### List Addresses

```go
walletID := "w1731"
walletPass := "R@GK"
apiCode := "123" // Set apiCode to bypass the request limiter.
c := blockchain.NewClient(nil, walletID, walletPass, apiCode)
fmt.Println(c.Wallet.Addresses())
// [{13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY  20090584076}]
```

## Block Explorer API

### Get Address

Address represents https://blockchain.info/address/13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY?format=json
resource.

```go
// Set apiCode to bypass the request limiter.
apiCode := "123"
c := blockchain.NewClient(nil, "", "", apiCode)
fmt.Println(c.Blockchain.Address("13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY"))
```

## Testing

```shell
$ make test
$ make lint
```
