package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/counter", createCounterHandler)
	router.POST("/counter/:name/increment", incrementCounterHandler)
	router.GET("/counter/:name", getCounterHandler)
	router.GET("/counters", listCountersHandler)

	return router
}

func createCounterHandler(c *gin.Context) {
	var newCounter Counter
	if err := c.BindJSON(&newCounter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	counters.CreateCounter(newCounter.Name)
	c.Status(http.StatusCreated)
}

func incrementCounterHandler(c *gin.Context) {
	name := c.Param("name")
	counters.IncrementCounter(name)
	c.Status(http.StatusOK)
}

func getCounterHandler(c *gin.Context) {
	name := c.Param("name")
	if counter, ok := counters.GetCounter(name); ok {
		c.JSON(http.StatusOK, counter)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func listCountersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, counters.GetAllCounters())
}
