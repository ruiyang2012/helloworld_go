package hello

import (
	"net/http"
	"github.com/gin-gonic/gin"
    "appengine"
    "appengine/datastore"
	"strconv"
)

type Counter struct {
    Count int
}

var actionList = []string{"add", "subtract", "multiple"}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func SetupRoutes() *gin.Engine {
	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define your handlers
	r.GET("/", func(c *gin.Context){
		c.String(200, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context){
		c.String(200, "pong")
	})

	r.GET("/v1/count", func(c *gin.Context){
		appengineContext := appengine.NewContext(c.Request)
	    key := datastore.NewKey(appengineContext, "Counter", "mycounter", 0, nil)
	    count := new(Counter)
	    err := datastore.RunInTransaction(appengineContext, func(c appengine.Context) error {
	        // Note: this function's argument c shadows the variable c
	        //       from the surrounding function.
	        err := datastore.Get(c, key, count)
	        if err != nil && err != datastore.ErrNoSuchEntity {
	            return err
	        }
	        
	        _, err = datastore.Put(c, key, count)
	        return err
	    }, nil)
	    if err != nil {
			c.String(500, "error in datastore")
	        return
	    }
		c.String(200, strconv.Itoa(count.Count))
	})	
	
	r.DELETE("/v1/count", func(c *gin.Context){
		appengineContext := appengine.NewContext(c.Request)
	    key := datastore.NewKey(appengineContext, "Counter", "mycounter", 0, nil)
	    count := new(Counter)
	    err := datastore.RunInTransaction(appengineContext, func(c appengine.Context) error {
	        // Note: this function's argument c shadows the variable c
	        //       from the surrounding function.
	        err := datastore.Get(c, key, count)
	        if err != nil && err != datastore.ErrNoSuchEntity {
	            return err
	        }
	        count.Count = 0
	        _, err = datastore.Put(c, key, count)
	        return err
	    }, nil)
	    if err != nil {
			c.String(500, "error in datastore")
	        return
	    }
		c.String(200, "reset")
	})	
	
	r.POST("/v1/count/:action/:number", func(c *gin.Context){
		num := c.Param("number")
		act := c.Param("action")
		
		if !stringInSlice(act, actionList) {
			c.String(404, "action not supported")
			return			
		}
		
		i, err := strconv.Atoi(num)
		if err != nil {
			c.String(404, "not a number")
			return
		}
		appengineContext := appengine.NewContext(c.Request)
	    key := datastore.NewKey(appengineContext, "Counter", "mycounter", 0, nil)
	    count := new(Counter)
	    err0 := datastore.RunInTransaction(appengineContext, func(c appengine.Context) error {
	        // Note: this function's argument c shadows the variable c
	        //       from the surrounding function.
	        err := datastore.Get(c, key, count)
	        if err != nil && err != datastore.ErrNoSuchEntity {
	            return err
	        }
			switch act {
			case "add":
				count.Count += i
			case "subtract":
				count.Count -= i
			case "multiple":
				count.Count *= i
			}
	        
	        _, err = datastore.Put(c, key, count)
	        return err
	    }, nil)
	    if err0 != nil {
			c.String(500, "error in datastore")
	        return
	    }		
		c.String(200, strconv.Itoa(count.Count))
	})

	return r
}

// This function's name is a must. App Engine uses it to drive the requests properly.
func init() {
	// Handle all requests using net/http
	http.Handle("/", SetupRoutes())
}

