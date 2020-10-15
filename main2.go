package main

import f "fmt"

func sayHello() {
	f.Println("Hello Jerry")
	doStuff1()
}

func doStuff1() {
	i := 0
	var j bool
	f.Println(i)
	f.Println(j)
	f.Println("")
	f.Println(isZeroInt(0))
	f.Println(isZeroInt(9))
	f.Println(isZeroFloat(9))

	const nine = 9
	f.Println(isZeroFloat(nine))
}

func isZeroInt(num int) bool {
	return num == 0
}


func isZeroFloat(num float32) bool {
        return num == 0
}
