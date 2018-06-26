package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "Bucket Accounting"
	app.Usage = "Organizing viritual 'envelope' accounting"
	app.Commands = []cli.Command{
		{
			Name:    "bucket",
			Aliases: []string{"b"},
			Usage:   "working with buckets",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "create a new bucket",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "list",
					Usage: "list all bucket info",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First)
						return nil
					},
				},
			},
		},
		{
			Name:    "report",
			Aliases: []string{"b"},
			Usage:   "working with buckets",
			Subcommands: []cli.Command{
				{
					Name:  "show", //shows a single report
					Usage: "create a new bucket",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First)
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Whoopps!")
	}

}

func roundDay(input time.Time) time.Time {
	//round this date to the nearest day
	return input
}

func numDays(start, end time.Time) int {
	//find number of days between two dates
	return 6
}

// reader := bufio.NewReader(os.Stdin)
// fmt.Println("Welcome to bucket accounting")
// for {
// 	mainMenu(reader)
// 	text, _ := reader.ReadString('\n')
// 	fmt.Println(text)
// 	text = strings.Replace(text, "\n", "", -1)
// 	if text == "exit" {
// 		break
// 	}
// }
