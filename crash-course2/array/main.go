package main



func main() {
	// Array
	var a [5]int
	a[2] = 7
	println(a[2])
	println(a[3])

	// Array literal
	b := [5]int{1, 2, 3, 4, 5}
	println(b[2])
	println(b[3])

	// Array length
	println(len(b))

	// Array of arrays
	c := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	println(c[1][2])
	println(c[0][1])
	println(c[1][0])
}