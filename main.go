package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/menghua6/token-net/api"
	"github.com/menghua6/token-net/db"
)

var port string

func main() {
	flag.StringVar(&db.MysqlUserName, "mysql-username", "", "mysql username")
	flag.StringVar(&db.MysqlPassword, "mysql-password", "", "mysql password")
	flag.StringVar(&db.MysqlHost, "mysql-host", "", "mysql host")
	flag.StringVar(&db.MysqlPort, "mysql-port", "", "mysql port")
	flag.StringVar(&db.MysqlDbName, "mysql-db-name", "", "mysql-db-name")
	flag.StringVar(&port, "port", "", "port")
	flag.Parse()

	if err := db.InitMysql(); err != nil {
		log.Println(err.Error())
		return
	}

	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		response, err := api.Introduction()
		if err != nil {
			log.Println(err.Error())
			context.String(400, "请求错误")
		} else {
			context.String(200, response)
		}
	})

	r.GET("/token/:limit", func(context *gin.Context) {
		limit := context.Param("limit")
		response, err := api.Token(limit)
		if err != nil {
			log.Println(err.Error())
			context.String(400, "请求错误")
		} else {
			context.String(200, response)
		}
	})

	r.GET("/:token", func(context *gin.Context) {
		token := context.Param("token")
		if len(token) != 64 {
			context.String(200, "")
		} else {
			response, err := api.CheckMessages(token)
			if err != nil {
				log.Println(err.Error())
				context.String(400, "请求错误")
			} else {
				context.String(200, response)
			}
		}
	})

	r.GET("/:token/:message", func(context *gin.Context) {
		token := context.Param("token")
		message := context.Param("message")
		if len(token) != 64 {
			context.String(200, "")
		} else {
			response, err := api.SendMessage(token, message)
			if err != nil {
				log.Println(err.Error())
				context.String(400, "请求错误")
			} else {
				context.String(200, response)
			}
		}
	})

	r.Run(port)

}
