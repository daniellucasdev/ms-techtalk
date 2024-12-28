//nolint:ireturn,nolintlint
package bus

import (
	"github.com/asaskevich/EventBus"
)

//go:generate mockgen -destination=event_bus_mock.go -package=bus . EventBusI

type EventBusI interface {
	Publish(topic string, args ...interface{})
}

var bus EventBus.Bus //nolint:gochecknoglobals

func GetBus() EventBus.Bus {
	if bus == nil {
		bus = EventBus.New()
	}

	return bus
}
