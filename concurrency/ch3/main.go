package main

import (
	"fmt"
	"strconv"
	"time"
)

//for-select pattern
/*
チャネルを介してゴルーチン内外でデータのやり取り、エラーハンドリング、終了処理について。

1. ゴルーチン外部から、チャネルを介して文字を挿入("1", "bad0")
2. ゴルーチン内部で、チャンネルから得た文字を数字に変換(エラーかどうか場合分けして、結果をresultに挿入)
3. ゴルーチン外部から、doneをcloseすることで、ゴルーチンが終了し、それに伴いresultチャネルもcloseされる。
4. ゴルーチン外部で、resultを取得していたforが終了する。

Result: 1
error: strconv.Atoi: parsing "bad1": invalid syntax
doSomethingを終了
タスク終了
Finished !!
*/

type Result struct {
	Error    error
	Response int
}

func doSomething(done <-chan interface{}, printStr <-chan string) <-chan Result {
	result := make(chan Result)

	go func() {
		defer fmt.Println("タスク終了")
		defer close(result)
		for {
			select {
			case s := <-printStr:

				i, err := strconv.Atoi(s)
				if err != nil {
					result <- Result{Error: err, Response: 0}

				} else {

					result <- Result{Error: err, Response: i}
				}

			case <-done:
				return
			}
		}
	}()
	return result
}

func main() {
	done := make(chan interface{})
	printStr := make(chan string)
	result := doSomething(done, printStr)

	//1秒後にdoSomethingを終了

	go func() {
		pushData := []string{"1", "bad1"}
		for _, v := range pushData {
			printStr <- string(v)
		}
		time.Sleep(1 * time.Second)
		fmt.Println("doSomethingを終了")
		close(done)
	}()

	//doSomething終了に伴いresultがcloseされる。
	for res := range result {
		if res.Error != nil {
			fmt.Println("error:", res.Error)
			continue
		}
		fmt.Println("Result:", res.Response)
	}

	fmt.Println("Finished !!")
}
