package bot

import (
	"fmt"
	"time"

	"github.com/blue-factory/statemachine"
	"github.com/mtavano/buda-go"
)

type Bot struct {
	sm             *statemachine.StateMachine
	referencePrice float64 // from local exchange
	buda           BudaClient

	config *Config

	// ticker use only
	now               TimeWrapper
	lastRequestMadeAt time.Time
}

// Config is a bot config type
type Config struct {
	TargetVolume     float64 // in local currency
	RateToAction     float64 // the fall that we are looking for
	TimeoutInSeconds int
	Currency         string
}

func New(config *Config, budaClient BudaClient) *Bot {
	b := &Bot{
		buda:   budaClient,
		config: config,
	}

	b.sm = statemachine.New(
		&statemachine.Event{Name: eventGetTicker},
		map[string]statemachine.State{
			eventGetTicker: {
				EventHandler: b.GetTickerHandler,
				Destination:  []string{eventGetTicker, eventCalculate},
			},
			eventCalculate: {
				EventHandler: b.CalculateHandler,
				Destination:  []string{eventCalculate, eventGetTicker, eventCreateOrder},
			},
			eventCreateOrder: {
				EventHandler: b.BuyCryptoHandler,
				Destination:  []string{eventFinish},
			},
			eventFinish: {
				EventHandler: b.FinishHandler,
				Destination:  []string{statemachine.EventAbort},
			},
		},
		nil,
	)

	return b
}

func (b *Bot) Start() {
	//b.sm.Run()
	fmt.Println()
	fmt.Println(b.sm.RenderMermaid())
}

// BudaClient Definition of buda client interface
type BudaClient interface {
	GetTicker(pair string) (*buda.Ticker, error)
	CreateOrder(pair string, req *buda.CreateOrderRequest) (*buda.Order, error)
	GetOrder(id string) (*buda.Order, error)
}

// TimeWrapper is the type alias for the .Now() function
type TimeWrapper func() time.Time

// payload passed through events
type eventPayload struct {
	ticker   *buda.Ticker
	buyPrice float64
}

// parseData parses the generic map into a eventPayload type
func (b *Bot) parseData(data map[string]interface{}) *eventPayload {
	return data["payload"].(*eventPayload)
}
