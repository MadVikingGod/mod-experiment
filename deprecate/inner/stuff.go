// Deprecated: This package is deprecated to test builds with deprecated packages
package inner

import "fmt"

type Stuff struct{}

func (Stuff) Blarg() {
	fmt.Println("Inner Blarg")
}
