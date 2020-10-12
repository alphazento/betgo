package main

import (
	"github.com/alphazento/betgo/app"
)

type User struct {
	ID   uint64
	Name string
}

func main() {
	app.InitServices()

	// c := app.ConfigService()
	d := app.LogService()

	d.Info("This is tony test")

	app.Run()

	// ports := app.ConfigService().Get("app.services.globalproxy.depends_on")
	// if arrs, ok := ports.([]interface{}); ok {
	// 	for _, v := range arrs {
	// 		fmt.Printf("ports: %#v\n", v)
	// 	}
	// } else {
	// 	fmt.Print("No")
	// }

	// req := requests.Requests()
	// resp, _ := req.Get("https://www.google.com")
	// fmt.Println(resp.R.StatusCode)
	// fmt.Println(resp.R.Header)

	// users := []User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}

	// r := app.RouterService()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"Blog":   "www.flysnow.org",
	// 		"wechat": "flysnow_org",
	// 	})
	// })
	// r.GET("/users", func(c *gin.Context) {
	// 	// c.JSON(200, gin.H{
	// 	// 	"Blog":   "www.flysnow.org",
	// 	// 	"wechat": "flysnow_org",
	// 	// })
	// 	c.JSON(200, users)
	// })

	// r.GET("/users/:id", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"id": c.Param("id"), "user": users[0]})
	// })
	// r.Run(":8080")
}
