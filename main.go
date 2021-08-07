package main

import (

	//"io"
	"bufio"
	"fmt"
	"io"
	"os"
)

/*type TestStruct struct {
	A int
	B int
	C int
	D string
}*/

func testFunction() int {
	const url = "C:\\Users\\Asus\\source\\repos\\Project2\\Debug\\marketdata-hc1710-20170608.txt"

	// 開檔
	inputFile, Error := os.Open(url)
	// 判斷是否開檔錯誤
	if Error != nil {
		fmt.Println("開檔錯誤!")
		return 0
	}
	// 離開時自動執行關檔
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for i := 0; i < 10; i++ {
		inputString, err := inputReader.ReadString('\n')

		if err == io.EOF {
			fmt.Println("已讀取到檔尾!!")
			break
		}
		fmt.Printf("%s\n", inputString)
	}

	fmt.Printf("End\n")
	return 3
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
	a := testFunction()
	fmt.Printf("a=%d\n", a)

	b := 10
	c := &b
	fmt.Printf("b=%d, *c=%d", b, *c)
}
