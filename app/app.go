package app

import (
	"fmt"

	"github.com/alphazento/betgo/pkg/config"
	"github.com/alphazento/betgo/pkg/container"
	"github.com/alphazento/betgo/pkg/logger"
	"github.com/alphazento/betgo/pkg/orm"
	"github.com/alphazento/betgo/pkg/router"
	"github.com/alphazento/betgo/pkg/user"
	"github.com/gin-gonic/gin"
)

func InitServices() {
	container.Bind("config", config.Factory, true)
	container.Bind("log", logger.Factory, true)
	container.Bind("router", router.Factory, true)
	container.Bind("user", user.Factory, true)
	if res, err := orm.Connection("mysql").Query("select * from ink_url_rewrites limit 1"); err == nil {
		fmt.Println(res)
	}

	res, _ := orm.Connection("mysql").Table("ink_url_rewrites").First()
	// res 类型为 map[string]interface{}
	fmt.Println(res)
}

func ConfigService() *config.Config {
	return container.Get("config").(*config.Config)
}

func LogService() *logger.Logger {
	return container.Get("log").(*logger.Logger)
}

func RouterService() *gin.Engine {
	return container.Get("router").(*gin.Engine)
}

func Run() {
	if services, ok := ConfigService().Get("app.boot_services").([]interface{}); ok {
		container.Boot(services)
	}
	RouterService().Run(":8080")
}
