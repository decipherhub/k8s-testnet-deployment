package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "testnet deployment tool for k8s"
	app.Usage = "you can easily deploy blockchain testnet for any cloud vendors"

	backendFlag := []cli.Flag{
		cli.StringFlag{
			Name:  "vendor",
			Value: "IBM",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "backend",
			Usage: "select cloud vendors who serves you k8s cluster",
			Flags: backendFlag,
			Action: func(c *cli.Context) error {
				backend := c.String("vendor")
				switch backend {
				case "IBM":
					fmt.Printf("you selected %s backend\n", backend)
				case "aws":
					fmt.Printf("you selected %s backend\n", backend)
				case "gcp":
					fmt.Printf("you selected %s backend\n", backend)
				case "azure":
					fmt.Printf("you selected %s backend\n", backend)
				default:
					fmt.Printf("ERR! %s is wrong backend\n", backend)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
