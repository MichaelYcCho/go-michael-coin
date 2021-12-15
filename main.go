package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/michael_cho77/go-michael-coin/blockchain"
)

const port string = ":8000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://127.0.0.1%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
