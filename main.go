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

	const url = "C:\\Users\\Asus\\source\\repos\\Project2\\Debug\\marketdata-hc1710-20170608.txt"

	// 開檔
	inputFile, Error := os.Open(url)
	// 判斷是否開檔錯誤
	if Error != nil {
		fmt.Println("開檔錯誤!")
		return
	}
	// 離開時自動執行關檔
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for i := 0; i < 10; i++ {
		inputString, err := inputReader.ReadString('\n')

		if err == io.EOF {
			fmt.Println("已讀取到檔尾!!")
			return
		}
		fmt.Printf("%s\n", inputString)
	}
}
