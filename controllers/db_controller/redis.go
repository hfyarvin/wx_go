package db_controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func PingRedis(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6397",
		Password: "",
		DB:       0, //use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		fmt.Println(http.StatusOK, pong)
	}
}
