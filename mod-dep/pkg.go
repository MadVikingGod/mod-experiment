package deprecate

import (
	"fmt"
)

type Stuff struct{}

func (Stuff) MoreStuff() {
	fmt.Println("Blarg")
}
