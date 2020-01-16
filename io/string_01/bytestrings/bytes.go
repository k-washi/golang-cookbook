package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

//WorkWithBuffer
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"
	b := Buffer(rawString)

	//easy convert
	fmt.Println(b.String())

	//generic io reader
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	reader := bytes.NewReader([]byte(rawString))

	//単語ごとに分離
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
	return nil
}

//SearchString
func SearchString() {
	s := "this is a test"
	fmt.Println(strings.Contains(s, "this"))

	//contain a or b or c
	fmt.Println(strings.ContainsAny(s, "abc"))

	fmt.Println(strings.HasPrefix(s, "thi"))

	fmt.Println(strings.HasSuffix(s, "test"))

	fmt.Println(strings.Split(s, " "))

	fmt.Println(strings.Title(s))

	//先頭と最後のスペースを削除
	fmt.Println(strings.TrimSpace(s))

}
