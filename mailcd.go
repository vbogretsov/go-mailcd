package mailcd

import (
	"net/rpc"
)

type Sender interface {
	Send(lang string, tid string, args interface{}) error
}

type rpcSender struct {
	client *rpc.Client
}

func (self *rpcSender) Send(lang string, name string, args interface{}) error {
	res := struct{}{}
	req := map[string]interface{}{
		"Lang": lang,
		"Name": name,
		"Args": args,
	}
	return client.Call("Maild.Send", req, &res)
}

func NewRPC(client *rpc.Client) Sender {
	return &rpcSender{client: client}
}
