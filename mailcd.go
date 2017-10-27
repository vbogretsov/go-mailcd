package mailcd

import (
	"github.com/streadway/amqp"
	"github.com/vmihailenco/msgpack"
)

type Address struct {
	Name  string
	Email string
}

type Request struct {
	TemplateLang string
	TemplateName string
	TemplateArgs interface{}
	To           []Address
	Cc           []Address
}

type Sender interface {
	Send(*Request) error
}

type AMQPSender struct {
	exchange string
	key      string
	channel  *amqp.Channel
}

func (s *AMQPSender) Send(r *Request) error {
	body, err := msgpack.Marshal(r)

	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		Body: body,
	}

	return s.channel.Publish(s.exchange, s.key, false, false, msg)
}

func (s *AMQPSender) Close() error {
	return s.channel.Close()
}

func NewAMQPSender(conn *amqp.Connection, exchange, key string) (*AMQPSender, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	sender := AMQPSender{
		channel:  channel,
		exchange: exchange,
		key:      key,
	}

	return &sender, nil
}
