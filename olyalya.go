package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
	"strings"
	"github.com/dzyanis/olyalya/cmd"
	"github.com/dzyanis/olyalya/client"
)

var (
	Client *client.Client
	Cmd = cmd.NewCmd()
)

func init() {
	Client = client.NewClient("localhost", 8080)

	Cmd.Add("HELP", &cmd.Command{
		Title: "Function show information about other functions",
		Description: "Example: HELP <FUNCTION_NAME>",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			funcName := ""
			if len(args) > 1 {
				funcName = args[1]
			} else {
				funcName = args[0]
			}

			command, ok := c.Commands[ funcName ]
			if !ok {
				return "", cmd.ErrCommandNotExist
			}

			return fmt.Sprintf("%s\n%s", command.Title, command.Description), nil
		},
	})

	Cmd.Add("ECHO", &cmd.Command{
		Title: "Prints string",
		Description: "Example: ECHO \"Hello World!\"",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			if len(args) < 2 {
				return "", errors.New("Not enough arguments")
			}
			s := strings.Trim(args[1], `"`)
			return s, nil
		},
	})

	Cmd.Add("DB/CREATE", &cmd.Command{
		Title: "Create Database",
		Description: "Example: DB/CREATE dbname",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			if len(args) < 2 {
				return "", errors.New("Not enough arguments")
			}

			err := Client.Create(args[1])
			if err!=nil {
				return "", err
			}
			return "OK", nil
		},
	})

	Cmd.Add("DB/LIST", &cmd.Command{
		Title: "Show list of database",
		Description: "Example: DB/LIST",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			list, err := Client.DbList()
			if err!=nil {
				return "", err
			}
			result := ""
			for i, e := range list {
				result = result + fmt.Sprintf("%d) %s\n", i+1, e)
			}
			return strings.Trim(result, "\n"), nil
		},
	})
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("O(lya-lya) greets you")

	for {
		fmt.Print("> ")
		cli, err := reader.ReadString('\n')
		if err!=nil {
			fmt.Errorf("ERROR: %s", err)
			continue
		}

		result, err := Cmd.Run(cli)
		if err!=nil {
			fmt.Errorf("ERROR: %s", err)
			continue
		}
		fmt.Println(result)
	}
}
