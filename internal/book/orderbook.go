package book

import "sort"

type Tick struct {
	Symbol string
	Price  float64
	Volume int64
}

type Subscription struct {
	ClientID string
	Symbols  map[string]bool
}

func RouteTick(tick Tick, subscriptions []Subscription) []string {
	recipients := []string{}
	for _, sub := range subscriptions {
		if sub.Symbols[tick.Symbol] {
			recipients = append(recipients, sub.ClientID)
		}
	}
	sort.Strings(recipients)
	return recipients
}
