/**
*  GoFlush
*  (c) 2017 Ipan Ardian
*
*  A command line app to flush or delete Redis keys easily
*  For details, see the web site: https://github.com/ipanardian/goflush
*  The MIT License
 */

package main

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/urfave/cli"
)

var args []string

func deleteKey(key string) []string {
	return []string{"redis-cli", "del", key}
}

func flushAll() []string {
	return []string{"redis-cli", "flushall"}
}

func execute(c *cli.Context) {
	binary, lookErr := exec.LookPath("redis-cli")
	if lookErr != nil {
		panic(lookErr)
	}
	if c.NArg() > 0 {
		key := c.Args().Get(0)
		args = deleteKey(key)
	} else {
		args = flushAll()
	}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "GoFlush"
	app.Usage = "A command line app to flush or delete Redis keys easily"
	app.UsageText = "goflush (no argument will flush all keys)\n\t goflush [key] (delete specific key)"
	app.Version = "1.0.0"
	app.Author = "Ipan Ardian <https://github.com/ipanardian>"
	app.Action = func(c *cli.Context) error {
		execute(c)
		return nil
	}
	app.Run(os.Args)
}
