package main

import (
	"os"

	"github.com/ancuongnguyen07/cracklePop/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
