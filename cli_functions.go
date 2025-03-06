package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	// ensure username is provided
	if len(cmd.arguments) == 0 {
		return errors.New("error: command missing arguments")
	}

	// update config file
	newValue := cmd.arguments[0]
	err := s.config.SetUser(newValue)
	if err != nil {
		return err
	}

	fmt.Printf("username successfully set to %v\n", newValue)
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	fn, exists := c.cmds[cmd.commandName]
	if !exists {
		return fmt.Errorf("unknown command '%v'", cmd.commandName)
	}
	return fn(s, cmd)
}

