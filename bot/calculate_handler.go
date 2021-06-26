package bot

import (
	"strconv"

	"github.com/blue-factory/statemachine"
	"github.com/pkg/errors"
)

const (
	eventCalculate = "event-calculate"
)

var bitSize = 64

func (b *Bot) CalculateHandler(e *statemachine.Event) (*statemachine.Event, error) {
	payload := b.parseData(e.Data)

	ask, err := strconv.ParseFloat(payload.ticker.MinAsk[0], bitSize)
	if err != nil {
		return nil, errors.Wrap(err, "bot: Bot CalculateHandler strconv.ParseFloat error")
	}

	// if is the first execution
	if b.referencePrice == 0 {
		b.referencePrice = ask
	}

	if ask <= (b.referencePrice*1 + (b.config.RateToAction / 100)) {
		return &statemachine.Event{Name: eventGetTicker}, nil
	}

	return &statemachine.Event{
		Name: eventCreateOrder,
		Data: map[string]interface{}{
			"payload": &eventPayload{
				buyPrice: ask,
			},
		},
	}, nil
}
