package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	. "rest/apis"
	"net/http"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"It works")
	})

	router.POST("/person", AddPersonApi)

	router.GET("/persons", GetPersonsApi)

	router.GET("/person/:id", GetPersonApi)

	router.PUT("/person/:id", UpdatePersonApi)

	router.DELETE("/person/:id", DeletePersonApi)

	return router
}
