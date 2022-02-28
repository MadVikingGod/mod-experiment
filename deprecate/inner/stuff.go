package inner

type Stuff struct{}

func (Stuff) Blarg() {
	fmt.Println("Inner Blarg")
}
