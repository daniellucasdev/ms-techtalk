package handlers

import (
	"github.com/asaskevich/EventBus"
	"github.com/braiphub/go-core/log"
	"github.com/braiphub/go-scaffold/internal/domain/service"
	"github.com/braiphub/go-scaffold/internal/events/bus"
	"github.com/braiphub/go-scaffold/internal/events/event"
	"github.com/braiphub/go-scaffold/internal/infra/anticorruption/msbooks"
)

type EventHandler struct {
	logger         log.LoggerI
	bus            EventBus.Bus
	msBooksAdapter *msbooks.MsBooksAdapter
	chapterService *service.ChapterService
}

func NewEventHandler(
	logger log.LoggerI,
	msBooksAdapter *msbooks.MsBooksAdapter,
	chapterService *service.ChapterService,
) *EventHandler {
	return &EventHandler{
		logger:         logger,
		bus:            bus.GetBus(),
		msBooksAdapter: msBooksAdapter,
		chapterService: chapterService,
	}
}

//nolint:errcheck // it'll be changed soon with another implementation with error logging
func (handler *EventHandler) StartListeners() {
	handler.bus.SubscribeAsync(event.BookCreated, handler.LogBookCreated, false)
	handler.bus.SubscribeAsync(event.BookCreated, handler.InitFirstChapter, false)

	handler.bus.SubscribeAsync(event.ChapterCreated, handler.LogChapterCreated, false)
	handler.bus.SubscribeAsync(event.ChapterCreated, handler.NotifyChapterCreatedToMsBooks, false)
}
