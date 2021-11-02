package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string   `form:"name"`
	Address string   `form:"address"`
	Test    []string `form:"test"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Test)
		log.Println(len(person.Test))
	}
	c.String(200, "Success")
}
