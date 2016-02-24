# Go client for [the Blockchain.info API](https://blockchain.info/api/blockchain_wallet_api)

The API client design is inspired by [go-github](https://github.com/google/go-github/)
and [stripe-go](https://github.com/stripe/stripe-go/).

## List addresses

```go
client = blockchain.NewClient(nil, "w1731", "R@GK")
fmt.Println(client.Address.List())
// [{15zyMv6T4SGkZ9ka3dj1BvSftvYuVVB66  20090584076}]
```

## Testing

```shell
$ make test
$ make lint
```
