package mailcd

import (
	"net/rpc"
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
	Send(Request) error
}

type rpcSender struct {
	client *rpc.Client
}

func (s *rpcSender) Send(r Request) error {
	reply := struct{}{}
	return s.client.Call("Maild.Send", r, &reply)
}

func NewRPCSender(client *rpc.Client) Sender {
	return &rpcSender{client: client}
}
