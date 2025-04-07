package dig

import (
	"github.com/braiphub/go-core/cache"
	"github.com/braiphub/go-core/hashid"
	"github.com/braiphub/go-core/log"
	"github.com/braiphub/go-core/queue"
	"github.com/braiphub/ms-tech-talk/internal/api/http/controller"
	"github.com/braiphub/ms-tech-talk/internal/domain/repository"
	"github.com/braiphub/ms-tech-talk/internal/domain/service"
	handlers "github.com/braiphub/ms-tech-talk/internal/events/handler"
	"github.com/braiphub/ms-tech-talk/internal/infra/anticorruption/msproducts"
	"gorm.io/gorm"
)

type IoCContainer struct {
	logger   *log.ZapLoggerAdapter
	readDB   *gorm.DB
	writeDB  *gorm.DB
	rabbitMQ *queue.RabbitMQConnection
	cache    cache.Cacherer
	hasher   hashid.Hasher
}

func NewIoCContainer(
	logger *log.ZapLoggerAdapter,
	readDB *gorm.DB,
	writeDB *gorm.DB,
	rabbitMQ *queue.RabbitMQConnection,
	cache cache.Cacherer,
	hasher hashid.Hasher,
) *IoCContainer {
	return &IoCContainer{
		logger:   logger,
		readDB:   readDB,
		writeDB:  writeDB,
		rabbitMQ: rabbitMQ,
		cache:    cache,
		hasher:   hasher,
	}
}

func (c *IoCContainer) SubscriptionController() *controller.SubscriptionController {
	return controller.NewSubscriptionController(
		c.SubscriptionService(),
	)
}

func (c *IoCContainer) OfferService() *service.OfferService {
	return service.NewOfferService(
		c.WriteOfferRepository(),
		c.ReadOfferRepository(),
	)
}

func (c *IoCContainer) SubscriptionService() *service.SubscriptionService {
	return service.NewSubscriptionService(
		c.WriteSubscriptionRepository(),
	)
}

func (c *IoCContainer) WriteOfferRepository() *repository.WriteOfferRepository {
	return repository.NewWriteOfferRepository(c.writeDB)
}

func (c *IoCContainer) ReadOfferRepository() *repository.ReadOfferRepository {
	return repository.NewReadOfferRepository(c.readDB)
}

func (c *IoCContainer) WriteSubscriptionRepository() *repository.WriteSubscriptionRepository {
	return repository.NewWriteSubscriptionRepository(c.writeDB)
}

func (c *IoCContainer) MsProductsAdapter() *msproducts.Adapter {
	return msproducts.NewAdapter(c.rabbitMQ, c.OfferService())
}

func (c *IoCContainer) EventHandler() *handlers.EventHandler {
	return handlers.NewEventHandler(
		c.logger,
	)
}

func (c *IoCContainer) Logger() *log.ZapLoggerAdapter {
	return c.logger
}

func (c *IoCContainer) RabbitClient() *queue.RabbitMQConnection {
	return c.rabbitMQ
}
