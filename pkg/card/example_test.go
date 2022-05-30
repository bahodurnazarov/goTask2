package card

import (
	"fmt"
	"github.com/bahodurnazarov/goTask2/pkg/types"
)


func ExampleControl(card *types.Card) *types.Card{
	if card.Activity != true{
		fmt.Println("Card is not Active")
		card.Activity = true
		fmt.Println("Now we changet it to Active")
	} else {
		fmt.Println("Cars is Active")
	}
	
	return card

	// Output: Card is not Active

}