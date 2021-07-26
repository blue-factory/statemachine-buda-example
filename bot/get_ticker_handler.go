package bot

import (
	"log"
	"strings"
	"time"

	"github.com/blue-factory/statemachine"
)

const (
	eventGetTicker = "event_get_ticker"
)

func (b *Bot) GetTickerHandler(e *statemachine.Event) (*statemachine.Event, error) {
	nextRequestStartAt := b.lastRequestMadeAt.Add(time.Second * time.Duration(b.config.TimeoutInSeconds))
	now := b.now()
	if now.Before(nextRequestStartAt) { // to avoid blocks by too many requests
		return &statemachine.Event{Name: eventGetTicker}, nil
	}

	ticker, err := b.buda.GetTicker(strings.ToUpper(b.config.Currency))
	if err != nil {
		log.Printf("bot: Bot GetTickerHandler b.buda.GetTicker error %s", err.Error())
		return &statemachine.Event{Name: eventGetTicker}, nil
	}

	b.lastRequestMadeAt = now

	return &statemachine.Event{
		Name: eventCalculate,
		Data: map[string]interface{}{
			"payload": &eventPayload{ticker: ticker},
		},
	}, nil
}
