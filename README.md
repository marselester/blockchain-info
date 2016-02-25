# Go client for [the Blockchain.info API](https://blockchain.info/api/blockchain_wallet_api)

The API client design is inspired by [go-github](https://github.com/google/go-github/)
and [stripe-go](https://github.com/stripe/stripe-go/).

## Wallet API

### List Addresses

```go
c := blockchain.NewClient(nil, "w1731", "R@GK")
fmt.Println(c.Wallet.Addresses())
// [{13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY  20090584076}]
```

## Block Explorer API

### Get Address

Address represents https://blockchain.info/address/13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY?format=json
resource.

```go
c := blockchain.NewClient(nil, "", "")
fmt.Println(c.Blockchain.Address("13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY"))
```

## Testing

```shell
$ make test
$ make lint
```
