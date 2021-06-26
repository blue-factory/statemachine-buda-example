package bot

import (
	"strings"
	"time"

	"github.com/blue-factory/statemachine"
	"github.com/pkg/errors"
)

const (
	eventGetTicker = "event_get_ticker"
)

func (b *Bot) GetTickerHandler(e *statemachine.Event) (*statemachine.Event, error) {
	nextRequestStartAt := b.lastRequestMadeAt.Add(time.Second * time.Duration(b.config.TimeoutInSeconds))
	now := b.now()
	if now.Before(nextRequestStartAt) {
		return e, nil
	}

	ticker, err := b.buda.GetTicker(strings.ToUpper(b.config.Currency))
	if err != nil {
		return nil, errors.Wrap(err, "bot: Bot GetTickerHandler b.buda.GetTicker error")
	}

	b.lastRequestMadeAt = now

	return &statemachine.Event{
		Name: eventCalculate,
		Data: map[string]interface{}{
			"payload": &eventPayload{ticker: ticker},
		},
	}, nil
}
