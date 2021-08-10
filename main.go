package main

import (
	"fmt"
	"sync"
	"testGo/test"
)

//"io"

/*type TestStruct struct {
	A int
	B int
	C int
	D string
}*/

func main() {
	fmt.Println("Main started!")
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
	var wg sync.WaitGroup
	var ch chan string
	ch = make(chan string)
	wg.Add(1)
	go test.TestGoRoutine1(ch, &wg, 1)
	wg.Add(1)
	go test.TestGoRoutine2(ch, &wg, 2)

	wg.Wait()
	fmt.Println("Main finished!")
}
