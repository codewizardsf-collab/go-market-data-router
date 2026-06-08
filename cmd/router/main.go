package main

import (
	"fmt"

	"github.com/haris/enterprise-engineering-portfolio/go-market-data-router/internal/book"
)

func main() {
	recipients := book.RouteTick(book.Tick{Symbol: "AAPL", Price: 214.42, Volume: 1000}, []book.Subscription{
		{ClientID: "desk-2", Symbols: map[string]bool{"MSFT": true}},
		{ClientID: "desk-1", Symbols: map[string]bool{"AAPL": true}},
	})
	fmt.Println(recipients)
}
