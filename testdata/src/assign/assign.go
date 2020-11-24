package main

type A struct {
	x *int
}

func main() {
	var aSlice []A

	println("loop expecting 10, 11, 12, 13")

	for _, p := range []int{10, 11, 12, 13} {
		a := A{}
		a.x = &p // want "taking a pointer for the loop variable p"
		aSlice = append(aSlice, a)
	}

	println(`slice expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range aSlice {
		printp(p.x)
	}
}

func printp(p *int) {
	println(*p)
}
