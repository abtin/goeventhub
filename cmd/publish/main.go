package main

import (
	"context"
	"fmt"
	eventhub "github.com/Azure/azure-event-hubs-go/v3"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Name: "publish - A cli to publish messages to Azure Event hubs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "connectionString",
				Aliases: []string{"c"},
				Usage:   "Event hubs connection string",
				EnvVars: []string{"EVENTHUB_CONNECTION_STRING"},
			},
		},
		Action: func(c *cli.Context) error {
			hub, err := eventhub.NewHubFromConnectionString(c.String("connectionString"))
			if err != nil {
				return err
			}

			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			// send a single message into a random partition
			err = hub.Send(ctx, eventhub.NewEventFromString("hello, Gopher!"))
			if err != nil {
				return err
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fail(err)
	}

}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}
