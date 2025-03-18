package dig

import (
	"github.com/braiphub/go-core/cache"
	"github.com/braiphub/go-core/hashid"
	"github.com/braiphub/go-core/log"
	"github.com/braiphub/go-core/queue"
	"github.com/braiphub/go-scaffold/internal/api/http/controller"
	"github.com/braiphub/go-scaffold/internal/domain/repository"
	"github.com/braiphub/go-scaffold/internal/domain/service"
	handlers "github.com/braiphub/go-scaffold/internal/events/handler"
	"github.com/braiphub/go-scaffold/internal/infra/anticorruption/msbooks"
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

func (c *IoCContainer) BookService() *service.BookService {
	return service.NewBookService(
		c.WriteBookRepository(),
		c.ReadBookRepository(),
		c.logger,
	)
}

func (c *IoCContainer) ChapterService() *service.ChapterService {
	return service.NewChapterService(
		c.WriteChapterRepository(),
		c.ReadChapterRepository(),
		c.logger,
	)
}

func (c *IoCContainer) BookController() *controller.BookController {
	return controller.NewBookController(c.BookService())
}

func (c *IoCContainer) ChapterController() *controller.ChapterController {
	return controller.NewChapterController(c.ChapterService())
}

func (c *IoCContainer) WriteBookRepository() *repository.WriteBookRepository {
	return repository.NewWriteBookRepository(c.writeDB)
}

func (c *IoCContainer) ReadBookRepository() *repository.ReadBookReadRepository {
	return repository.NewBookReadRepository(c.readDB)
}

func (c *IoCContainer) WriteChapterRepository() *repository.WriteChapterRepository {
	return repository.NewWriteChapterRepository(c.writeDB, c.hasher)
}

func (c *IoCContainer) ReadChapterRepository() *repository.ReadChapterReadRepository {
	return repository.NewChapterReadRepository(c.readDB)
}

func (c *IoCContainer) MsBooksAdapter() *msbooks.MsBooksAdapter {
	return msbooks.NewMsBooksAdapter(
		c.rabbitMQ,
		c.BookService(),
	)
}

func (c *IoCContainer) EventHandler() *handlers.EventHandler {
	return handlers.NewEventHandler(
		c.logger,
		c.MsBooksAdapter(),
		c.ChapterService(),
	)
}

func (c *IoCContainer) Logger() *log.ZapLoggerAdapter {
	return c.logger
}

func (c *IoCContainer) RabbitClient() *queue.RabbitMQConnection {
	return c.rabbitMQ
}
