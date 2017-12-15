package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	MACIIOT_DATABASE = "root:12345678@tcp(127.0.0.1:3306)/maxiiot_test?charset=utf8"
)

// NewEngine 返回xorm引擎
func NewEngine(driverName, dataSourceName string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("database connection error: %s", err)
	}
	for {
		if err := engine.Ping(); err != nil {
			log.Errorf("ping database error,will retry in 2s: %s", err)
			time.Sleep(time.Second * 2)
		} else {
			break
		}
	}
	engine.ShowSQL(true) //打印SQL
	engine.Logger().SetLevel(core.LOG_DEBUG)
	return engine, nil
}
