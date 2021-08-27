package main

import (
	"context"

	"0chain.net/smartcontract/magmasc/bench/state-generator/cli"
)

func main() {
	app := cli.New()

	if err := cli.Start(context.Background(), app); err != nil {
		panic(err)
	}
}
