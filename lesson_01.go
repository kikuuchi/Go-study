package main

import (
	"fmt";
	"reflect";
);

func main() {
	// 変数に値を代入する方法

	// var 変数名 データ型を指定し、値を代入する方法
	var message_01 string
	message_01 = "おはよう";
	fmt.Println(message_01);
	// reflect.TypeOf を使用してデータ型を表示する
	fmt.Println(reflect.TypeOf(message_01));

	// データ型を省略し、値を代入する方法
	var message_02 = "こんにちわ";
	fmt.Println(message_02);
	fmt.Println(reflect.TypeOf(message_02));

	// さらにシンプルに、値を代入する方法
	message_03 := "こんばんわ";
	fmt.Println(message_03);
	fmt.Println(reflect.TypeOf(message_03));
}