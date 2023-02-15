package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/levor/to-do/internal/config"
	"log"
)

func NewConnection(cfg *config.Config) (*gorm.DB, error) {
	connection, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True", // username:password@protocol(host)/dbname?param=value
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbScheme))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if cfg.IsDebugMode {
		connection.LogMode(true)
	}
	return connection, nil
}
