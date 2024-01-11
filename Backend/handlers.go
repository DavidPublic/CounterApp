package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// setupRouter sets up the routes and handlers for the Gin router
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/counter", createCounterHandler)
	router.POST("/counter/:name/increment", incrementCounterHandler)
	router.GET("/counter/:name", getCounterHandler)
	router.DELETE("/counter/:name", deleteCounterHandler)
	router.GET("/counters", listCountersHandler)

	return router
}

// createCounterHandler handles the creation of a new counter
func createCounterHandler(c *gin.Context) {
    var newCounter Counter
    if err := c.BindJSON(&newCounter); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Pass the global DynamoDB client and the new counter to the CreateCounter function
    err := CreateCounter(db, &newCounter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusCreated)
}

// incrementCounterHandler handles incrementing the value of a counter
func incrementCounterHandler(c *gin.Context) {
	name := c.Param("name")
	err := IncrementCounter(name, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// getCounterHandler handles retrieving a specific counter
func getCounterHandler(c *gin.Context) {
	name := c.Param("name")
	counter, err := GetCounter(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if counter == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Counter not found"})
		return
	}

	c.JSON(http.StatusOK, counter)
}

// deleteCounterHandler handles the deletion of a counter
func deleteCounterHandler(c *gin.Context) {
	name := c.Param("name")
	err := DeleteCounter(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// listCountersHandler handles listing all counters
func listCountersHandler(c *gin.Context) {
	counters, err := GetAllCounters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, counters)
}
