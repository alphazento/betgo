package main

import (
	"fmt"
	"github.com/alphazento/betgo/pkg/setting"

	"github.com/asmcos/requests"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint64
	Name string
}

func main() {
	configs := setting.configs;
	setting.load('test');
	req := requests.Requests()
	resp, _ := req.Get("https://www.google.com")
	fmt.Println(resp.R.StatusCode)
	fmt.Println(resp.R.Header)

	users := []User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"Blog":   "www.flysnow.org",
		// 	"wechat": "flysnow_org",
		// })
		c.JSON(200, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"id": c.Param("id"), "user": users[0]})
	})
	r.Run(":8080")
}
