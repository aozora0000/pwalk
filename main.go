package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"strings"
)

var version = "unknown"
var revision = "unknown"

func main() {
	app := cli.NewApp()
	app.Usage = "Awk like Process Generator"
	app.Version = version + "@" + revision
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "sep",
			Aliases: []string{"S"},
			Usage:   "separate char from stdin",
			Value:   "",
		},
		&cli.BoolFlag{
			Name:    "exception",
			Aliases: []string{"E"},
			Usage:   "stop process when error",
			Value:   false,
		},
		&cli.BoolFlag{
			Name:    "dry-run",
			Aliases: []string{"D"},
			Usage:   "create run command string",
		},
	}
	app.Action = func(context *cli.Context) error {
		if context.NArg() == 0 {
			return errors.New("required command argument")
		}
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var line = scanner.Text()
			for _, command := range context.Args().Slice() {
				if context.String("sep") != "" {
					for i, char := range strings.Split(line, context.String("sep")) {
						command = strings.Replace(command, fmt.Sprintf("%%%v", i+1), char, -1)
					}
				} else {
					command = strings.Replace(command, "%1", line, -1)
				}
				if context.Bool("dry-run") {
					fmt.Println(command)
				} else {
					cmd := exec.Command("sh", []string{"-c", command}...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()
					if err != nil && context.Bool("exception") {
						return err
					}
				}
			}
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
