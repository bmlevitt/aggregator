package main

import (
	"log"
	"os"

	"github.com/bmlevitt/aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {log.Fatalf("error reading config file: %v\n", err)}

	cliState := state{config: &cfg}
	cmds := &commands{cmds: make(map[string]func(*state, command) error),}
	cmds.register("login", handlerLogin)
	args := os.Args

	if len(args) < 2 {log.Fatalf("Error: No command provided")} 
	if len(args) < 3 {log.Fatalf("Error: no username provided")}

	cmd := command{commandName: args[1], arguments: args[2:]}
	cmds.run(&cliState, cmd)
}