package db

import (
	"errors"
	"fmt"
	"github.com/apiboxgo/library-utils/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DbDriverPostgres = "postgres"

func ConnectionFactory(config *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DbHost,
		config.DbUser,
		config.DbPassword,
		config.DbName,
		config.DbPort,
		config.DbSSLMode,
	)

	switch config.DbDriver {
	case DbDriverPostgres:
		return gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	return nil, errors.New(fmt.Sprintf("Driver %s not found", config.DbDriver))
}
