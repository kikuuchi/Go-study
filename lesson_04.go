package main

import (
	"fmt";
	"math/rand"
	"time"
);

func main() {
	// ループ処理

	// 1から4までを表示
	for i := 1; i<= 4; i++ {
		fmt.Println(i)
	}

	// 1から10までの数字をランダムで5回表示する
	rand.Seed(time.Now().Unix()) 
	for i := 1; i<= 5; i++ {
		fmt.Println(rand.Intn(10))
	}

	// for…range (foreach)
	// v にインデックスが代入されるのを防ぐため、_, を記述する
	for _, v := range []string{"おはよう", "こんにちわ", "こんばんわ"} {
		fmt.Println(v)
	}

	// 連想配列(map)の場合
	// k にキー、v に値が代入される
	for k, v := range map[string]string{"key1": "おはよう", "key2": "こんにちわ", "key3": "こんばんわ"} {
		fmt.Println(k, v)
	}
}