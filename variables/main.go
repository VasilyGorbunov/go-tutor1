package main

func main() {
	var i int = 25
	var autoInt = -25
	var bigInt int64 = 100500
	var unsignedInt uint = 90
	var unsignedBigInt uint64 = 1005000000

	println("integers", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	var f32 float32 = 3.14
	var f64 float64 = 3.1452

	println("floats", f32, f64)

	var b bool = true
	println("bools", b)

	var hello string = "Hi\n\t"
	var world = "World"
	println(hello, world)

	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("defaults", defaultInt, defaultFloat, defaultString, defaultBool)
}
