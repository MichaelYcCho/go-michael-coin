package main

import (
	"github.com/michael_cho77/go-michael-coin/blockchain"
	"github.com/michael_cho77/go-michael-coin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
