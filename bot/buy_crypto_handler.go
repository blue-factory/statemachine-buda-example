package bot

import (
	"log"

	"github.com/blue-factory/statemachine"
	"github.com/mtavano/buda-go"
	"github.com/pkg/errors"
)

const (
	eventCreateOrder = "event_create_order"

	typeOrderBid   = "Bid"
	typePriceLimit = "market"
)

func (b *Bot) BuyCryptoHandler(e *statemachine.Event) (*statemachine.Event, error) {
	payload := b.parseData(e.Data)

	order, err := b.buda.CreateOrder(b.config.Currency, &buda.CreateOrderRequest{
		Type:      typeOrderBid,
		PriceType: typePriceLimit,
		Amount:    b.config.TargetVolume / payload.buyPrice,
	})
	if err != nil {
		return nil, errors.Wrap(err, "but: Bot.BuyCryptoHandler b.buda.CreateOrder error")
	}

	log.Printf("Bot.BuyCryptoHandler order created %+v", order)

	return &statemachine.Event{Name: eventFinish}, nil
}
