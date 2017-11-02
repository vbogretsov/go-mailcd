package main

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/streadway/amqp"
	"github.com/vbogretsov/amqprpc"
	"github.com/vbogretsov/go-mailcd"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	clientCodec, err := amqprpc.NewClientCodec(conn, "maild")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
	defer clientCodec.Close()

	client := rpc.NewClientWithCodec(clientCodec)
	sender := mailcd.NewRPCSender(client)

	req := mailcd.Request{
		TemplateLang: "en",
		TemplateName: "example",
		TemplateArgs: map[string]interface{}{
			"Username": "Donald",
		},
		To: []mailcd.Address{
			{
				Email: "to1@mail.com",
			},
		},
	}

	if err := sender.Send(req); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	fmt.Println("sent")
}
