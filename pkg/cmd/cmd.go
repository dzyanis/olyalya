package cmd

import (
	"errors"
)

var (
	ErrCommandNotExist = errors.New("Command is not exist")
)

type Command struct {
	Title       string
	Description string
	Handler     func(*Cmd, []string, string) (string, error)
}

type Cmd struct {
	Commands map[string]*Command
}

func NewCmd() *Cmd {
	return &Cmd{
		Commands: make(map[string]*Command),
	}
}

func (c *Cmd) Add(name string, rec *Command) {
	c.Commands[name] = rec
}

func (c *Cmd) Run(cli string) (string, error) {
	lexer := NewLexer()
	lexems, err := lexer.Parse(cli)
	if err != nil {
		return "", err
	}
	if len(lexems) < 1 {
		return "", ErrCommandNotExist
	}

	command, ok := c.Commands[lexems[0]]
	if !ok {
		return "", ErrCommandNotExist
	}

	return command.Handler(c, lexems, cli)
}
