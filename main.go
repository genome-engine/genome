package main

import (
	"github.com/genome-engine/genome/cli"
)

func main() {
	err := cli.GetCli().Execute()
	if err != nil {
		println(err.Error())
		return
	}
}
