package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func debug(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", root)
	router.POST("/echo", echo)

	router.Run(":" + port)
}

func root(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}

func echo(c *gin.Context) {
	var data gin.H

	err := c.BindJSON(&data)
	debug(err)

	jsonBytes, err := json.Marshal(data)
	debug(err)

	fmt.Println(">>", string(jsonBytes))

	c.JSON(http.StatusOK, data)
}
