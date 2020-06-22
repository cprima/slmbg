package main

import (
	"github.com/cprior/slmbg/cmd"
	"github.com/cprior/slmbg/screensize"
)

func init() {}

func main() {

	cmd.Execute()
	screensize.GetMonitors()

}
