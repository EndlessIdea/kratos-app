package data

import (
	"github.com/EndlessIdea/kratos-app/app/helloworld/internal/conf"
	"github.com/EndlessIdea/kratos-app/pkg/infra"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	mysqlClient    *gorm.DB
	postgresClient *gorm.DB
	redisClient    *redis.Client
}

// NewData .
func NewData(config *conf.Data, logger log.Logger) (result *Data, cleanup func(), err error) {
	result = &Data{}

	result.mysqlClient, err = infra.NewMysqlClient(config.Mysql)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "init mysql client error")
	}

	result.postgresClient, err = infra.NewPostgresClient(config.Postgres)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "init postgres client error")
	}

	result.redisClient, err = infra.NewRedisClient(config.Redis)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "init redis client error")
	}

	cleanup = func() {
		logger := log.NewHelper(logger)
		logger.Info("closing the data resources")

		mysqlDB, _ := result.mysqlClient.DB()
		if err := mysqlDB.Close(); err != nil {
			logger.Errorf("close mysql connection error %+v", err)
		}

		pgDB, _ := result.postgresClient.DB()
		if err := pgDB.Close(); err != nil {
			logger.Errorf("close pg connection error %+v", err)
		}

		if err := result.redisClient.Close(); err != nil {
			logger.Errorf("close redis connection error %+v", err)
		}
	}

	return result, cleanup, nil
}
