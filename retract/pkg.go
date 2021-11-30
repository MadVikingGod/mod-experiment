package retract

import (
	"fmt"

	"github.com/madvikinggod/mod-experiment/internal"
)

type Stuff struct{}

func (Stuff) Blarg() {
	fmt.Println("Blarg")
	fmt.Println(internal.A{})
}
