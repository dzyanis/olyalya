package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
	"strings"
	"strconv"
	"github.com/dzyanis/olyalya/cmd"
	"github.com/dzyanis/olyalya/client"
)

var (
	Client *client.Client
	Cmd = cmd.NewCmd()
)

func handlerSetTTL(c *cmd.Cmd, args []string, line string) (string, error) {
	if len(args) < 3 {
		return "", errors.New("Not enough arguments")
	}

	ttl, err := strconv.Atoi(args[2]);
	if err!=nil {
		return "", err
	}

	err = Client.SetTTL(args[1], ttl)
	if err!=nil {
		return "", err
	}

	return "OK", nil
}

func handlerDeleteTTL(c *cmd.Cmd, args []string, line string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("Not enough arguments")
	}

	err := Client.DelTTL(args[1])
	if err!=nil {
		return "", err
	}

	return "OK", nil
}

func init() {
	Client = client.NewClient("localhost", 8080)

	Cmd.Add("HELP", &cmd.Command{
		Title: "Function show information about other functions",
		Description: "Example: HELP <FUNCTION_NAME>",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			if len(args) > 1 {
				command, ok := c.Commands[ args[1] ]
				if !ok {
					return "", cmd.ErrCommandNotExist
				}

				return fmt.Sprintf("%s\n%s", command.Title, command.Description), nil
			}

			command, _ := c.Commands[args[0]]
			result := fmt.Sprintf("%s\n%s\n\nList of commands:", command.Title, command.Description);
			for name, command := range c.Commands {
				result = result + fmt.Sprintf("\n%s - %s", name, command.Title);
			}

			return result, nil
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
	Cmd.Add("EXIT", &cmd.Command{
		Title: "Exit from the program",
		Description: "",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			fmt.Println("Bye!")
			os.Exit(0)
			return "", errors.New("Something really went wrong")
		},
	})

	Cmd.Add("INST/CREATE", &cmd.Command{
		Title: "Create an instance",
		Description: "Example: INST/CREATE dbname",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			if len(args) < 2 {
				return "", errors.New("Not enough arguments")
			}

			err := Client.InstCreate(args[1])
			if err!=nil {
				return "", err
			}
			return "OK", nil
		},
	})

	Cmd.Add("INST/LIST", &cmd.Command{
		Title: "Show list of instance",
		Description: "Example: INST/LIST",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			list, err := Client.InstList()
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

	Cmd.Add("INST/SELECT", &cmd.Command{
		Title: "Select an instance",
		Description: "Example: INST/SELECT instance_name",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			if len(args) < 2 {
				return "", errors.New("Not enough arguments")
			}

			err := Client.InstSelect(args[1])
			if err!=nil {
				return "", err
			}
			return "OK", nil
		},
	})

	Cmd.Add("SET", &cmd.Command{
		Title: "Set value",
		Description: "Example: SET name \"value\" ttl",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			if len(args) < 3 {
				return "", errors.New("Not enough arguments")
			}

			ttl := 0
			if len(args) > 3 {
				ttl, _ = strconv.Atoi(args[3]);
			}

			err := Client.Set(args[1], args[2], ttl)
			if err!=nil {
				return "", err
			}

			return "OK", nil
		},
	})
	Cmd.Add("GET", &cmd.Command{
		Title: "Get value",
		Description: "Example: GET name",
		Handler: func(c *cmd.Cmd, args []string, line string) (string, error) {
			if len(args) < 2 {
				return "", errors.New("Not enough arguments")
			}

			value, err := Client.Get(args[1])
			if err!=nil {
				return "", err
			}

			return value, nil
		},
	})

	Cmd.Add("TTL/SET", &cmd.Command{
		Title: "Set time to live",
		Description: "Example: TTL/SET mayfly 86400",
		Handler: handlerSetTTL,
	})

	Cmd.Add("TTL/DEL", &cmd.Command{
		Title: "Remove time to live",
		Description: "Example: TTL/DEL mayfly",
		Handler: handlerDeleteTTL,
	})
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("O(lya-lya) greets you")

	for {
		fmt.Printf("%s> ", Client.InstName())
		cli, err := reader.ReadString('\n')
		if err!=nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			continue
		}

		result, err := Cmd.Run(cli)
		if err!=nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			continue
		}
		fmt.Println(result)
	}
}
