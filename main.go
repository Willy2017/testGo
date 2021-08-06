package main

import (
	"bufio"
	"fmt"

	//"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/*type TestStruct struct {
	A int
	B int
	C int
	D string
}*/

func main() {
	// var a int
	a := 8
	fmt.Printf("hello world %d", a)
	fmt.Println("type of a is", reflect.TypeOf(a))

	var now time.Time = time.Now()
	if a == 8 {
		fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if grade >= 60 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// logger := log.Logger{};
	/*
		var a [2]TestStruct

		a[0] = TestStruct{1, 2, 3, "hello"}
		a[1] = TestStruct{4, 5, 6, "test"}

		fmt.Println(a[0].A, a[0].B, a[0].C, a[0].D)
	*/
	/*
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
			inputString, Error := inputReader.ReadString('\n')

			if Error == io.EOF {
				fmt.Println("已讀取到檔尾!!")
				return
			}
			fmt.Printf("%s\n", inputString)
		}*/
}
