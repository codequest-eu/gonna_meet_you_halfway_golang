package broadcaster

import (
	"encoding/json"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//mqttBroadcaster is a implementation of a Broadcaster interface
//mqttBroadcaster use mqtt protocol to communicat
type mqttBroadcaster struct {
	client mqtt.Client
}

var handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("TOPIC: %s\n", msg.Topic())
	log.Printf("MSG: %s\n", msg.Payload())
}

//NewMQTTBroadcaster is a constructor of MQTTBroadcaster struct
func NewMQTTBroadcaster(broker string, user string, pass string) (Broadcaster, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(user)
	opts.SetPassword(pass)
	opts.SetClientID("half-way-server-" + os.Getenv("HOST"))
	opts.SetCleanSession(true)
	opts.SetDefaultPublishHandler(handler)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &mqttBroadcaster{c}, nil
}

//Close use to close connections
func (mq *mqttBroadcaster) Close() error {
	mq.client.Disconnect(1)
	return nil
}

//Publish message
func (mq *mqttBroadcaster) Publish(v interface{}, topic string) error {
	message, err := json.Marshal(v)
	if err != nil {
		return err
	}
	token := mq.client.Publish(topic, byte(1), true, message)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
