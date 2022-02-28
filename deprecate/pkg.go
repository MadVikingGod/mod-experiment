package deprecate

import (
	"fmt"

	"github.com/madvikinggod/mod-experiment/deprecate/inner"
)

type Stuff struct{}

func (Stuff) Blarg() {
	fmt.Println("Blarg")
	fmt.Println(inner.A{})
}
