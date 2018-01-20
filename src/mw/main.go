package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"mw/internal/build"
	"mw/internal/cmd"
	"mw/internal/get"
	"mw/internal/new"
	"mw/internal/run"
)

func init() {
	cmd.Commands = []*cmd.Command{
		run.CmdRun,
		build.CmdBuild,
		get.CmdGet,
		new.CmdNew,
	}
}

func main() {
	flag.Usage = cmd.Usage

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		mwUsage()
		os.Exit(2)
	}

	subCommand := args[0]

	for _, c := range cmd.Commands {
		if c.CmdName == subCommand {
			if c.HasNoFlags {
				args = args[1:]
			} else {
				c.CmdFlag.Parse(args[1:])
				args = c.CmdFlag.Args()
			}
			c.Run(c, args)
			os.Exit(0)
			return
		}
	}

	log.Printf("Unknown command '%s'", subCommand)
	os.Exit(2)
}

func init() {
	cmd.Usage = mwUsage
}

func mwUsage() {
	fmt.Println("usage: mw [command] [arguments...]")
	os.Exit(1)
}
