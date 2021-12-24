package main

// What is function result
func main() {
	v := 5
	p := &v
	println(*p)

	changePointer(p)
	println(*p)
}

func changePointer(p *int) {
	v := 3
	p = &v
}

//5,5 coz p = &v override copy of pointer but not write value

func changePointerCorrect(p *int) {
	v := 3
	*p = v
}