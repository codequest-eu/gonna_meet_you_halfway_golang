package broadcaster

import "io"

//Broadcaster is responsible for open a communication channel and message exchange
type Broadcaster interface {
	io.Closer
	Publish(v interface{}, topic string) error
}
