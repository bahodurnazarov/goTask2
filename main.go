package main

import (
	"github.com/bahodurnazarov/goTask/pkg/card"
	"github.com/bahodurnazarov/goTask/pkg/types"
)

func main() {
	card.Control(&types.Card{Activity: false})
}
