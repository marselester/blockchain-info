# Go client for [the Blockchain.info API](https://blockchain.info/api/blockchain_wallet_api)

The API client design is inspired by [go-github](https://github.com/google/go-github/)
and [stripe-go](https://github.com/stripe/stripe-go/).

## List Wallet Addresses

```go
client = blockchain.NewClient(nil, "w1731", "R@GK")
fmt.Println(client.Wallet.Addresses())
// [{13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY  20090584076}]
```

## Get Address Summary

Address represents https://blockchain.info/address/13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY?format=json
resource.

```go
client = blockchain.NewClient(nil, "", "")
fmt.Println(client.Data.Address("13R9dBgKwBP29JKo11zhfi74YuBsMxJ4qY"))
```

## Testing

```shell
$ make test
$ make lint
```
