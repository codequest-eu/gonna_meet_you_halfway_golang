package broadcaster

import (
	"log"

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
	opts.SetClientID("half-way")
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
func (mq *mqttBroadcaster) Publish(message []byte, topic string) error {
	token := mq.client.Publish(topic, 1, false, message)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
