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
    router.DELETE("/counter/:name", deleteCounterHandler)
    router.GET("/counters", listCountersHandler)

    return router
}

func createCounterHandler(c *gin.Context) {
    var newCounter Counter
    if err := c.BindJSON(&newCounter); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    CreateCounter(&newCounter)
    c.Status(http.StatusCreated)
}

func incrementCounterHandler(c *gin.Context) {
    name := c.Param("name")
    IncrementCounter(name, 1)
    c.Status(http.StatusOK)
}

func getCounterHandler(c *gin.Context) {
    name := c.Param("name")
    counter, err := GetCounter(name)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Counter not found"})
        return
    }
    c.JSON(http.StatusOK, counter)
}

func deleteCounterHandler(c *gin.Context) {
    name := c.Param("name")
    DeleteCounter(name)
    c.Status(http.StatusOK)
}

func listCountersHandler(c *gin.Context) {
    counters := GetAllCounters()
    c.JSON(http.StatusOK, counters)
}
