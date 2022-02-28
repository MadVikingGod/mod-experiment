package retract

import (
	"fmt"
)

type Stuff struct{}

func (Stuff) Blarg() {
	fmt.Println("Blarg")
}
