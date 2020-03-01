package main

import (
    "fmt"
    "os"
    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "hello_cli"
    app.Usage = "Print hello world"
    app.Flags = []cli.Flag {
        &cli.StringFlag {
			Name: "name",
			Aliases: []string{"n"},
            Value: "World",
            Usage: "Who to say hello to.",
        },
    }
    app.Action = func(c *cli.Context) error {
        name := c.String("name")
        fmt.Printf("Hello %s!\n", name)
        return nil
    }
    app.Run(os.Args)
}