# go-mailcd

Go client for the [maild](https://github.com/vbogretsov/maild) service.

## Usage

```go

import (
    "log"
    "net/rpc"

    "github.com/streadway/amqp"
    "github.com/vbogretsov/go-mailcd"
)


func createAMQPConnection() (*amqp.Connection, error) {
    // create and return AMQP connection.
}

func main() {
    con, err := createAMQPConnection()
    if err != nil {
        log.Fatal(err)
    }
    defer con.Close()

    sender, err := mailcd.New(con, "maild")
    if err != nil {
        log.Fatal(err)
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

    if err := sender.Send(req); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        return
    }
}

```

## Licence

See the LICENCE file.
