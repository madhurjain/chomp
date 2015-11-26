package main

import (
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"os"
)

const MQTT_URL = "tcp://m11.cloudmqtt.com:18959"
const MQTT_USERNAME = "sreudmfs"
const MQTT_PASSWORD = "ivMAiXmR-v-1"
const MQTT_CLIENTID = "petfeeder-server"
const MQTT_ENDPOINT = "/feeder"

var client *MQTT.Client

func InitMQTTClient() {

	opts := MQTT.NewClientOptions().AddBroker(MQTT_URL)
	opts.SetClientID(MQTT_CLIENTID)
	opts.SetDefaultPublishHandler(onMessage)
	opts.SetUsername(MQTT_USERNAME)
	opts.SetPassword(MQTT_PASSWORD)

	//create and start a client using the above ClientOptions
	client = MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	//subscribe to the topic and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := client.Subscribe("/feeder/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	token := client.Publish("/feeder/status", 0, false, "hello from server")
	token.Wait()
}

//define a function for the default message handler
var onMessage MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	if msg.Topic() == "/feeder/status/log" {
		h.broadcast <- msg.Payload()
	}
}

func sendMessageToDevice(topic, message string) {
	if client != nil {
		token := client.Publish("/feeder/control", 0, false, message)
		token.Wait()
	}
}
