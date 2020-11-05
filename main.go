package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var count = 0

func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post call",
	})
}

func GetHomePage(c *gin.Context) {
	count += 1
	currentTime := time.Now()
	var res string = ""
	response, err := http.Get("https://picsum.photos/200/300")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		res = b64.StdEncoding.EncodeToString(data)
	}
	c.JSON(200, gin.H{
		"message":  "Hello world",
		"count":    count,
		"current":  currentTime.Format("2006.01.02 15:04:05"),
		"response": res,
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Print("hello world")
	r := gin.Default()
	r.GET("/", GetHomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings)
	r.GET("/path/:name/:age", PathParameters)
	r.Run()
}
