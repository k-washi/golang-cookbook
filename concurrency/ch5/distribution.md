# Golang goroutine

[Distributed Computing with Go: Practical concurrency and parallelism for Go applications (English Edition)](https://www.amazon.co.jp/dp/B076H8KDB6/ref=dp-kindle-redirect?_encoding=UTF8&btkr=1)

## M;N scheduler

 + M : ゴルーチンの数
 + N : Threadの数

## 項目

+ Goroutine(G)
+ OS thread or machine (M)
+ Context or processor (P)

M1, M2, ...MnのThreadがあるとする。P1, P2のスケジューラを設定する。
もし、G1, ..., G20のGoroutineがあるとすると、P1にG1~G10が割り振られ、同時にG11〜G20がP2に割り振られる。そして、P1からG１を取り出し、空いているMnで実行する。これをG2, ..., G10まで繰り返す。P2も同様の流れで処理する。

もし、P2のGorutineが空になったら、P1の残りのGorutineを半分奪う。
