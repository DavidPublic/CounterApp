package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "sort"
    "sync"
)

// Counter struct represents a counter with an ID, name, and value.
type Counter struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Value int    `json:"value"`
}

// CountersMap is a thread-safe map of counters with an auto-incrementing ID.
type CountersMap struct {
    sync.RWMutex
    m      map[string]*Counter
    nextID int
}

// NewCountersMap initializes a new CountersMap.
func NewCountersMap() *CountersMap {
    return &CountersMap{
        m:      make(map[string]*Counter),
        nextID: 1,
    }
}

// CreateCounter creates a new counter with the given name.
func (cm *CountersMap) CreateCounter(name string) {
    cm.Lock()
    defer cm.Unlock()
    cm.m[name] = &Counter{ID: cm.nextID, Name: name, Value: 0}
    cm.nextID++
}

// IncrementCounter increments the counter by 1.
func (cm *CountersMap) IncrementCounter(name string) {
    cm.Lock()
    defer cm.Unlock()
    if counter, ok := cm.m[name]; ok {
        counter.Value++
    }
}

// GetCounter returns the counter with the given name.
func (cm *CountersMap) GetCounter(name string) (*Counter, bool) {
    cm.RLock()
    defer cm.RUnlock()
    counter, ok := cm.m[name]
    return counter, ok
}

// GetAllCounters returns all counters sorted by ID.
func (cm *CountersMap) GetAllCounters() []*Counter {
    cm.RLock()
    defer cm.RUnlock()
    counters := make([]*Counter, 0, len(cm.m))
    for _, counter := range cm.m {
        counters = append(counters, counter)
    }

    // Sorting the counters slice by ID
    sort.Slice(counters, func(i, j int) bool {
        return counters[i].ID < counters[j].ID
    })

    return counters
}

var counters = NewCountersMap()

func setupRouter() *gin.Engine {
    router := gin.Default()
    
    // CORS middleware
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        // Handle browser pre-flight requests
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    })

    // Define your routes here
    router.POST("/counter", func(c *gin.Context) {
        var newCounter Counter
        if err := c.BindJSON(&newCounter); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        counters.CreateCounter(newCounter.Name)
        c.Status(http.StatusCreated)
    })

    router.POST("/counter/:name/increment", func(c *gin.Context) {
        name := c.Param("name")
        counters.IncrementCounter(name)
        c.Status(http.StatusOK)
    })

    router.GET("/counter/:name", func(c *gin.Context) {
        name := c.Param("name")
        if counter, ok := counters.GetCounter(name); ok {
            c.JSON(http.StatusOK, counter)
        } else {
            c.Status(http.StatusNotFound)
        }
    })

    router.GET("/counters", func(c *gin.Context) {
        c.JSON(http.StatusOK, counters.GetAllCounters())
    })

    return router
}

func main() {
    router := setupRouter()
    router.Run(":8080")
}
