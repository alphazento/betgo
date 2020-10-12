package orm

import (
	"fmt"

	"github.com/alphazento/betgo/pkg/config"
	"github.com/alphazento/betgo/pkg/container"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
	"github.com/mitchellh/mapstructure"
)

func connect(name string) *gorose.Engin {
	if conf, ok := container.Get("config").(*config.Config); ok {
		configs := conf.Get("database.connections." + name)
		var dbConf gorose.Config
		if err := mapstructure.Decode(configs, &dbConf); err != nil {
			fmt.Printf("database.connection.%#v is not correct %#v\n", name, configs)
		}
		if engin, err := gorose.Open(&dbConf); err == nil {
			return engin
		}
	}
	return nil
}

func Connection(name string) gorose.IOrm {
	key := "db_" + name
	if ins := container.Get(key); ins != nil {
		if db, ok := container.Get(key).(*gorose.Engin); ok {
			return db.NewOrm()
		}
	}

	if db := connect(name); db != nil {
		container.Put(key, db)
		return db.NewOrm()
	}
	return nil
}
