# Statemaachine buda example

This repo is a simple example of a basic buy-strategy applied over Buda.com exchange.

You only need to setup the following environment variables before running the main file

```
BUDA_API_KEY= // Buda.con api key
BUDA_API_SECRET= // Buda.com api secret
TARGET_VOLUME= // how much CLP we want to expent
RATE_TO_ACTION= // expected fall price to buy
CURRENCY= // currency to trade {BTC; ETH; LTC; BCH}
```

Then you can simply run

```bash
go run ./cmd/main.go
```

The strategy diagram is the following


