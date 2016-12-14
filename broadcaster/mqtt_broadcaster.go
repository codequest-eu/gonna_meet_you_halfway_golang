package broadcaster

import (
	"encoding/json"
	"log"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
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

var subHander mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("TOPIC: %s\n", msg.Topic())
	log.Printf("MSG: %s\n", msg.Payload())
}

//NewMQTTBroadcaster is a constructor of MQTTBroadcaster struct
func NewMQTTBroadcaster(broker string, user string, pass string) (Broadcaster, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(user)
	opts.SetPassword(pass)
	opts.SetClientID("half-way-server")
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

func (mq *mqttBroadcaster) SubscribeMeetingSuggestion(topic string) error {
	token := mq.client.Subscribe(topic, byte(1), subHander)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

//Publish message
func (mq *mqttBroadcaster) PublishMeetingSuggestion(sugestion models.MeetingSuggestion, topic string) error {
	message, err := json.Marshal(sugestion)
	if err != nil {
		return err
	}
	token := mq.client.Publish(topic, byte(1), false, message)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
