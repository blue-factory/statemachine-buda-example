package bot

import (
	"time"

	"github.com/blue-factory/statemachine"
	"github.com/mtavano/buda-go"
)

type Bot struct {
	sm             *statemachine.StateMachine
	referencePrice float64 // from local exchange
	buda           BudaClient

	config Config

	// ticker use only
	now               TimeWrapper
	lastRequestMadeAt time.Time
}

type Config struct {
	TargetVolume     float64 // in local currency
	RateToAction     float64 // the fall that we are looking for
	TimeoutInSeconds int
	Currency         string
}

type BudaClient interface {
	GetTicker(pair string) (*buda.Ticker, error)
	CreateOrder(pair string, req *buda.CreateOrderRequest) (*buda.Order, error)
	GetOrder(id string) (*buda.Order, error)
}

type TimeWrapper func() time.Time

type eventPayload struct {
	ticker   *buda.Ticker
	buyPrice float64
}

func (b *Bot) parseData(data map[string]interface{}) *eventPayload {
	return data["payload"].(*eventPayload)
}

var bitSize = 64
