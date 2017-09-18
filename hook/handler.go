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

func printEvent(event *Event) {
	fmt.Println("############  begin  ##############")
	fmt.Printf("event.ID:%s\n", event.ID)
	fmt.Printf("event.Actor.Name:%s\n", event.Actor.Name)
	fmt.Printf("event.Action:%s\n", event.Action)
	fmt.Printf("event.Request.Host:%s\n", event.Request.Host)
	fmt.Printf("event.Request.Addr:%s\n", event.Request.Addr)
	fmt.Printf("event.Request.ID:%s\n", event.Request.ID)
	fmt.Printf("event.Request.Method:%s\n", event.Request.Method)
	fmt.Printf("event.Source.Addr:%s\n", event.Source.Addr)
	fmt.Printf("event.Source.InstanceID:%s\n", event.Source.InstanceID)
	fmt.Printf("event.Target.Repository:%s\n", event.Target.Repository)
	fmt.Printf("event.Target.FromRepository:%s\n", event.Target.FromRepository)
	fmt.Printf("event.Target.Tag:%s\n", event.Target.Tag)
	fmt.Printf("event.Target.URL:%s\n", event.Target.URL)
	fmt.Println("#############  end  ################")
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
		printEvent(&event)
	}
	c.JSON(200, gin.H{
		"OK": "OK",
	})
}
