package svc

import (
	"cloud-disk/core/internal/config"
	"cloud-disk/core/internal/middleware"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/rest"
	bolt "go.etcd.io/bbolt"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	BoltDB *bolt.DB
	Auth	rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c),
		BoltDB: models.InitBolt(c),
		Auth:	 middleware.NewAuthMiddleware().Handle,
	}
}
