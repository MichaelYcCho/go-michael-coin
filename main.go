package main

import (
	"github.com/michael_cho77/go-michael-coin/explorer"
	"github.com/michael_cho77/go-michael-coin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(8000)
}
