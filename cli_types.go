package main

import (
	"github.com/bmlevitt/aggregator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	commandName string
	arguments []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}