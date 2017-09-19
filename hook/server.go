package hook

import (
	"time"

	"github.com/Sirupsen/logrus"
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
	startAt(r, "0.0.0.0:8080")
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
	logrus.Infoln(c.Request.URL.Path, ":", latency)
}

func auditMiddlewareHandler(c *gin.Context) {
	logrus.Infoln("c.Request.URL.RawPath:", c.Request.URL.RawPath)
	c.Next()
	logrus.Infoln("c.Writer.Status():", c.Writer.Status(), "--c.Writer.Size():", c.Writer.Size())
}
