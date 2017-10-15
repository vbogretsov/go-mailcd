# go-mailcd

Go client for the [maild](https://github.com/vbogretsov/maild) service.

## Usage

```go

import (
    "log"
    "net/rpc"

    "github.com/vbogretsov/go-mailcd"
)


func createRPCClient() (*rpc.Client, error) {
    // create RPC client.
}

func main() {
    client, err := createRPCClient()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    sender := mailcd.NewRPC(client)

    err := sender.Send("en", "email-template-name", map[string]interface{}{
        "To": []string{
            "to1@mail.com",
            "to2@mail.com",
        },
        "Cc": []string{
            "cc1@mail.com",
            "cc2@mail.com"
        },
        "EmailTemplateArg1": "arg1",
        "EmailTemplateArg2": "arg2",
    })

    if err != nil {
        log.Fatal(err)
    }
}

```

## Licence

See the LICENCE file.
