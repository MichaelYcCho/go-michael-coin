package main

import (
	"github.com/michael_cho77/go-michael-coin/cli"
	"github.com/michael_cho77/go-michael-coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
