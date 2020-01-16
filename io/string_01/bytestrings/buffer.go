package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

//Buffer io.Reader interfaceの実装
func Buffer(rawString string) *bytes.Buffer {
	rawBytes := []byte(rawString)

	//バファーを作成(1)
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	//method 2
	b = bytes.NewBuffer(rawBytes)

	//method 3
	b = bytes.NewBufferString(rawString)

	return b

}

//ToString
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
