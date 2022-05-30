package card

import (
	"fmt"
	"github.com/bahodurnazarov/goTask/pkg/types"
)

func Hello() {
	fmt.Printf("HEllo go ")


}

func Control(card *types.Card) *types.Card{
	if card.Activity != true{
		fmt.Println("Card is not Active")
		card.Activity = true
		fmt.Println("Now we changet it to Active")
	} else {
		fmt.Println("Cars is Active")
	}
	return card
}
