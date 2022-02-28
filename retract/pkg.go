package retract

import (
	"fmt"
)

type Stuff struct{}

func (Stuff) MoreStuff() {
	fmt.Println("MoreStuff")
}
