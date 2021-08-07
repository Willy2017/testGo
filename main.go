package main

import (

	//"io"

	"fmt"
)

/*type TestStruct struct {
	A int
	B int
	C int
	D string
}*/

func testFunction(a *int) {
	*a = 2
}

func main() {
	//rand.Seed(time.Now().Unix())
	//target := rand.Intn(100)
	//println(target)

	// logger := log.Logger{};
	/*
		var a [2]TestStruct

		a[0] = TestStruct{1, 2, 3, "hello"}
		a[1] = TestStruct{4, 5, 6, "test"}

		fmt.Println(a[0].A, a[0].B, a[0].C, a[0].D)
	*/
	a := 1
	fmt.Printf("before testFunction, a=%d\n", a)
	testFunction(&a)
	fmt.Printf("after testFunction, a=%d\n", a)

	b := 10
	c := &b
	fmt.Printf("b=%d, *c=%d", b, *c)
}
