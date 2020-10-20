package enums

import f "fmt"

type gender int

const (
	UNKNOWN = iota
	MALE
	FEMALE
)

func genderInfo(g gender) {
	f.Println(g)
}

func Main() {
	i := 0
	var j bool
	f.Println("hello world!!!!!!!!!!!!!")
	f.Println(i)
	f.Println(j)
	f.Println("")
	f.Println(isZeroInt(0))
	f.Println(isZeroInt(9))
	f.Println(isZeroFloat(9))

	const nine = 9
	f.Println(isZeroFloat(nine))

	genderInfo(5)
}

func isZeroInt(num int) bool {
	return num == 0
}

func isZeroFloat(num float32) bool {
	return num == 0
}
