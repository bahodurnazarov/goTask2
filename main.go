package main

import (
	"github.com/bahodurnazarov/goTask2/pkg/card"
	"github.com/bahodurnazarov/goTask2/pkg/types"
)

func main() {
	card.Control(&types.Card{Activity: false})
}
