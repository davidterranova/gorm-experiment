package pg

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Schema   string `envconfig:"SCHEMA" default:"public"`
	Username string `envconfig:"USERNAME" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Host     string `envconfig:"HOST" required:"true"`
	Port     int    `envconfig:"PORT" default:"5432"`
	Name     string `envconfig:"NAME" required:"true"`
	Params   string `envconfig:"PARAMS" default:"sslmode=disable"`
}

func Open(cfg DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d search_path=%s %s",
		cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port, cfg.Schema, cfg.Params,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
