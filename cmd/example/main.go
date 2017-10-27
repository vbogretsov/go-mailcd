package main

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
	"github.com/vbogretsov/go-mailcd"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	sender, err := mailcd.NewAMQPSender(conn, "maild", "maild")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
	defer sender.Close()

	req := mailcd.Request{
		TemplateLang: "en",
		TemplateName: "test",
		TemplateArgs: map[string]interface{}{
			"Username": "Donald",
		},
		To: []mailcd.Address{
			{
				Email: "to1@mail.com",
			},
		},
	}

	if err := sender.Send(&req); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	fmt.Println("sent")
}
