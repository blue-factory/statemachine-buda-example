package bot

import "github.com/blue-factory/statemachine"

const (
	eventFinish = "event-finish"
)

func (b *Bot) FinishHandler(e *statemachine.Event) (*statemachine.Event, error) {
	return &statemachine.Event{Name: statemachine.EventAbort}, nil
}
