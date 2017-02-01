package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/dzyanis/olyalya/client"
)

var (
	commands = map[string]func(args... interface{}) (string, error){
		"HELP": handlerHelp,
		//"EXIT": handlerHelp,

		"LIST": handlerHelp,
		"CREATE": handlerCreate,
		"SELECT": handlerHelp,
		"DESTROY": handlerHelp,

		"GET": handlerGet,
		"DEL": handlerHelp,
		"SET": handlerSet,
		"TTL": handlerHelp,
		"HAS": handlerHelp,

		"ARR/INDEX/GET": handlerHelp,
		"ARR/INDEX/SET": handlerHelp,
		"ARR/INDEX/DEL": handlerHelp,

		"HASH/KEY/GET": handlerHelp,
		"HASH/KEY/SET": handlerHelp,
		"HASH/KEY/DEL": handlerHelp,
	}

	Client *client.Client
)

func init() {
	Client = client.NewClient("localhost", 8080)
}

func handlerHelp (args... interface{}) (string, error) {
	return "Good luck! You really cool =*", nil
}

func handlerCreate (args... interface{}) (string, error) {
	err := Client.Create("dz")
	return "", err
}

func handlerSet (args... interface{}) (string, error) {
	err := Client.Set("author", "Dzyanis Kuzmenka")
	return "", err
}
func handlerGet (args... interface{}) (string, error) {
	s, err := Client.Get("author")
	return s, err
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("O(lya-lya) greets you")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		for cmd, f := range commands {
			if strings.Compare("EXIT", text) == 0 {
				fmt.Println("Bye!")
				return
			}
			if strings.Compare(cmd, text) == 0 {
				r, e := f()
				if e!=nil {
					fmt.Println("ERRPR: ", e.Error())
				} else {
					fmt.Println(r)
				}
			}
		}
	}
}
