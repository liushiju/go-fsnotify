/*
author: @liushiju
time: 2023-04-11
*/

package main

import (
	"os"

	"github.com/liushiju/go-fsnotify/cmd/app"
)

func main() {
	cmd := app.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	return
}
