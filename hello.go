package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	var quotes string
	quotes = quote.Go()
	fmt.Println(quotes)
}
