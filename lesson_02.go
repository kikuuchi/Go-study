package main

import (
	"fmt";
);

func main() {
	// 配列に値を代入する方法

	// var 変数名 データ型を指定し、値を代入する方法
	var message_array_01 [3]string
	message_array_01[0] = "おはよう"
	message_array_01[1] = "こんにちわ"
	message_array_01[2] = "こんばんわ"
	fmt.Println(message_array_01);

	// 配列の宣言と同時に初期化を行う
	var message_array_02 = [3]string{"おはよう", "こんにちわ", "こんばんわ"}
	fmt.Println(message_array_02);

	// さらにシンプルに値を代入する方法
	message_array_03 := [3]string{"おはよう", "こんにちわ", "こんばんわ"}
	fmt.Println(message_array_03);

	// 要素数の記述を省略する方法
	message_array_04 := [...]string{"おはよう", "こんにちわ", "こんばんわ"}
	fmt.Println(message_array_04);
}