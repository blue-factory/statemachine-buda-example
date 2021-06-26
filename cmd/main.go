package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/blue-factory/statemachine-buda-example/bot"
	"github.com/mtavano/buda-go"
)

var bitSize = 64

func main() {
	// env vars
	budaAPIKey := os.Getenv("BUDA_API_KEY")
	budaAPISecret := os.Getenv("BUDA_API_SECRET")
	targetVolumeStr := os.Getenv("TARGET_VOLUME")
	rateToAcctionStr := os.Getenv("RATE_TO_ACTION")
	currency := os.Getenv("CORRENCY")

	targetVolume, err := strconv.ParseFloat(targetVolumeStr, bitSize)
	check(err)

	rateToAction, err := strconv.ParseFloat(rateToAcctionStr, bitSize)
	check(err)

	currency = strings.ToUpper(currency)

	// setup
	budaClient := buda.New(budaAPIKey, budaAPISecret, http.DefaultClient)

	b := bot.New(&bot.Config{
		TargetVolume:     targetVolume,
		RateToAction:     rateToAction,
		TimeoutInSeconds: 1, // 1 request per second
		Currency:         currency,
	},
		budaClient,
	)

	// run
	b.Start()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
