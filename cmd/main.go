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
var shouldRunOption = "1"

func main() {
	// env vars
	budaAPIKey := os.Getenv("BUDA_API_KEY")
	budaAPISecret := os.Getenv("BUDA_API_SECRET")
	targetVolumeStr := os.Getenv("TARGET_VOLUME")
	rateToAcctionStr := os.Getenv("RATE_TO_ACTION")
	currency := os.Getenv("CURRENCY")
	shouldRun := os.Getenv("SHOULD_RUN")

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
		TimeoutInSeconds: 2, // 1 request every 2 seconds
		Currency:         currency,
	},
		budaClient,
	)

	// run
	b.Render()
	if shouldRun == shouldRunOption {
		b.Start()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
