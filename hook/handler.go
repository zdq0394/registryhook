package hook

import (
	"fmt"
	"time"

	"io/ioutil"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func pingHanlder(c *gin.Context) {
	time.Sleep(time.Second * 2)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func eventsHandler(c *gin.Context) {
	fmt.Println("eventsHandler")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	var envelope Envelope
	// fmt.Println(string(data))
	json.Unmarshal(data, &envelope)
	events := envelope.Events
	for _, event := range events {
		fmt.Println(event.Target.Digest)
	}
}
