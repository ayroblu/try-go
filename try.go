package try

import "fmt"

func main() {
	fmt.Printf("%+v\n", 0x2p10)
	fmt.Printf("%+v\n", 0x2.p10)
	fmt.Printf("%+v\n", 0x2p32)
	fmt.Printf("%+v\n", 0x2p64)
	fmt.Printf("%+v\n", 0x2p65)
	fmt.Printf("%+v\n", 0x2p128)
	fmt.Println("Hello, World!")
	t := struct{ Name string }{Name: "Alice"}
	fmt.Println(t) // Output: {Alice}
	s := []string{}
	fmt.Println("s", s)
	s = append(s, "hi")
	fmt.Println("max int", ^uint(0))

	fmt.Println("KB", KB)
	fmt.Println("KB", MB)
	var thing = Thing{}
	read(thing)
}

func read(t Printer) {
}

const (
	_          = iota // ignore first value by assigning to blank identifier
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type Printer interface {
	Read() string
}
type Thing struct {
	a int
}

// Read implements Printer.
func (t Thing) Read() string {
	panic("unimplemented")
}
