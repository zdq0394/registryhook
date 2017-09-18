package hook

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	defaultListenAddr = ":8080"
)

// Start ...
func Start() {
	r := gin.Default()
	r.Use(latencyMiddlewareHandler)
	r.Use(auditMiddlewareHandler)
	// set routes
	setRoutes(r)
	// start
	startAt(r, "127.0.0.1:8080")
}

func startAt(r *gin.Engine, addrs ...string) {
	switch len(addrs) {
	case 0:
		go func() {
			r.Run(defaultListenAddr)
		}()
	default:
		for _, addr := range addrs {
			tAdd := addr
			go func() {
				r.Run(tAdd)
			}()
		}
	}
	select {}

}

func latencyMiddlewareHandler(c *gin.Context) {
	start := time.Now().UnixNano() / 1000000
	c.Next()
	end := time.Now().UnixNano() / 1000000
	latency := end - start
	fmt.Println(c.Request.URL.Path, ":", latency)
}

func auditMiddlewareHandler(c *gin.Context) {
	fmt.Println("c.Request.URL.RawPath:", c.Request.URL.RawPath)
	c.Next()
	fmt.Println("c.Writer.Status():", c.Writer.Status())
	fmt.Println("c.Writer.Size():", c.Writer.Size())
}
