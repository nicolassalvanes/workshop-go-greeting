package main

import "github.com/gin-gonic/gin"
import "github.com/nicolassalvanes/workshop-go-greeting/rest"

type Greeting struct {
	Greeting string `json:"greeting"`
}

func main() {
	r := gin.Default()
	client := rest.New()
	response := new(Greeting)
	client.Get("https://go-workshop-meli.herokuapp.com/greeting", nil, response)

	r.GET("/greet", addCors, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"owner": "Nicolás Salvanés",
			"greeting": response.Greeting,
			"repository": "https://github.com/nicolassalvanes/workshop-go-greeting",
		})
	})
	r.Run()
}

func addCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Authorization, X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Allow-Request-Method")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	c.Header("Allow", "GET, POST, OPTIONS, PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}