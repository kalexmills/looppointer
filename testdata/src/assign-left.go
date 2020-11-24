package main

type A struct {
	x *int
}

func main() {

	println("loop expecting 10, 11, 12, 13")

	var x int
	a := A{&x}
	for _, p := range []int{10, 11, 12, 13} {
		*a.x = p // want nothing
	}

	println(`slice expecting "13"`)
	printp(a.x)
}

func printp(p *int) {
	println(*p)
}
