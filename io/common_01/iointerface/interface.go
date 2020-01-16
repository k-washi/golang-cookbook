package iointerface

import (
	"fmt"
	"io"
	"os"
)

//Copy copies outにコピーかつ標準出力
func Copy(in io.ReadSeeker, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout)

	//一般的なcopy, 大量のデータが有る場合危険
	/*
		if _, err := io.Copy(w, in); err != nil {
			return err
		}
	*/

	in.Seek(0, 0)
	//64byteのチャンクを確保してコピー
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	fmt.Println()
	return nil
}
