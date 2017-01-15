package main

import (
	"fmt"
	stdLog "log"
	"os"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	mqtt.DEBUG = stdLog.New(os.Stdout, "", stdLog.LstdFlags|stdLog.LUTC)
	mqtt.WARN = stdLog.New(os.Stdout, "", stdLog.LstdFlags|stdLog.LUTC)
	mqtt.ERROR = stdLog.New(os.Stdout, "", stdLog.LstdFlags|stdLog.LUTC)
	mqtt.CRITICAL = stdLog.New(os.Stdout, "", stdLog.LstdFlags|stdLog.LUTC)

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("gotrivial")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topicName := "go-mqtt/sample"
	qos := uint8(0)
	retain := false

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d", i)
		token := c.Publish(topicName, qos, retain, text)
		token.Wait()
	}

	time.Sleep(6 * time.Second)

	if token := c.Unsubscribe(topicName); token.Wait() && token.Error() != nil {
		log.Error(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}
