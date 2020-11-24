package main

func main() {

	println("loop expecting 10, 11, 12, 13")

	for _, p := range []int{10, 11, 12, 13} {
		x := &p // want nothing
	}

	println(`slice expecting "13"`)
	printp(x)
}

func printp(p *int) {
	println(*p)
}
