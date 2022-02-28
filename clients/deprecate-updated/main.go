package main

import (
	"github.com/madvikinggod/mod-experiment/deprecate"
	"github.com/madvikinggod/mod-experiment/deprecate/inner"
)

func main() {
	deprecate.Stuff{}.Blarg()
	inner.Stuff{}.Blarg()
}
