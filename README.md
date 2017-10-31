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

    sender := mailcd.NewRPCSender(client)

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
